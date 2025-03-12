package api

import (
	"effective-gin/configs"
	"effective-gin/internal/handlers"
	"effective-gin/utils"
	"effective-gin/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
)

var cfg = utils.Must(configs.LoadConfig("./configs/config.json"))

func InitRouter() *gin.Engine {
	if cfg.GinConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	//TODO: scope설정은 추후 config로 변경 또는 상황에 따라 변경하고 에러처리 확인
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Use(GinLogger())
	r.Use(gin.Recovery())

	rootGroup := r.Group("/")
	rootGroup.GET("", handlers.VersionHandler)
	rootGroup.GET("version", handlers.VersionHandler)
	rootGroup.GET("health", handlers.HealthHandler)
	rootGroup.GET("info", handlers.InfoHandler)
	rootGroup.GET("ping", handlers.PingHandler)

	return r
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		GinLogger := logger.NewLogger(cfg.GinConfig.LogPath)
		requestLogger := GinLogger.WithFields(logger.Fields{
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"clientIP": c.ClientIP(),
		})
		c.Set("logger", requestLogger)
		c.Next()
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()

		requestLogger.WithFields(logger.Fields{
			"latency": latency,
			"status":  statusCode,
		}).Infof("Request handled")
	}
}
