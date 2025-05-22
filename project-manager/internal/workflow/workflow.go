package workflow

import (
	"fmt"
	"project-manager/internal/logs"
	"project-manager/internal/system"
	"project-manager/pkg/gpio"
	"sync"
	"time"
)

const (
	StatusIdle       = iota // Idle: slow blink
	StatusProcessing        // Processing: medium blink
	StatusCompleted         // Completed: fast blink for 10 seconds
)

const (
	BlinkSlow   = 2000 * time.Millisecond
	BlinkMedium = 1000 * time.Millisecond
	BlinkFast   = 200 * time.Millisecond
)

// Controller maneja el flujo del sistema
type Controller struct {
	ledPin       int
	buttonPin    int
	status       int
	ledControl   *gpio.LEDController
	btnControl   *gpio.ButtonController
	processMu    sync.Mutex
	isProcessing bool
}

// NewController crea un nuevo controlador de flujo
func NewController(ledPin, buttonPin int) (*Controller, error) {
	led, err := gpio.NewLEDController(ledPin)
	if err != nil {
		logs.Log.Error(fmt.Sprintf("error al inicializar LED: %v", err))
		return nil, err
	}

	button, err := gpio.NewButtonController(buttonPin)
	if err != nil {
		logs.Log.Error("error al inicializar botón")
		led.Close()
		return nil, err
	}

	wfc := &Controller{
		ledPin:       ledPin,
		buttonPin:    buttonPin,
		status:       StatusIdle,
		ledControl:   led,
		btnControl:   button,
		isProcessing: false,
	}

	button.SetOnPressCallback(wfc.handleButtonPress)

	return wfc, nil
}

// Start starts the system
func (wfc *Controller) Start() {
	wfc.ledControl.SetBlinkSpeed(BlinkSlow)
	wfc.ledControl.StartBlinking()

	wfc.btnControl.StartMonitoring()

	logs.Log.Info("Sistema iniciado. Esperando que se presione el botón")
}

// handleButtonPress manages button press
func (wfc *Controller) handleButtonPress() {
	err := system.StartRun()
	if err != nil {
		return
	}

	logs.Log.Info("Botón presionado")

	wfc.processMu.Lock()
	defer wfc.processMu.Unlock()

	if wfc.status == StatusIdle && !wfc.isProcessing {
		wfc.isProcessing = true

		wfc.status = StatusProcessing
		wfc.ledControl.SetBlinkSpeed(BlinkMedium)

		go func() {
			logs.Log.Info("Iniciando proceso...")
			wfc.runSystem()

			wfc.processMu.Lock()
			wfc.status = StatusCompleted
			wfc.ledControl.SetBlinkSpeed(BlinkFast)
			wfc.processMu.Unlock()

			logs.Log.Info("Proceso completado")

			time.Sleep(10 * time.Second)

			wfc.processMu.Lock()
			wfc.status = StatusIdle
			wfc.isProcessing = false
			wfc.ledControl.SetBlinkSpeed(BlinkSlow)
			wfc.processMu.Unlock()

			logs.Log.Info("Sistema listo para nuevo proceso.")
		}()
	}
}

// Stop detiene el sistema y libera recursos
func (wfc *Controller) Stop() {
	logs.Log.Info("Deteniendo sistema")
	wfc.btnControl.StopMonitoring()
	wfc.ledControl.StopBlinking()
	wfc.btnControl.Close()
	wfc.ledControl.Close()
	logs.Log.Info("Sistema detenido")
}
