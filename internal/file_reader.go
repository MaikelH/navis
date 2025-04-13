package internal

import (
	"os"

	"github.com/goccy/go-yaml"
)

// ReadDeploymentFile reads and parses a YAML deployment configuration file.
// It takes a file path as input and returns a pointer to a DeploymentConfig struct
// along with any error encountered during reading or unmarshaling.
func ReadDeploymentFile(filePath string) (*DeploymentConfig, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var deploymentConfig DeploymentConfig
	err = yaml.Unmarshal(yamlFile, &deploymentConfig)
	if err != nil {
		return nil, err
	}

	return &deploymentConfig, nil
}
