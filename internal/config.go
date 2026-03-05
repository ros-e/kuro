package internal

type Provider string
type Trigger string

const (
	ProviderGitHub Provider = "github"
	ProviderGitLab Provider = "gitlab"
	ProviderGitea  Provider = "gitea"
)

const (
	TriggerPush   Trigger = "push"
	TriggerManual Trigger = "tag"
)

type RootConfig struct {
	Version  string                  `yaml:"version"`
	Projects map[string]ProjectEntry `yaml:"projects"`
}

type ProjectEntry struct {
	Name string `yaml:"name"`
	Dir  string `yaml:"dir"`
}

type ProjectConfig struct {
	Services map[string]ServiceEntry `yaml:"services"`
}

type ServiceEntry struct {
	Name string `yaml:"name"`
	Dir  string `yaml:"dir"`
}

type ServiceConfig struct {
	Name       string   `yaml:"name"`
	Provider   Provider `yaml:"provider"`
	Trigger    Trigger  `yaml:"trigger"`
	AutoDeploy bool     `yaml:"auto_deploy"`
}
