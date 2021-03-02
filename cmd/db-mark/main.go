package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/sunzhenkai/db-mark/pkg/config"
	"github.com/sunzhenkai/db-mark/test/basic"
)

func main() {
	// config
	config.SetDefault()
	config.ReadConfig("configs")
	fmt.Printf("db.type \t: %s \n", viper.Get("db.type"))
	fmt.Printf("db.max.qps \t: %d \n", viper.GetInt("db.max.qps"))

	// test
	basic.DoTest()
}
