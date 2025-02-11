package main

import (
	"fmt"
	"os"
	"time"
)

func writeHIDG0(text string) error {
	hid, err := os.OpenFile("/dev/hidg0", os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("error opening hidg0: %w", err)
	}
	defer hid.Close()

	for _, char := range text {
		keycode, ok := Key[char]
		if !ok {
			continue
		}

		_, err := hid.Write([]byte{0x00, 0x00, keycode, 0x00, 0x00, 0x00, 0x00, 0x00})
		if err != nil {
			return fmt.Errorf("error writing keypress: %w", err)
		}

		time.Sleep(10 * time.Millisecond)

		_, err = hid.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		if err != nil {
			return fmt.Errorf("error unpressing key: %w", err)
		}
	}
	return nil
}

func main() {
	text := "hola como estas\nbien y vos.\n"
	if err := writeHIDG0(text); err != nil {
		fmt.Println("Error:", err)
	}
}
