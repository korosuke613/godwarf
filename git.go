package main

import "github.com/go-git/go-git/v5"

func makeWorktree(config *Config) (*git.Worktree, error) {
	r, err := git.PlainOpen(config.Path)
	if err != nil {
		return nil, err
	}

	w, err := r.Worktree()
	if err != nil {
		return nil, err
	}

	return w, nil
}

func pull(config *Config) error {
	w, err := makeWorktree(config)
	if err != nil {
		return err
	}

	err = w.Pull(&config.PullOptions)
	if err != nil {
		if err.Error() == "already up-to-date" {
			return nil
		}
		return err
	}

	return nil
}
