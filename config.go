package main

import (
	"github.com/go-git/go-git/v5"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Slack struct {
	ApiKey string `yaml:"apiKey"`
}

type Notify struct {
	Slack Slack
}

type Script struct {
	WorkingDirectory string `yaml:"workingDirectory"`
	Commands         string
}

type Scripts struct {
	Before *Script
	After  *Script
}

type Config struct {
	Path        string
	Schedule    string
	Scripts     Scripts
	Notify      Notify
	DisablePull bool            `yaml:"disablePull"`
	PullOptions git.PullOptions `yaml:"pullOptions"`
}

type Configs map[string]Config

func readConfig(filePath string, logger *zap.SugaredLogger) (*Configs, error) {
	l := logger.With(
		"action", "read configs",
		"file", filePath,
	)

	l.Info("Start")
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	configs, err := parseConfigs(bytes)
	l.Info("Finish")

	return configs, err
}

func parseConfigs(configToml []byte) (*Configs, error) {
	configs := make(map[string]Config)

	err := yaml.Unmarshal(configToml, &configs)
	if err != nil {
		return nil, err
	}

	return (*Configs)(&configs), nil
}
