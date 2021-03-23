package config

import "github.com/spf13/viper"

// config for mark
type config interface {
	// generator
	GetStringSize() int
	GetListSize() int
	GetMapSize() int

	// benchmark
	GetLoopNum() int
	GetDataItemSize() int
}

// ConfigImpl config impl
type Impl struct {
	stringSize   int
	listSize     int
	mapSize      int
	loopNum      int
	dataItemSize int
}

func (c Impl) GetStringSize() int {
	return c.stringSize
}

func (c Impl) GetListSize() int {
	return c.listSize
}

func (c Impl) GetMapSize() int {
	return c.mapSize
}

func (c Impl) GetLoopNum() int {
	return c.loopNum
}

func (c Impl) GetDataItemSize() int {
	return c.dataItemSize
}

func GenConfig() config {
	var stringSize = viper.GetInt("bench.data.string-size")
	var loop = viper.GetInt("bench.loop")
	var cfg config = Impl{stringSize: stringSize, loopNum: loop}
	return cfg
}
