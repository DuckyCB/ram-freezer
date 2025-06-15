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

const RFC3339Micro = "2006-01-02T15:04:05.000000Z"

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

type RFLogger struct {
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

func NewRFLogger() (*RFLogger, error) {
	logFilePath := getOutPath()

	return &RFLogger{
		logFilePath: logFilePath,
		console:     log.New(os.Stdout, "", 0),
	}, nil
}

func (l *RFLogger) Log(level LogLevel, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, file, line, ok := runtime.Caller(1)
	if file == "log.go" {
		_, file, line, ok = runtime.Caller(2)
	}
	source := "unknown"
	if ok {
		parts := strings.Split(file, "/")
		if len(parts) > 0 {
			source = parts[len(parts)-1]
		}
	}

	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(RFC3339Micro),
		Level:     level,
		Message:   message,
		Source:    source,
		Line:      line,
	}

	jsonBytes, err := json.Marshal(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LOG_ERROR: Error marshaling log entry: %v\n", err)
		return
	}

	logOutput := string(jsonBytes)

	logFile, err := os.OpenFile(l.logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LOG_ERROR: Error abriendo archivo de log %s: %v", l.logFilePath, err)
	}
	defer logFile.Close()

	if _, err := logFile.WriteString(logOutput + "\n"); err != nil {
		fmt.Fprintf(os.Stderr, "LOG_ERROR: Error escribiendo en archivo de log: %v", err)
	}

	l.console.Println(fmt.Sprintf("%s: %s", level, message))
}

func (l *RFLogger) Debug(message string) {
	l.Log(Debug, message)
}

func (l *RFLogger) Info(message string) {
	l.Log(Info, message)
}

func (l *RFLogger) Warn(message string) {
	l.Log(Warn, message)
}

func (l *RFLogger) Error(message string) {
	l.Log(Error, message)
}

func (l *RFLogger) Fatal(message string) {
	l.Log(Fatal, message)
}
