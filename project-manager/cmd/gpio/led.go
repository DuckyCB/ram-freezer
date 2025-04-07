package gpio

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// LEDController controla el parpadeo del LED
type LEDController struct {
	pin        int
	isBlinking bool
	stopChan   chan bool
	blinkSpeed time.Duration
	mu         sync.Mutex
}

// NewLEDController crea un nuevo controlador para el LED
func NewLEDController(pin int) (*LEDController, error) {
	if !checkGPIOAccess() {
		log.Println("No se puede acceder al sistema de archivos GPIO. ¿Estás ejecutando como sudo?")
		os.Exit(1)
	}

	gpioPin, err := getGPIOPin(pin)

	err = initGPIO(gpioPin, "out")
	if err != nil {
		log.Printf("Error inicializando led: %v", err)
		os.Exit(1)
	}

	err = writeGPIO(gpioPin, 0)
	if err != nil {
		return nil, fmt.Errorf("error al apagar LED en pin %d (gpio %d): %v", pin, gpioPin, err)
	}

	return &LEDController{
		pin:        gpioPin,
		stopChan:   make(chan bool),
		blinkSpeed: 500 * time.Millisecond,
		isBlinking: false,
	}, nil
}

// SetBlinkSpeed cambia la velocidad de parpadeo
func (lc *LEDController) SetBlinkSpeed(speed time.Duration) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	wasBlinking := lc.isBlinking
	if wasBlinking {
		lc.stopChan <- true
		close(lc.stopChan)
		lc.isBlinking = false
	}

	lc.blinkSpeed = speed

	if wasBlinking {
		lc.startBlinkingInternal()
	}
}

// StartBlinking inicia el parpadeo en segundo plano
func (lc *LEDController) StartBlinking() {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if !lc.isBlinking {
		lc.startBlinkingInternal()
	}
}

func (lc *LEDController) startBlinkingInternal() {
	lc.stopChan = make(chan bool)
	lc.isBlinking = true

	go func() {
		for {
			select {
			case <-lc.stopChan:
				writeGPIO(lc.pin, 0)
				return
			default:
				writeGPIO(lc.pin, 1)
				time.Sleep(lc.blinkSpeed / 2)
				writeGPIO(lc.pin, 0)
				time.Sleep(lc.blinkSpeed / 2)
			}
		}
	}()
}

func (lc *LEDController) StopBlinking() {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if lc.isBlinking {
		lc.stopChan <- true
		close(lc.stopChan)
		lc.isBlinking = false
	}
}

func (lc *LEDController) TurnOn() {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if lc.isBlinking {
		lc.stopChan <- true
		close(lc.stopChan)
		lc.isBlinking = false
	}

	writeGPIO(lc.pin, 1)
}

func (lc *LEDController) TurnOff() {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if lc.isBlinking {
		lc.stopChan <- true
		close(lc.stopChan)
		lc.isBlinking = false
	}

	writeGPIO(lc.pin, 0)
}

func (lc *LEDController) Close() error {
	lc.TurnOff()
	return cleanupGPIO(lc.pin)
}
