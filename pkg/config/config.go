package config

import "github.com/spf13/viper"

// Config config for mark
type Config interface {
	// GetStringSize generator
	GetStringSize() int
	GetListSize() int
	GetMapSize() int

	// GetLoopNum benchmark
	GetLoopNum() int
	GetDataItemSize() int
}

// Impl ConfigImpl config impl
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

func GenConfig() Config {
	var stringSize = viper.GetInt("bench.data.string-size")
	var stringListSize = viper.GetInt("bench.data.string-list-size")
	var loop = viper.GetInt("bench.loop")
	var cfg Config = Impl{stringSize: stringSize, loopNum: loop, listSize: stringListSize}
	return cfg
}
