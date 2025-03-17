package configs

import (
	"effective-gin/utils"
	"encoding/json"
	"fmt"
	"os"
)

const ConfigFilePath = "./configs/config.json"

type Config struct {
	Server struct {
		Port    string `yaml:"port"`
		Host    string `yaml:"host"`
		LogPath string `yaml:"logPath"`
	} `yaml:"server"`
	GinConfig struct {
		Environment string `yaml:"environment"`
		LogLevel    string `yaml:"logLevel"`
		LogPath     string `yaml:"logPath"`
	} `yaml:"ginConfig"`
	Database struct {
		Dialect  string `yaml:"dialect"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func LoadConfig(configPath string) (config *Config, err error) {
	jsonFile := utils.Must(os.ReadFile(configPath))
	if err := json.Unmarshal(jsonFile, &config); err != nil {
		return config, fmt.Errorf("JSON 파싱 오류: %w", err)
	}
	return
}
