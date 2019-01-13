package local

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/tkhr0/bananaci-prototype/lib"
	"github.com/tkhr0/bananaci-prototype/lib/build/provider"
	"github.com/tkhr0/bananaci-prototype/lib/config"
)

const containerRegistryPath string = "localhost"

type LocalBuilder struct {
	*provider.BuildProvider
	lib.Repository
}

func NewLocalBuilder(repository lib.Repository, cloneDir string) (*LocalBuilder, error) {
	_, err := exec.LookPath("git")
	if err != nil {
		return nil, fmt.Errorf("cannot find git: %s", err)
	}

	return &LocalBuilder{
		Repository:    repository,
		BuildProvider: provider.NewBuildProvider(cloneDir),
	}, nil
}

// Build methods is build image
func (l *LocalBuilder) Build() error {
	out, err := clone(l.Repository.CloneURL, l.CloneDir)

	if err != nil {
		return fmt.Errorf("clone failed: %s", err)
	}
	fmt.Printf("%d\n", out)

	ok, err := l.loadConfig()
	if !ok {
		return fmt.Errorf("failed to load config: %s", err)
	}

	prjRoot := filepath.Join(l.CloneDir, l.Repository.Name)

	for _, svc := range l.Config.Services {
		name := filepath.Join(containerRegistryPath,
			l.Repository.FullName(),
			svc.Name,
		)
		tag := fmt.Sprintf("%s:%s", name, "last") // TODO: tag

		fmt.Printf("Run: docker build -f %s -t %s %s\n",
			filepath.Join(prjRoot, svc.Dockerfile),
			tag,
			filepath.Join(prjRoot, svc.Context))

		out, err := exec.
			Command("docker", "build",
				"-f", filepath.Join(prjRoot, svc.Dockerfile),
				"-t", tag,
				filepath.Join(prjRoot, svc.Context)).
			CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to build %s image: %s", name, err)
		}

		fmt.Printf("%s\n", out)
	}

	return nil
}

// loadConfig is
func (l *LocalBuilder) loadConfig() (bool, error) {
	fmt.Println(filepath.Join(l.CloneDir, l.Name, config.FilePath))

	c, err := config.UnmarshalFromFile(
		filepath.Join(l.CloneDir, l.Name, config.FilePath),
	)

	if err != nil {
		return false, fmt.Errorf("faild to load config: %s", err)
	}

	l.Config = *c
	return true, nil
}

func clone(cloneURL string, cloneDir string) ([]byte, error) {
	out, err := exec.
		Command("git", "clone",
			"--depth", "1",
			"--single-branch",
			cloneURL, cloneDir).
		Output()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return out, nil
}
