package handlers

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// response models
type VersionResponse struct {
	Version     string `json:"version"`
	BuildCommit string `json:"build_commit"`
	BuildDate   string `json:"build_date"`
	GoVersion   string `json:"go_version"`
}

type HealthResponse struct {
	Message      string                `json:"message"`
	Version      string                `json:"version"`
	Uptime       string                `json:"uptime"`
	Timestamp    string                `json:"timestamp"`
	Dependencies map[string]Dependency `json:"dependencies"`
}
type Dependency struct {
	PackageName string `json:"package_name,omitempty"`
	Version     string `json:"version,omitempty"`
}

type InfoResponse struct {
	GoVersion      string `json:"go_version"`
	Uptime         string `json:"uptime"`
	Hostname       string `json:"hostname"`
	OS             string `json:"os"`
	GoroutineCount int    `json:"goroutine_count"`
	CPUCount       int    `json:"cpu_count"`
}

type PingResponse struct {
	Message string `json:"message"`
}

var (
	Version     = "dev"
	BuildCommit = "unknown"
	BuildDate   = "unknown"
	startTime   time.Time
)

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

// @Summary Get application version information
// @Description Returns the application version, build commit, build date, and Go version
// @Tags info
// @Produce json
// @Success 200 {object} handlers.VersionResponse
// @Router /version [get]
func VersionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, VersionResponse{
		Version:     Version,
		BuildCommit: BuildCommit,
		BuildDate:   BuildDate,
		GoVersion:   runtime.Version(),
	})
}

// @Summary Get application health status
// @Description Returns the application health status, version, uptime, timestamp, and dependencies
// @Tags health
// @Produce json
// @Success 200 {object} handlers.HealthResponse
// @Router /health [get]
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Message:   "testOk",
		Version:   "testVersion",
		Uptime:    getUptime(),
		Timestamp: time.Now().Format(time.RFC3339),
		Dependencies: map[string]Dependency{
			"testDatabaseDependency": {
				PackageName: "testMariaDB",
				Version:     "999.999",
			},
		},
	})
}

// @Summary Get application information
// @Description Returns the application Go version, uptime, hostname, OS, goroutine count, and CPU count
// @Tags info
// @Produce json
// @Success 200 {object} handlers.InfoResponse
// @Router /info [get]
func InfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, InfoResponse{
		GoVersion:      runtime.Version(),
		Uptime:         getUptime(),
		Hostname:       getHostName(),
		OS:             runtime.GOOS + "/" + runtime.GOARCH,
		GoroutineCount: runtime.NumGoroutine(),
		CPUCount:       runtime.NumCPU(),
	})
}

// @Summary Ping endpoint
// @Description Returns pong
// @Tags health
// @Produce json
// @Success 200 {object} handlers.PingResponse
// @Router /ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, PingResponse{
		Message: "pong",
	})
}
