package pipeline

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/sunzhenkai/db-mark/pkg/client"
	"github.com/sunzhenkai/db-mark/pkg/config"
	"github.com/sunzhenkai/db-mark/pkg/data"
	"github.com/sunzhenkai/db-mark/pkg/generator"
	"github.com/sunzhenkai/db-mark/pkg/util"
	"log"
	"os"
	"time"
)

// Pipeline for mark
type Pipeline struct {
	Dt  data.Data
	Cfg config.Config
	Clt client.Client
}

func NewPipeline() *Pipeline {
	return &Pipeline{Cfg: config.GenConfig()}
}

func (p *Pipeline) prepare() {
	p.Cfg = config.GenConfig()
	p.Dt = *generator.GenData(p.Cfg)
	p.Clt = client.NewRedisClient()
}

func (p *Pipeline) process() {
	tm := metrics.NewTimer()
	_ = metrics.Register("mark", tm)
	go metrics.Log(metrics.DefaultRegistry, time.Second*1, log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	loop := p.Cfg.GetLoopNum()
	for i := 0; i < loop; i++ {
		s := len(p.Dt.StringList)
		for j := 0; j < s; j++ {
			start := time.Now()
			p.Clt.Put(p.Dt.KeyList[j], p.Dt.StringList[j])
			tm.Update(time.Since(start))
		}
	}
	fmt.Printf("%v\n", util.ToIntArray(tm.Percentiles([]float64{0.1, 0.5, 0.99, 0.999}), 1))
}

func (p *Pipeline) quit() {
	println("pipeline is done")
}

func (p *Pipeline) Do() {
	p.prepare()
	p.process()
	p.quit()
}
