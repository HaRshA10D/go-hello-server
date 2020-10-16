package logger

import (
	"io"
	"os"
)

// A good convention to use
// https://stackoverflow.com/questions/2031163/when-to-use-the-different-log-levels/2031195#2031195
//

// LogLevel defines levels for logger
type LogLevel uint8

const (
	// DebugLevel is the lowest and used to debug issues, not for production
	DebugLevel LogLevel = iota
	// InfoLevel is the default level for all events in app
	InfoLevel
	// WarnLevel can be used for events that might cause errors
	WarnLevel
	// ErrorLevel can be used for errors that needs investigation
	ErrorLevel
	// FatalLevel can be used critical failure that needs immediate investigation
	FatalLevel
)

// Logger enables caller to log a message at certain level.
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	WithFields(fields map[string]interface{}) Logger
}

// Config holds level of logger, writer to write logs.
type Config struct {
	Level LogLevel
	Output io.Writer
}

var logger Logger

// Init initializes the logger with config
func Init(config Config) {
	hostName, _ := os.Hostname()
	logger = newZapLogger(config).WithFields(map[string]interface{}{
		"host": hostName,
	})
}

// GetLogger returns current app's logger
func GetLogger() Logger {
	return logger
}

// Debug logs the message at debug level
func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}

// Info logs the message at info level
func Info(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

// Warn logs the message at warn level
func Warn(msg string, args ...interface{}) {
	logger.Warn(msg, args...)
}

// Error logs the message at error level
func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}

// Fatal logs the message at fatal level
func Fatal(msg string, args ...interface{}) {
	logger.Fatal(msg, args...)
}
