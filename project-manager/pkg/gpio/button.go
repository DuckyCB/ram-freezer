package gpio

import (
	"fmt"
	"os"
	"project-manager/internal/logs"
	"sync"
	"time"
)

// ButtonController controla un botón
type ButtonController struct {
	pin             int
	pressEvents     chan bool
	stopChan        chan bool
	isRunning       bool
	mu              sync.Mutex
	onPressCallback func()
	debounceTime    time.Duration
}

// NewButtonController crea un nuevo controlador para un botón
func NewButtonController(pin int) (*ButtonController, error) {
	logs.Log.Info("Creando nuevo botón")

	if !checkGPIOAccess() {
		logs.Log.Error("No se puede acceder al sistema de archivos GPIO. ¿Estás ejecutando como sudo?")
		os.Exit(1)
	}

	gpioPin, err := getGPIOPin(pin)

	err = initGPIO(gpioPin, "in")
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error al inicializar botón en pin %d: %v", pin, err))
		return nil, err
	}

	return &ButtonController{
		pin:          gpioPin,
		pressEvents:  make(chan bool),
		stopChan:     make(chan bool),
		isRunning:    false,
		debounceTime: 50 * time.Millisecond,
	}, nil
}

// SetOnPressCallback establece la función que se llamará cuando se pulse el botón
func (bc *ButtonController) SetOnPressCallback(callback func()) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.onPressCallback = callback
}

// SetDebounceTime configura el tiempo para el debounce del botón
func (bc *ButtonController) SetDebounceTime(duration time.Duration) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.debounceTime = duration
}

// StartMonitoring comienza a monitorear el estado del botón
func (bc *ButtonController) StartMonitoring() {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	if !bc.isRunning {
		bc.startMonitoringInternal()
	}
}

func (bc *ButtonController) startMonitoringInternal() {
	bc.stopChan = make(chan bool)
	bc.pressEvents = make(chan bool)
	bc.isRunning = true

	// Goroutine para leer el estado del botón
	go func() {
		var lastState int = 0
		var currentState int
		var err error

		// Leer el estado inicial
		lastState, err = readGPIO(bc.pin)
		if err != nil {
			logs.Log.Error(err.Error())
			return
		}

		for {
			select {
			case <-bc.stopChan:
				return
			default:
				currentState, err = readGPIO(bc.pin)
				if err != nil {
					logs.Log.Error(err.Error())
					time.Sleep(100 * time.Millisecond)
					continue
				}

				// Detectar flanco de bajada (1->0) para botones pull-up o
				// flanco de subida (0->1) para botones pull-down
				// Este ejemplo detecta flanco de bajada (botón pull-up)
				if lastState == 1 && currentState == 0 {
					bc.pressEvents <- true
				}

				lastState = currentState
				time.Sleep(15 * time.Millisecond) // Polling rate
			}
		}
	}()

	go func() {
		for {
			select {
			case <-bc.stopChan:
				return
			case <-bc.pressEvents:
				if bc.onPressCallback != nil {
					bc.onPressCallback()
				}
				time.Sleep(bc.debounceTime)
			}
		}
	}()
}

// StopMonitoring detiene el monitoreo del botón
func (bc *ButtonController) StopMonitoring() {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	if bc.isRunning {
		bc.stopChan <- true
		close(bc.stopChan)
		close(bc.pressEvents)
		bc.isRunning = false
	}
}

// Close libera los recursos asociados al botón
func (bc *ButtonController) Close() error {
	bc.StopMonitoring()
	return cleanupGPIO(bc.pin)
}
