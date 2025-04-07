package main

import (
	"fmt"
	"project-manager/cmd/gpio"
	"sync"
	"time"
)

const (
	buttonPin = 17
	ledPin    = 27
)

const (
	StatusIdle       = iota // Idle: slow blink
	StatusProcessing        // Processing: medium blink
	StatusCompleted         // Completed: fast blink for 10 seconds
)

const (
	BlinkSlow   = 1000 * time.Millisecond
	BlinkMedium = 500 * time.Millisecond
	BlinkFast   = 100 * time.Millisecond
)

// WorkflowController maneja el flujo del sistema
type WorkflowController struct {
	ledPin       int
	buttonPin    int
	status       int
	ledControl   *gpio.LEDController
	btnControl   *gpio.ButtonController
	processMu    sync.Mutex
	isProcessing bool
}

// NewWorkflowController crea un nuevo controlador de flujo
func NewWorkflowController(ledPin, buttonPin int) (*WorkflowController, error) {
	led, err := gpio.NewLEDController(ledPin)
	if err != nil {
		return nil, fmt.Errorf("error al inicializar LED: %v", err)
	}

	button, err := gpio.NewButtonController(buttonPin)
	if err != nil {
		led.Close()
		return nil, fmt.Errorf("error al inicializar botón: %v", err)
	}

	wfc := &WorkflowController{
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
func (wfc *WorkflowController) Start() {
	wfc.ledControl.SetBlinkSpeed(BlinkSlow)
	wfc.ledControl.StartBlinking()

	wfc.btnControl.StartMonitoring()

	fmt.Println("Sistema iniciado. Presiona el botón para comenzar el proceso.")
}

// handleButtonPress manages button press
func (wfc *WorkflowController) handleButtonPress() {
	wfc.processMu.Lock()
	defer wfc.processMu.Unlock()

	if wfc.status == StatusIdle && !wfc.isProcessing {
		wfc.isProcessing = true

		wfc.status = StatusProcessing
		wfc.ledControl.SetBlinkSpeed(BlinkMedium)

		go func() {
			fmt.Println("Iniciando proceso...")
			wfc.runSystem()

			wfc.processMu.Lock()
			wfc.status = StatusCompleted
			wfc.ledControl.SetBlinkSpeed(BlinkFast)
			wfc.processMu.Unlock()

			fmt.Println("Proceso completado. Parpadeo rápido por 10 segundos.")

			time.Sleep(10 * time.Second)

			wfc.processMu.Lock()
			wfc.status = StatusIdle
			wfc.isProcessing = false
			wfc.ledControl.SetBlinkSpeed(BlinkSlow)
			wfc.processMu.Unlock()

			fmt.Println("Sistema listo para nuevo proceso.")
		}()
	}
}

// Stop detiene el sistema y libera recursos
func (wfc *WorkflowController) Stop() {
	wfc.btnControl.StopMonitoring()
	wfc.ledControl.StopBlinking()
	wfc.btnControl.Close()
	wfc.ledControl.Close()
	fmt.Println("Sistema detenido.")
}
