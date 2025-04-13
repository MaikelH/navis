package internal

import (
	"reflect"
	"testing"
)

func TestReadDeploymentFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *DeploymentConfig
		wantErr bool
	}{
		{
			name: "Read Simple file",
			args: args{
				filePath: "./test/test-0.yaml",
			},
			want: &DeploymentConfig{
				Deployments: []Deployment{
					{
						Environment: "dev",
						ComposeFile: "docker.compose.yml",
						Auth:        nil,
						Targets: []Target{
							{
								Address: "127.0.0.1",
							},
						},
						EnvVariables: nil,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Read File with env variables",
			args: args{
				filePath: "./test/test-1.yaml",
			},
			want: &DeploymentConfig{
				Deployments: []Deployment{
					{
						Environment: "dev",
						ComposeFile: "docker.compose.yml",
						Auth:        &Auth{Key: "test"},
						Targets: []Target{
							{
								Address: "127.0.0.1",
							},
						},
						EnvVariables: map[string]string{
							"TEST_VARIABLE": "test",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDeploymentFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDeploymentFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDeploymentFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
