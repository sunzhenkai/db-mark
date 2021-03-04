package config

// Config for mark
type Config interface {
	GetStringSize() int32
	GetListSize() int32
	GetMapSize() int32
	GetLoopNum() int32
}
