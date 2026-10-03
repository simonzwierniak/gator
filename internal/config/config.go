package config

import (
	"path/filepath"
	"os"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
			return "", err
		}
	path := filepath.Join(homeDir, configFileName)
	return path, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	homePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(homePath, []byte(data), 0666)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	homePath, err := getConfigFilePath()
	if err != nil {
			return Config{}, err
		}
	data, err := os.ReadFile(homePath)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}
	return config, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username 
	err := write(*cfg)
	if err != nil {
		return err
	}
	return nil
}