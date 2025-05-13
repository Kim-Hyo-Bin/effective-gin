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

	// Graceful shutdown 설정
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// 서비스 시작
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			mainLogger.Errorf("server listen and serve error: %v", err)
		}
	}()

	// 시그널 대기
	<-quit
	mainLogger.Info("Shutting down server...")

	// 5초 타임아웃 설정 (또는 설정 파일에서 가져오기)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Graceful shutdown 시작
	if err := srv.Shutdown(ctx); err != nil {
		mainLogger.Errorf("server forced to shutdown: %v", err)
	}

	mainLogger.Info("Server exiting")
}

// setupServer 함수 (추가 또는 확인 필요)
func setupServer(handler http.Handler, cfg *configs.Config) *http.Server {
	// TODO: cfg를 사용하여 서버 주소, 타임아웃 등을 설정
	serverAddress := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port) // 예시
	return &http.Server{
		Addr:    serverAddress,
		Handler: handler,
		// ReadTimeout:    10 * time.Second, // 예시
		// WriteTimeout:   10 * time.Second, // 예시
		// IdleTimeout:    120 * time.Second, // 예시
	}
}
