package main

import (
	"github.com/gogotsenghsien/simple-rate-limit/src/api/restful"
	"github.com/gogotsenghsien/simple-rate-limit/src/app"
)

func main() {
	// prepare all dep.
	container := app.BuildContainer()

	// run server
	if err := container.Invoke(func(server *restful.Server) {
		server.Run()
	}); err != nil {
		panic(err)
	}
}
