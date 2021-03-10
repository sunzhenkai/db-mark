package config

// Config for mark
type Config interface {
	// generator
	GetStringSize() int
	GetListSize() int
	GetMapSize() int

	// benchmark
	GetLoopNum() int
	GetDataItemSize() int
}
