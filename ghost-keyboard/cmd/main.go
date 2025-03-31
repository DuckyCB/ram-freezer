package main

import (
	"bufio"
	"controller/cmd/keycodes"
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
)

// wait_time is the time to wait between key presses
var wait_time = 100 * time.Millisecond

// wait_time_key is the time to wait pressing a key
var wait_time_key = 30 * time.Millisecond

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
		fmt.Println("Processing line:", line)
		
		// wait for a specific time
		if strings.HasPrefix(line, "wait") {
			parts := strings.SplitN(line, " ", 2) // Divide en ["wait", "5"]
			
			if len(parts) < 2 {
				fmt.Println("Falta el tiempo después de 'wait': Default 1")
				parts = append(parts, "1") // Si no hay tiempo, usar 1 segundo por defecto
			}

			waitTime, err := strconv.Atoi(parts[1]) // Convierte el tiempo a entero
			if err != nil {
				fmt.Println("Error: tiempo inválido después de 'wait'")
				return fmt.Errorf("error: tiempo inválido después de 'wait'")
			}

			fmt.Printf("Sleeping %d seconds\n", waitTime)
			time.Sleep(time.Duration(waitTime) * time.Second)
			continue // Skip to the next line
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

	time.Sleep(wait_time_key)

	_, err = hid.Write(keycodes.Empty)
	if err != nil {
		return fmt.Errorf("error unpressing key: %w", err)
	}

	fmt.Printf("Sleeping %v\n", wait_time)
	time.Sleep(wait_time) // Sleep for a while to simulate key press duration
	fmt.Printf("Wrote key: %c, 0x%X\n", char, keycode)

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
					r := rune(part[0])
					if k, ok := keycodes.Key[r]; ok {
						keys = append(keys, k)
					} else if k, ok := keycodes.KeyShift[r]; ok {
						modKey |= keycodes.KEYCODE_LEFT_SHIFT
						keys = append(keys, k)
					}
				}

				if len(keys) >= 6 {
					break
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
		fmt.Printf("Error writing keypress: %v\n", err)
		return
	}
	time.Sleep(wait_time_key)

	_, err = hid.Write(keycodes.Empty)
	if err != nil {
		fmt.Printf("Error unpressing key: %v\n", err)
		return
	}

	fmt.Printf("Sleeping %v\n", wait_time)
	time.Sleep(wait_time)
	fmt.Printf("Wrote key: %s, %x, %v\n", key, modKey, keys)
}

func main() {
	filePath := flag.String("f", "", "File to use")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Use: ./ghost-keyboard -f file")
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
