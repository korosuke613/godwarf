package main

import (
	"go.uber.org/zap"
	"os/exec"
	"path"
	"strings"
)

type scriptClient struct {
	l      *zap.SugaredLogger
	config *Config
}

func makeScriptClient(config *Config, logger *zap.SugaredLogger) (*scriptClient, error) {
	l := logger.With(
		"path", config.Path,
		"action", "exec script",
	)

	return &scriptClient{
		l:      l,
		config: config,
	}, nil
}

func (sc *scriptClient) beforeExec() {
	if sc.config.Scripts.Before != nil {
		sc.exec("before", sc.config.Scripts.Before)
	}
}

func (sc *scriptClient) afterExec() {
	if sc.config.Scripts.After != nil {
		sc.exec("after", sc.config.Scripts.After)
	}
}

func (sc *scriptClient) exec(name string, script *Script) {
	customLogger := sc.l.With(
		"timing", name)

	customLogger.Info("Start")

	scripts := strings.Split(script.Commands, "\n")
	for _, s := range scripts {
		if s == "" {
			continue
		}
		cmd := exec.Command("sh", "-c", s)
		cmd.Dir = path.Join(sc.config.Path, script.WorkingDirectory)
		output, err := cmd.Output()
		if err != nil {
			customLogger.Errorw("Failed",
				"command", s,
				"reason", err.Error())
			break
		}
		customLogger.Debugw(string(output),
			"command", s)
	}

	customLogger.Info("Finish")
}
