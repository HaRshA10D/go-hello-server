package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	*zap.SugaredLogger
}

func newZapLogger(config Config) Logger {
	zapEncoderConfig := zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey: "level",
		TimeKey: "time",
	}
	zapEncoder := zapcore.NewConsoleEncoder(zapEncoderConfig)
	zapCore := zapcore.NewCore(zapEncoder, zapcore.AddSync(config.Output), zapcore.InfoLevel)

	l := zap.New(zapCore).Sugar()
	defer l.Sync()

	return &zapLogger{l}
}

func (l *zapLogger) Debug(msg string, args ...interface{}) {
	l.Debugf(msg, args...)
}

func (l *zapLogger) Info(msg string, args ...interface{}) {
	l.Infof(msg, args...)
}

func (l *zapLogger) Warn(msg string, args ...interface{}) {
	l.Warnf(msg, args...)
}

func (l *zapLogger) Error(msg string, args ...interface{}) {
	l.Errorf(msg, args...)
}

func (l *zapLogger) Fatal(msg string, args ...interface{}) {
	l.Fatalf(msg, args...)
}

func (l *zapLogger) WithFields(fields map[string]interface{}) Logger {
	_fields := make([]interface{}, 0)
	for k, v := range fields {
		_fields = append(_fields, k)
		_fields = append(_fields, v)
	}
	return &zapLogger{l.With(_fields...)}
}
