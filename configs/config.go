package configs

import (
	"effective-gin/utils"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	App struct {
		Environment string `yaml:"environment"`
		LogLevel    string `yaml:"log_level"`
	} `yaml:"app"`
}

func LoadConfig(configPath string) (Config, error) {
	jsonFile := utils.Must(os.ReadFile(configPath))

	var config Config

	if err := json.Unmarshal(jsonFile, &config); err != nil {
		return config, fmt.Errorf("JSON 파싱 오류: %w", err)
	}

	return config, nil
}
