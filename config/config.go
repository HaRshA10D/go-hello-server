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
	server ServerConfig
}

// Load loads the config from given file
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
		server: getServerConfig(),
	}
}

// Logger returns logging configuration
func Logger() *logger.Config { return appConfig.logger }
// Server returns http server config
func Server() ServerConfig { return appConfig.server }

func getServerConfig() ServerConfig {
	return newServerConfig()
}

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
