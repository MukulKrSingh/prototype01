package logger

import (
	"log"
	"os"
	"time"
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

// Init initializes the logger
func Init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	warningLogger = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
}

// Info logs an informational message
func Info(message string) {
	if infoLogger == nil {
		Init()
	}
	infoLogger.Println(message)
}

// Warn logs a warning message
func Warn(message string) {
	if warningLogger == nil {
		Init()
	}
	warningLogger.Println(message)
}

// Error logs an error message
func Error(message string, err error) {
	if errorLogger == nil {
		Init()
	}
	if err != nil {
		errorLogger.Printf("%s: %v", message, err)
	} else {
		errorLogger.Println(message)
	}
}

// Fatal logs an error message and exits
func Fatal(message string, err error) {
	if errorLogger == nil {
		Init()
	}
	if err != nil {
		errorLogger.Fatalf("%s: %v", message, err)
	} else {
		errorLogger.Fatal(message)
	}
}

// RequestLogger logs HTTP request details
func RequestLogger(method, path, remoteAddr string, statusCode int, duration time.Duration) {
	if infoLogger == nil {
		Init()
	}
	infoLogger.Printf(
		"method=%s path=%s remote=%s status=%d duration=%s",
		method,
		path,
		remoteAddr,
		statusCode,
		duration,
	)
}
