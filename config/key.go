package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func mustGetInt(key string) int {
	mustHave(key)
	return viper.GetInt(key)
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
