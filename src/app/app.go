package app

import (
	"github.com/gogotsenghsien/simple-rate-limit/src/api/restful"
	"github.com/gogotsenghsien/simple-rate-limit/src/api/restful/handlers"
	"github.com/gogotsenghsien/simple-rate-limit/src/caches"
	"github.com/gogotsenghsien/simple-rate-limit/src/configs"
	"github.com/gogotsenghsien/simple-rate-limit/src/logs"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(logs.NewLogger)
	_ = container.Provide(configs.NewConfig)
	_ = container.Provide(caches.NewRedisConn)
	_ = container.Provide(caches.NewRedis)
	_ = container.Provide(handlers.NewPostHandler)
	_ = container.Provide(restful.NewServer)
	return container
}
