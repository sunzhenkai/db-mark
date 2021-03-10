package generator

import (
	"github.com/sunzhenkai/db-mark/pkg/config"
	"github.com/sunzhenkai/db-mark/pkg/data"
	"github.com/sunzhenkai/db-mark/pkg/rand"
)

func GenString(cfg config.Config) data.Data {
	n := cfg.GetStringSize()
	var dt = data.Data{}

	for i := 0; i < n; i++ {
		append(dt.StringList, rand.RandomString(cfg.GetDataItemSize()))
	}

	return dt
}
