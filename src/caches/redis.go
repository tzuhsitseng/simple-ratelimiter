package caches

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	*redis.Client
}

func NewRedis(conn *RedisConn) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conn.Host,
		Password: conn.Password,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return &Redis{client}, nil
}
