package main

import (
	"github.com/go-git/go-git/v5"
	"gopkg.in/yaml.v3"
)

type Slack struct {
	ApiKey string `yaml:"apiKey"`
}

type Notify struct {
	Slack Slack
}

type Script struct {
	Before string
	After  string
}

type Config struct {
	Path        string
	Schedule    string
	Script      Script
	Notify      Notify
	DisablePull bool            `yaml:"disablePull"`
	PullOptions git.PullOptions `yaml:"pullOptions"`
}

func ReadConfig(configToml []byte) (*map[string]Config, error) {
	config := make(map[string]Config)

	err := yaml.Unmarshal(configToml, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
