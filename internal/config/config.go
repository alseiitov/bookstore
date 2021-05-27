package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type (
	Config struct {
		API      API      `json:"api"`
		Database Database `json:"database"`
	}

	API struct {
		Port int `json:"port"`
	}

	Database struct {
		User     string `json:"user"`
		Password string
		Host     string `json:"host"`
		Port     int    `json:"port"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	var cfg Config

	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}

	cfg.Database.Password = os.Getenv("DB_PASSWORD")

	return &cfg, nil
}
