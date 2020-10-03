package logger

import (
	"io"
	"os"
)

// A good convention to use
// https://stackoverflow.com/questions/2031163/when-to-use-the-different-log-levels/2031195#2031195
//

type LogLevel uint8

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
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

func GetLogger() Logger {
	return logger
}

func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	logger.Fatal(msg, args...)
}
