package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// SetDefault : set default configs
func SetDefault() {
	viper.SetDefault("db.max.qps", 1000)
}

// ReadConfig : read configs from config file
func ReadConfig(pt string) {
	viper.SetConfigName("db")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(pt)
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error occurred while read config : %s", err))
	}
}
