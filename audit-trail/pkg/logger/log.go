package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
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
	logFile *os.File
	logger  *log.Logger
	console *log.Logger
}

func NewRFLogger() (*SimpleLogger, error) {
	logDir := "/opt/ram-freezer/bin/logs"

	today := time.Now().In(time.FixedZone("America/Montevideo", -3*60*60)).Format("2006-01-02")
	logFilePath := fmt.Sprintf("%s/%s.log", logDir, today)

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	return &SimpleLogger{
		logFile: file,
		logger:  log.New(file, "", 0),
		console: log.New(os.Stdout, "", 0),
	}, nil
}

func (l *SimpleLogger) Close() {
	if l.logFile != nil {
		l.logFile.Close()
	}
}

func (l *SimpleLogger) Log(level LogLevel, message string) {
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
	l.logger.Println(logOutput)
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
