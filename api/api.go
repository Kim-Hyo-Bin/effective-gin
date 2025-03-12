package api

import (
	"effective-gin/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	rootGroup := r.Group("/")
	rootGroup.GET("version", handlers.VersionHandler)
	rootGroup.GET("health", handlers.HealthHandler)
	rootGroup.GET("info", handlers.InfoHandler)
	rootGroup.GET("ping", handlers.PingHandler)

	return r
}
