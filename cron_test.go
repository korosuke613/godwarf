package main

import (
	"github.com/robfig/cron/v3"
	"reflect"
	"testing"
)

func Test_makeCron(t *testing.T) {
	type args struct {
		configs *map[string]Config
	}
	tests := []struct {
		name    string
		args    args
		want    *[]cron.Entry
		wantErr bool
	}{
		{
			name: "setting cron",
			args: args{
				configs: &map[string]Config{
					"renovate": {
						Path: "./hoge",
					},
					"renovate2": {
						Path:     "./foo",
						Schedule: "0 * * * *",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeCron(tt.args.configs)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeCron() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeCron() got = %v, want %v", entries, tt.want)
			}
		})
	}
}
