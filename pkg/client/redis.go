package client

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var ctx = context.Background()

type RedisClient struct {
	Clt *redis.Client
}

func NewRedisClient() *RedisClient {
	clt := redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
	})
	return &RedisClient{Clt: clt}
}

func (p *RedisClient) Get(key string) string {
	return p.Clt.Get(ctx, key).String()
}

func (p *RedisClient) Put(key string, value string) {
	status := p.Clt.SetEX(ctx, key, value, time.Second * 1)
	if status.Err() != nil {
		println("Put into redis failed, e=", status.Err().Error())
	}
}
