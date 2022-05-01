package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gogotsenghsien/simple-rate-limit/src/caches"
	"github.com/gogotsenghsien/simple-rate-limit/src/configs"
	"github.com/gogotsenghsien/simple-rate-limit/src/logs"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	redis  *caches.Redis
	logger *logs.Logger
	config *configs.Config
}

func NewPostHandler(redis *caches.Redis, logger *logs.Logger, config *configs.Config) *PostHandler {
	return &PostHandler{
		redis:  redis,
		logger: logger,
		config: config,
	}
}

func (h *PostHandler) AddPost(c echo.Context) error {
	// get client ip address
	ipExtractor := echo.ExtractIPFromXFFHeader()
	ip := ipExtractor(c.Request())
	h.logger.Info("Try to add post", logs.FieldIP, ip)

	// get requested nums per minute by redis
	pipe := h.redis.TxPipeline()
	key := fmt.Sprintf("ReqLimit:%s:%s", time.Now().Format("15-04"), ip)
	increased := pipe.Incr(key)
	_ = pipe.Expire(key, time.Minute)
	if _, err := pipe.Exec(); err != nil {
		h.logger.Error("Get request limit failed", logs.FieldError, err)
		return c.String(http.StatusInternalServerError, "Get request limit failed")
	}

	// check requested nums whether greater than limit
	cnt := increased.Val()
	if cnt > h.config.GetInt64("limits.request") {
		return c.String(http.StatusTooManyRequests, "Error")
	}
	return c.String(http.StatusOK, strconv.FormatInt(cnt, 10))
}
