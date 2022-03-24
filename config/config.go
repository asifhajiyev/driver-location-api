package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database struct {
		Name    string `yaml:"name"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"database"`
}

func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}
