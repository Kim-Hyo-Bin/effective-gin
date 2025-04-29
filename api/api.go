package api

import (
	"effective-gin/api/handlers"
	"effective-gin/configs"
	"effective-gin/utils"
	"effective-gin/utils/logger"
	"net/http"
	"strings"
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
	r.Use(redirectAllToV2)

	rootGroup := r.Group("")
	InitV1Routes(rootGroup)
	InitV2Routes(rootGroup)

	return r
}

func InitV1Routes(RouterGroup *gin.RouterGroup) {
	v1Group := RouterGroup.Group("/v1")
	v1Group.GET("", handlers.VersionHandler)
	v1Group.GET("version", handlers.VersionHandler)
	v1Group.GET("health", handlers.HealthHandler)
	v1Group.GET("info", handlers.InfoHandler)
	v1Group.GET("ping", handlers.PingHandler)
}

func InitV2Routes(RouterGroup *gin.RouterGroup) {
	v2Group := RouterGroup.Group("/v2")
	v2Group.GET("", handlers.VersionHandler)
	v2Group.GET("version", handlers.VersionHandler)
	v2Group.GET("health", handlers.HealthHandler)
	v2Group.GET("info", handlers.InfoHandler)
	v2Group.GET("ping", handlers.PingHandler)
	v2Group.GET("example1", handlers.TestShapeCalculations1)
	v2Group.GET("example2", handlers.TestShapeCalculations2)
}

func redirectAllToV2(c *gin.Context) {
	if c.Request.URL.Path != "/v1" && !strings.HasPrefix(c.Request.URL.Path, "/v1/") &&
		c.Request.URL.Path != "/v2" && !strings.HasPrefix(c.Request.URL.Path, "/v2/") &&
		c.Request.URL.Path != "/swagger" && !strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
		c.Redirect(http.StatusMovedPermanently, "/v2"+c.Request.URL.Path)
		c.Abort()
	}
	c.Next()
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
