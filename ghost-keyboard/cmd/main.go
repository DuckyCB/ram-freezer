package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"ram-freezer/ghost-keyboard/internal/keycodes"
	"ram-freezer/ghost-keyboard/internal/logs"
	"strconv"
	"strings"
	"time"
)

// waitTime is the time to wait between key presses
var waitTime = 50 * time.Millisecond

// waitTimeKey is the time to wait pressing a key
var waitTimeKey = 30 * time.Millisecond

func readFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		logs.Log.Error(err.Error())
		return nil, err
	}

	return file, nil
}

func openHIDG0() (*os.File, error) {
	hid, err := os.OpenFile("/dev/hidg0", os.O_WRONLY, 0666)
	if err != nil {
		logs.Log.Error(err.Error())
		return nil, err
	}
	return hid, nil
}

func processFile(scanner bufio.Scanner) error {
	hid, err := openHIDG0()
	if err != nil {
		return err
	}
	defer hid.Close()

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "wait") {
			parts := strings.SplitN(line, " ", 2)

			if len(parts) < 2 {
				logs.Log.Warn("Falta el tiempo después de 'wait': Default 1")
				parts = append(parts, "1")
			}

			waitTime, err := strconv.Atoi(parts[1])
			if err != nil {
				logs.Log.Error(fmt.Sprintf("Error: tiempo inválido después de 'wait': %v", err))
				return err
			}

			time.Sleep(time.Duration(waitTime) * time.Second)
			continue
		}

		i := 0
		for i < len(line) {
			char := line[i]

			if char == '{' {
				endIdx := strings.Index(line[i:], "}")
				if endIdx == -1 {
					logs.Log.Error(fmt.Sprintf("error: } not found in line: %s", line))
					return fmt.Errorf("error: } not found in line: %s", line)
				}
				specialKeys := line[i+1 : i+endIdx]
				writeSpecialKey(specialKeys, hid)
				i += endIdx + 1
			} else {
				err = writeChar(char, hid)
				if err != nil {
					logs.Log.Error(err.Error())
				}
				i++
			}
			// TODO: agregar caso para escribir {
		}
	}

	if err := scanner.Err(); err != nil {
		logs.Log.Error(err.Error())
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

func writeChar(char uint8, hid *os.File) error {
	keycode, ok := keycodes.Key[rune(char)]
	shift := false
	if !ok {
		keycode, ok = keycodes.KeyShift[rune(char)]
		shift = true
		if !ok {
			logs.Log.Error("keycode not found")
			return fmt.Errorf("keycode not found")
		}
	}

	modifier := byte(keycodes.KEYCODE_NONE)
	if shift {
		modifier = keycodes.KEYCODE_LEFT_SHIFT
	}

	_, err := hid.Write([]byte{modifier, 0x00, keycode, 0x00, 0x00, 0x00, 0x00, 0x00})
	if err != nil {
		logs.Log.Error(fmt.Sprintf("error writing keypress %d: %w", char, err))
		return fmt.Errorf("error writing keypress %d: %w", char, err)
	}

	time.Sleep(waitTimeKey)

	_, err = hid.Write(keycodes.Empty)
	if err != nil {
		logs.Log.Error(err.Error())
		return fmt.Errorf("error unpressing key: %w", err)
	}

	time.Sleep(waitTime) // Sleep for a while to simulate key press duration
	logs.Log.Info(fmt.Sprintf("Pressing key: %c\n", char))

	return nil
}

func writeSpecialKey(key string, hid *os.File) {
	var modKey byte = keycodes.KEYCODE_NONE

	keys := make([]byte, 0)

	if strings.Contains(key, "+") {
		parts := strings.Split(key, " + ")

		for _, part := range parts {
			if k, ok := keycodes.ModifierKey[part]; ok {
				modKey |= k
			} else if k, ok := keycodes.SpecialKey[part]; ok {
				keys = append(keys, k)
			} else if len(part) == 1 {
				// Si es una tecla normal, se procesa como un rune
				r := rune(part[0])
				if k, ok := keycodes.Key[r]; ok {
					keys = append(keys, k)
				} else if k, ok := keycodes.KeyShift[r]; ok {
					modKey |= keycodes.KEYCODE_LEFT_SHIFT
					keys = append(keys, k)
				}
			} else if strings.HasPrefix(part, ".") {
				// Si es un número, se procesa como un Keypad
				r := rune(part[1])
				if k, ok := keycodes.Keypad[r]; ok {
					keys = append(keys, k)
				} else {
					logs.Log.Warn(fmt.Sprintf("Key not found in Keypad: %s", part))
				}
			} else {
				logs.Log.Warn(fmt.Sprintf("key not found: %s", part))
			}

			if len(keys) >= 6 {
				logs.Log.Warn("Limiting to 6 keys")
				break // Limitar a 6 teclas
			}
		}
	} else {
		// Si no hay +, solo se procesa la tecla
		if k, ok := keycodes.ModifierKey[key]; ok {
			modKey |= k
		} else if k, ok := keycodes.SpecialKey[key]; ok {
			keys = append(keys, k)
		} else if len(key) == 1 {
			r := rune(key[0])
			if k, ok := keycodes.Key[r]; ok {
				keys = append(keys, k)
			} else if k, ok := keycodes.KeyShift[r]; ok {
				modKey |= keycodes.KEYCODE_LEFT_SHIFT
				keys = append(keys, k)
			}
		}
	}
	for len(keys) < 6 {
		keys = append(keys, keycodes.KEYCODE_NONE)
	}

	_, err := hid.Write([]byte{modKey, 0x00, keys[0], keys[1], keys[2], keys[3], keys[4], keys[5]})
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error writing keypress: %v", err))
		return
	}
	time.Sleep(waitTimeKey)

	_, err = hid.Write(keycodes.Empty)
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error unpressing key: %v\n", err))
		return
	}

	time.Sleep(waitTime)
	logs.Log.Info(fmt.Sprintf("Pressing special keys: %s\n", key))
}

func main() {
	logs.SetupLogger()

	filePath := flag.String("script", "", "script file to use")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Use: ./ghost-keyboard -script=file")
		os.Exit(1)
	}

	file, err := readFile(*filePath)
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error opening file %s: %v", *filePath, err))
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	err = processFile(*scanner)
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error processing file %s: %v", *filePath, err))
		return
	}
}
