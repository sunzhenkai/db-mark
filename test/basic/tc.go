package basic

import (
	"encoding/json"
	"github.com/sunzhenkai/db-mark/pkg/config"
)

func TestGetConfig() {
	cfg := config.GenConfig()
	println(cfg.GetLoopNum())
	b, err := json.MarshalIndent(cfg, "",  "  ")
	if err == nil {
		println(string(b))
	} else {
		println(err)
	}
}
