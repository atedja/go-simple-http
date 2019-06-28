package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Here you can add more additional configuration from the environment variables.
var Port string
var Timeout time.Duration

func init() {
	// Read config defaults.
	viper.SetConfigFile("config/defaults.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Unable to read config file")
	}

	// Viper is using dots (.) as separators.
	// So, we convert them to underscores to be compatible with env vars.
	for _, k := range viper.AllKeys() {
		e := strings.ToUpper(strings.Replace(k, ".", "_", -1))
		viper.BindEnv(k, e)
	}

	Port = fmt.Sprintf(":%d", viper.GetInt("app.port"))
	Timeout = time.Duration(viper.GetInt("app.timeout")) * time.Second
}
