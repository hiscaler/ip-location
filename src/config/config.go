package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Port int
}

func NewConfig() *Config {
	cfg := &Config{
		Port: 8080,
	}

	if file, err := os.Open("./config/config.json"); err == nil {
		defer file.Close()
		jsonByte, err := ioutil.ReadAll(file)
		if err == nil {
			json.Unmarshal(jsonByte, &cfg)
		}
	}

	return cfg
}
