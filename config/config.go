package config

import (
	"fmt"
	"github.com/HaRshA10D/go-hello-server/logger"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	appConfig config
	isLoaded = false
)

type config struct {
	logger *logger.Config
}

func Load() {
	if isLoaded {
		return
	}

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")

	viper.ReadInConfig()
	viper.AutomaticEnv()

	appConfig = config{
		logger: getLoggerConfig(),
	}
}

// Logger returns logging configuration
func Logger() *logger.Config { return appConfig.logger }

func getLoggerConfig() *logger.Config {
	logLevel := mustGetString("LOG_LEVEL")

	return &logger.Config{
		Level: parseLogLevel(logLevel),
		Output: os.Stdout,
	}
}

func parseLogLevel(lvl string) logger.LogLevel {
	var logLevel logger.LogLevel
	switch strings.ToLower(lvl) {
	case "debug":
		logLevel = logger.DebugLevel
	case "info":
		logLevel = logger.InfoLevel
	case "warn":
		logLevel = logger.WarnLevel
	case "error":
		logLevel = logger.ErrorLevel
	case "fatal":
		logLevel = logger.FatalLevel
	default:
		panic(fmt.Sprintf("Logger level mis-configured: %s", lvl))
	}
	return logLevel
}

func mustGetString(key string) string {
	mustHave(key)
	return viper.GetString(key)
}

func mustHave(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("Config not set for: %s", key))
	}
}
