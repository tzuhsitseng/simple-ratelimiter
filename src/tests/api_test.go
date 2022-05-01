package api_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gogotsenghsien/simple-rate-limit/src/api/restful"
	"github.com/gogotsenghsien/simple-rate-limit/src/api/restful/handlers"
	"github.com/gogotsenghsien/simple-rate-limit/src/caches"
	"github.com/gogotsenghsien/simple-rate-limit/src/configs"
	"github.com/gogotsenghsien/simple-rate-limit/src/logs"
	"github.com/ory/dockertest/v3"
)

func TestMain(m *testing.M) {
	// add resource pool
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// add docker resource
	resource, err := pool.Run("redis", "4.0.14", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// add retry func
	if err := pool.Retry(func() error {
		// add config
		config, err := configs.NewConfig()
		if err != nil {
			return err
		}

		// add logger
		logger, err := logs.NewLogger(config)
		if err != nil {
			return err
		}

		// add redis
		redis, err := caches.NewRedis(&caches.RedisConn{
			Host: fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")),
		})
		if err != nil {
			return err
		}

		// add handler
		postHandler := handlers.NewPostHandler(redis, logger, config)

		// start running server
		server := restful.NewServer(config, postHandler)
		go server.Run()
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// purge all resources then exit
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

func TestAddPost(t *testing.T) {
	e := httpexpect.New(t, "http://localhost:8080")

	// check the cases less than or equal to 60
	for i := 1; i <= 60; i++ {
		e.POST("/post").
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusOK).Text().Equal(strconv.FormatInt(int64(i), 10))
	}

	// check the case greater than 60
	e.POST("/post").
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(http.StatusTooManyRequests).Text().Equal("Error")
}
