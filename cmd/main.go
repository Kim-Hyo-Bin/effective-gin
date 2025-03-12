package main

import (
	"effective-gin/api"
	"effective-gin/configs"
	"effective-gin/utils"
	"fmt"
	"log"
)

func main() {
	cfg := utils.Must(configs.LoadConfig(configs.ConfigFilePath))
	r := api.InitRouter()
	if err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
