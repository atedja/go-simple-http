package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Here you can add more additional configuration from the environment variables.
var Port string
var Timeout time.Duration

func init() {
	viper.SetDefault("APP_PORT", 3000)
	viper.SetDefault("APP_TIMEOUT", 10)
	viper.AutomaticEnv()

	Port = fmt.Sprintf(":%d", viper.GetInt("APP_PORT"))
	Timeout = time.Duration(viper.GetInt("APP_TIMEOUT")) * time.Second
}
