package main

import (
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	type args struct {
		configYaml []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *map[string]Config
		wantErr bool
	}{
		{
			name: "Readable config",
			args: args{
				[]byte(`
renovate:
  path: "/tmp/hoge"
  script: 
    after: | 
      docker compose pull
      docker compose down
      docker compose up
  notify:
    slack:
      apiKey: "api"
`),
			},
			want: &map[string]Config{
				"renovate": {
					Path: "/tmp/hoge",
					Script: Script{
						Before: "",
						After: `docker compose pull
docker compose down
docker compose up
`,
					},
					Notify: Notify{
						Slack: Slack{
							ApiKey: "api",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfig(tt.args.configYaml)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
