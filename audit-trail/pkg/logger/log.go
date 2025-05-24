package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type LogLevel string

const (
	Debug LogLevel = "DEBUG"
	Info  LogLevel = "INFO"
	Warn  LogLevel = "WARN"
	Error LogLevel = "ERROR"
	Fatal LogLevel = "FATAL"
)

type LogEntry struct {
	Timestamp string   `json:"timestamp"`
	Level     LogLevel `json:"level"`
	Message   string   `json:"message"`
	Source    string   `json:"source"`
	Line      int      `json:"line"`
}

type SimpleLogger struct {
	logFilePath string
	console     *log.Logger
	mu          sync.Mutex
}

func getOutPath() string {
	content, err := os.ReadFile("/opt/ram-freezer/.out")
	if err != nil {
		fmt.Printf("no se pudo leer la ubicación de salida de logs. usando salida genérica.")
		return "/opt/ram-freezer/bin/ram-freezer.log"
	}
	return fmt.Sprintf("%s/ram-freezer.log", strings.TrimSpace(string(content)))
}

func NewRFLogger() (*SimpleLogger, error) {
	logFilePath := getOutPath()

	return &SimpleLogger{
		logFilePath: logFilePath,
		console:     log.New(os.Stdout, "", 0),
	}, nil
}

func (l *SimpleLogger) Log(level LogLevel, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, file, line, ok := runtime.Caller(1)
	source := "unknown"
	if ok {
		parts := strings.Split(file, "/")
		if len(parts) > 0 {
			source = parts[len(parts)-1]
		}
	}

	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		Level:     level,
		Message:   message,
		Source:    source,
		Line:      line,
	}

	jsonBytes, err := json.Marshal(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling log entry: %v\n", err)
		return
	}

	logOutput := string(jsonBytes)

	logFile, err := os.OpenFile(l.logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error abriendo archivo de log %s: %v\n", l.logFilePath, err)
		l.console.Println(message)
		return
	}
	defer logFile.Close()

	if _, err := logFile.WriteString(logOutput + "\n"); err != nil {
		fmt.Fprintf(os.Stderr, "Error escribiendo en archivo de log: %v\n", err)
	}

	l.console.Println(message)
}

func (l *SimpleLogger) Debug(message string) {
	l.Log(Debug, message)
}

func (l *SimpleLogger) Info(message string) {
	l.Log(Info, message)
}

func (l *SimpleLogger) Warn(message string) {
	l.Log(Warn, message)
}

func (l *SimpleLogger) Error(message string) {
	l.Log(Error, message)
}

func (l *SimpleLogger) Fatal(message string) {
	l.Log(Fatal, message)
}
