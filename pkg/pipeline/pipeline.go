package pipeline

import (
	"github.com/sunzhenkai/db-mark/pkg/client"
	"github.com/sunzhenkai/db-mark/pkg/config"
	"github.com/sunzhenkai/db-mark/pkg/data"
)

// Pipeline for mark
type Pipeline struct {
	data   data.Data
	config config.Config
	client client.Client
}

func (p Pipeline) prepare() {
	// prepare
}

func (p Pipeline) process() {
	// process
}

func (p Pipeline) quit() {
	// quit
}
