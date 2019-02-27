package provider

import (
	"github.com/tkhr0/bananaci-prototype/lib/config"
)

type BuildProvider struct {
	CloneDir string // git clone Dir
	config.Config
}

func NewBuildProvider(cloneDir string) *BuildProvider {
	return &BuildProvider{
		CloneDir: cloneDir,
	}
}
