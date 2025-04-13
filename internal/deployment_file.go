package internal

// Config represents the root structure of the YAML file.
type DeploymentConfig struct {
	Deployments []Deployment `yaml:"deployments"`
}

// Deployment represents a single deployment configuration block.
type Deployment struct {
	Environment  string            `yaml:"environment"`
	ComposeFile  string            `yaml:"compose_file"`
	Auth         *Auth             `yaml:"auth"`
	Targets      []Target          `yaml:"targets"`
	EnvVariables map[string]string `yaml:"env_variables"` // Represents the list of key-value pairs
}

// Auth holds the authentication details.
type Auth struct {
	Key string `yaml:"key"`
}

// Target represents a deployment target address.
type Target struct {
	Address string `yaml:"address"`
	Auth    *Auth  `yaml:"auth"`
}
