package caches

import (
	"strings"

	"github.com/gogotsenghsien/simple-rate-limit/src/configs"
)

type RedisConn struct {
	Host     string
	Password string
}

func NewRedisConn(config *configs.Config) (*RedisConn, error) {
	uri := config.GetString("redis.url")
	host := ""
	password := ""

	// try compatible with Heroku redis
	// e.g. redis://:pd3a2d4bd85868ce09020641447e3a4a15ca54fe9c4f83b40216ea27ae907fdd4@ec2-52-23-21-53.compute-1.amazonaws.com:26419
	if strings.Contains(uri, "redis://") {
		s := strings.Split(uri, "@")
		host = s[1]
		password = strings.Split(s[0], "redis://:")[1]
	} else {
		host = uri
	}

	return &RedisConn{
		Host:     host,
		Password: password,
	}, nil
}
