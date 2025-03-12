package handlers

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	//tmp
	Version     = "dev"
	BuildCommit = "unknown"
	BuildDate   = "unknown"
)
var startTime time.Time

func init() {
	startTime = time.Now()
}

func getUptime() string {
	uptime := time.Since(startTime)
	return fmt.Sprintf("%v", uptime.Round(time.Second))
}

func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

func VersionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version":      Version,
		"build_commit": BuildCommit,
		"build_date":   BuildDate,
		"go_version":   runtime.Version(),
	})
}

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "testok",
		"version":   "testVersion",
		"uptime":    getUptime(),
		"timestamp": time.Now().Format(time.RFC3339),
		"dependencies": map[string]string{
			"database": "testok",
		},
	})
}

func InfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"go_version":      runtime.Version(),
		"uptime":          getUptime(),
		"hostname":        getHostName(),
		"os":              runtime.GOOS + "/" + runtime.GOARCH,
		"goroutine_count": runtime.NumGoroutine(),
		"cpu_count":       runtime.NumCPU(),
	})
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
