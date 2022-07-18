package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"go.uber.org/zap"
)

var (
	alreadyUpToDate = "already up-to-date"
)

type remoteInfo struct {
	Name string
	Urls []string
}

type gitClient struct {
	l          *zap.SugaredLogger
	config     *Config
	repository *git.Repository
	remoteInfo remoteInfo
}

func makeGitClient(config *Config, logger *zap.SugaredLogger) (*gitClient, error) {
	r, err := git.PlainOpen(config.Path)
	if err != nil {
		return nil, err
	}
	remoteInfo, err := getRemoteInfo(r, config.PullOptions.RemoteName)
	if err != nil {
		return nil, err
	}

	l := logger.With(
		"path", config.Path,
		"remoteInfo", remoteInfo,
		"action", "pulling",
	)

	return &gitClient{
		l:          l,
		config:     config,
		repository: r,
		remoteInfo: *remoteInfo,
	}, nil
}

func getRemoteInfo(r *git.Repository, remoteName string) (*remoteInfo, error) {
	if remoteName == "" {
		remoteName = "origin"
	}

	remote, err := r.Remote(remoteName)
	if err != nil {
		return nil, err
	}

	remoteConfig := remote.Config()
	return &remoteInfo{remoteConfig.Name, remoteConfig.URLs}, nil
}

func (gc *gitClient) pull(config *Config) error {
	gc.l.Infow("Start")

	w, err := gc.repository.Worktree()
	if err != nil {
		return err
	}

	err = w.Pull(&config.PullOptions)
	if err != nil {
		if err.Error() == alreadyUpToDate {
			gc.l.Debugw(fmt.Sprintf("Git: %s", alreadyUpToDate))
		} else {
			gc.l.Errorw("Failed", "reason", err.Error())
		}
	}

	gc.l.Infow("Finish")
	return nil
}
