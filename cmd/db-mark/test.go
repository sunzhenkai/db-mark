package main

import (
	"github.com/sunzhenkai/db-mark/pkg/config"
	"github.com/sunzhenkai/db-mark/test/basic"
)

func main() {
	config.SetDefault()
	config.ReadConfig("configs")

	basic.DoTest()
	basic.TestGetConfig()
}