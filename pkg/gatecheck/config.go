package gatecheck

import (
	"github.com/gatecheckdev/gatecheck/pkg/artifact/grype"
	"github.com/gatecheckdev/gatecheck/pkg/artifact/semgrep"
)

type Config struct {
	ProjectName string         `yaml:"projectName" json:"projectName"`
	Grype       grype.Config   `yaml:"grype,omitempty" json:"grype,omitempty"`
	Semgrep     semgrep.Config `yaml:"semgrep,omitempty" json:"semgrep,omitempty"`
}

func NewConfig(projectName string) *Config {
	return &Config{
		ProjectName: projectName,
		Grype:       *grype.NewConfig(-1),
		Semgrep:     *semgrep.NewConfig(-1),
	}
}