package main

import (
	"context"
	"effective-gin/api"
	"effective-gin/configs"
	_ "effective-gin/docs"
	"effective-gin/utils"
	"effective-gin/utils/logger"

	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := utils.Must(configs.LoadConfig(configs.ConfigFilePath))
	mainLogger := logger.NewLogger(cfg.Server.LogPath)
	defer mainLogger.Info("Server stopped")
	mainLogger.Info("Starting server...")

	r := api.InitRouter()

	if cfg.GinConfig.Environment == "development" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	srv := setupServer(r, cfg)
	startServer(srv, mainLogger)
}

func setupServer(r *gin.Engine, cfg *configs.Config) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: r,
	}
}

func startServer(srv *http.Server, mainLogger *logrus.Logger) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			mainLogger.Fatalf("Failed to start server: %v", err)
		}
	}()
	waitForShutdown(srv, mainLogger)
}

func waitForShutdown(srv *http.Server, mainLogger *logrus.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	mainLogger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		mainLogger.Fatalf("Server forced to shutdown: %s", err)
	}
}
