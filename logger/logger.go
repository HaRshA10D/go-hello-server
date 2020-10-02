package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

// A good convention to use
// https://stackoverflow.com/questions/2031163/when-to-use-the-different-log-levels/2031195#2031195
//
// Interface to adapt multiple loggers, but going with zap for now
// type Logger interface {
//	 Debug(msg string, args ...interface{})
//	 Info(msg string, args ...interface{})
// }
//
// func GetLogger() *Logger {
//   return logger
// }

type LogLevel uint8

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var logger *zap.SugaredLogger

// LoggerConfig holds level of logger, writer to write logs
type Config struct {
	Level LogLevel
	Output io.Writer
}

// Init initializes the logger with config
func Init(config Config) {
	zapEncoderConfig := zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey: "level",
		TimeKey: "time",
	}
	zapEncoder := zapcore.NewConsoleEncoder(zapEncoderConfig)
	zapCore := zapcore.NewCore(zapEncoder, zapcore.AddSync(config.Output), zapcore.InfoLevel)

	logger = zap.New(zapCore).Sugar()
	defer logger.Sync()
}

func Debug(msg string, args ...interface{}) {
	logger.Debugf(msg, args)
}

func Info(msg string, args ...interface{}) {
	logger.Infof(msg, args)
}

func Warn(msg string, args ...interface{}) {
	logger.Warnf(msg, args)
}

func Error(msg string, args ...interface{}) {
	logger.Errorf(msg, args)
}

func Fatal(msg string, args ...interface{}) {
	logger.Fatalf(msg, args)
}
