package main

import (
	"os"
	"os/exec"
	"testing"
)

func Test_pull(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name: "git pullする",
		args: args{
			config: &Config{
				Path: "/tmp/godwarf/git_pull_test",
			},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.MkdirAll("/tmp/godwarf/git_pull_test", 0755)
			if err != nil {
				panic(err)
			}
			cmd := exec.Command("git", "clone", "--depth=2", "https://github.com/korosuke613/godwarf", "/tmp/godwarf/git_pull_test")
			err = cmd.Run()
			if err != nil {
				panic(err)
			}
			defer func() {
				err := os.RemoveAll("/tmp/godwarf/git_pull_test")
				if err != nil {
					panic(err)
				}
			}()

			if err := pull(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("pull() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
