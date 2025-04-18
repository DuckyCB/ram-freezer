package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"ram-freezer/ghost-keyboard/internal/keycodes"
	"strconv"
	"strings"
	"time"
)

// waitTime is the time to wait between key presses
var waitTime = 100 * time.Millisecond

// waitTimeKey is the time to wait pressing a key
var waitTimeKey = 30 * time.Millisecond

func readFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func openHIDG0() (*os.File, error) {
	hid, err := os.OpenFile("/dev/hidg0", os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening hidg0: %w", err)
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
				fmt.Println("Falta el tiempo después de 'wait': Default 1")
				parts = append(parts, "1")
			}

			waitTime, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error: tiempo inválido después de 'wait'")
				return fmt.Errorf("error: tiempo inválido después de 'wait'")
			}

			fmt.Printf("Sleeping %d seconds\n", waitTime)
			time.Sleep(time.Duration(waitTime) * time.Second)
			continue
		}

		i := 0
		for i < len(line) {
			char := line[i]

			if char == '{' {
				endIdx := strings.Index(line[i:], "}")
				if endIdx == -1 {
					return fmt.Errorf("error: } not found in line: %s", line)
				}
				specialKeys := line[i+1 : i+endIdx]
				writeSpecialKey(specialKeys, hid)
				i += endIdx + 1
			} else {
				err = writeChar(char, hid)
				// TODO: handle err
				i++
			}
			// TODO: agregar caso para escribir {
		}
	}

	if err := scanner.Err(); err != nil {
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
			return fmt.Errorf("keycode not found")
		}
	}

	modifier := byte(keycodes.KEYCODE_NONE)
	if shift {
		modifier = keycodes.KEYCODE_LEFT_SHIFT
	}

	_, err := hid.Write([]byte{modifier, 0x00, keycode, 0x00, 0x00, 0x00, 0x00, 0x00})
	if err != nil {
		return fmt.Errorf("error writing keypress %d: %w", char, err)
	}

	time.Sleep(waitTimeKey)

	_, err = hid.Write(keycodes.Empty)
	if err != nil {
		return fmt.Errorf("error unpressing key: %w", err)
	}

	fmt.Printf("Sleeping %v\n", waitTime)
	time.Sleep(waitTime) // Sleep for a while to simulate key press duration
	fmt.Printf("Wrote char: %c, 0x%X\n", char, keycode)

	return nil
}

func writeSpecialKey(key string, hid *os.File) {
	var modKey byte = keycodes.KEYCODE_NONE

	keys := make([]byte, 0)

	if strings.Contains(key, "+") {
		parts := strings.Split(key, " + ")

		for _, part := range parts {
			fmt.Println("Processing part:", part)
			fmt.Println("Len:", len(part))
			if k, ok := keycodes.ModifierKey[part]; ok {
				modKey |= k
			} else if k, ok := keycodes.SpecialKey[part]; ok {
				keys = append(keys, k)
			} else if len(part) == 1 {
				// Si es una tecla normal, se procesa como un rune
				fmt.Println("Processing rune:", part)
				r := rune(part[0])
				if k, ok := keycodes.Key[r]; ok {
					fmt.Println("Processing rune without shift:", part)
					keys = append(keys, k)
				} else if k, ok := keycodes.KeyShift[r]; ok {
					fmt.Println("Processing rune with shift:", part)
					modKey |= keycodes.KEYCODE_LEFT_SHIFT
					keys = append(keys, k)
				}
			} else if strings.HasPrefix(part, ".") {
				// Si es un número, se procesa como un Keypad
				fmt.Println("Processing number:", part)
				r := rune(part[1])
				if k, ok := keycodes.Keypad[r]; ok {
					keys = append(keys, k)
				} else {
					fmt.Println("Key not found in Keypad:", part)
				}
			} else {
				fmt.Println("Key not found:", part)
			}

			if len(keys) >= 6 {
				fmt.Println("Limiting to 6 keys")
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
	fmt.Println("Keys:", keys)
	for len(keys) < 6 {
		keys = append(keys, keycodes.KEYCODE_NONE)
	}
	fmt.Println("Keys:", keys)

	_, err := hid.Write([]byte{modKey, 0x00, keys[0], keys[1], keys[2], keys[3], keys[4], keys[5]})
	if err != nil {
		fmt.Printf("Error writing keypress: %v\n", err)
		return
	}
	time.Sleep(waitTimeKey)

	_, err = hid.Write(keycodes.Empty)
	if err != nil {
		fmt.Printf("Error unpressing key: %v\n", err)
		return
	}

	fmt.Printf("Sleeping %v\n", waitTime)
	time.Sleep(waitTime)
	fmt.Printf("Wrote key: %s, %x, %v\n", key, modKey, keys)
}

func main() {
	filePath := flag.String("script", "", "script file to use")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Use: ./ghost-keyboard -script=file")
		os.Exit(1)
	}

	file, err := readFile(*filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	err = processFile(*scanner)
	if err != nil {
		fmt.Println("Error processing file:", err)
	}
}
