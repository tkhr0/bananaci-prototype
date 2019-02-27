package local

import (
	"fmt"
	// "io/ioutil"
	"os"
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

func NewLocalBuilder(repository lib.Repository) (*LocalBuilder, error) {
	_, err := exec.LookPath("git")
	if err != nil {
		return nil, fmt.Errorf("cannot find git: %s", err)
	}

	// cloneDir, err := ioutil.TempDir("", "bananaci")
	//
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot make temporary dir: %s", err)
	// }

	cloneDir := "/tmp"

	return &LocalBuilder{
		Repository:    repository,
		BuildProvider: provider.NewBuildProvider(cloneDir),
	}, nil
}

// Build methods is build image
func (l *LocalBuilder) Build() error {
	prjRoot := filepath.Join(l.CloneDir, l.Repository.Name)

	if ok, _ := dirExists(prjRoot); !ok {
		out, err := clone(l.Repository.CloneURL, prjRoot)

		fmt.Printf("%s\n", out)
		if err != nil {
			return fmt.Errorf("clone failed: %s", err)
		}
	} else {
		fmt.Println("Skip clone")
	}

	ok, err := l.loadConfig()
	if !ok {
		return fmt.Errorf("failed to load config: %s", err)
	}
	fmt.Printf("loaded config: %+v\n", l.Config.Services)

	for _, image := range l.Config.Images {
		name := filepath.Join(containerRegistryPath,
			l.Repository.FullName(),
			image.Name,
		)
		tag := fmt.Sprintf("%s:%s", name, "latest") // TODO: tag

		fmt.Printf("Run: docker build -f %s -t %s %s\n",
			filepath.Join(prjRoot, image.Dockerfile),
			tag,
			filepath.Join(prjRoot, image.Context))

		out, err := exec.
			Command("docker", "build",
				"-f", filepath.Join(prjRoot, image.Dockerfile),
				"-t", tag,
				filepath.Join(prjRoot, image.Context)).
			CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to build %s image: %s", name, err)
		}

		fmt.Printf("%s\n", out)
	}

	return nil
}

func (l *LocalBuilder) Up() error {
	const chartDir string = "templates/helm/rails"

	prjRoot := filepath.Join(l.CloneDir, l.Repository.Name)

	tmpDir := filepath.Join(prjRoot, config.TmpDir)

	os.Mkdir(tmpDir, 0777) // TODO

	// define value file
	valueFiles := []string{filepath.Join(chartDir, "values.yaml")}

	overridePath := filepath.Join(prjRoot, config.Directory)
	overrideFile := filepath.Join(overridePath, "override.yaml")

	if _, err := os.Stat(overrideFile); err == nil {
		valueFiles = append(valueFiles, overrideFile)
	}

	var valueFilesOption []string
	for _, valueFile := range valueFiles {
		valueFilesOption = append(valueFilesOption, "--values", valueFile)
	}

	installArgs := []string{"install", chartDir}
	installArgs = append(installArgs, valueFilesOption...)
	out, err := exec.
		Command("helm", installArgs...).
		CombinedOutput()

	fmt.Printf("%s\n", out)

	if err != nil {
		return fmt.Errorf("failed to install: %s", err)
	}

	return nil
}

// loadConfig is
func (l *LocalBuilder) loadConfig() (bool, error) {
	configFilePath := filepath.Join(l.CloneDir, l.Repository.Name, config.FilePath)
	fmt.Println(configFilePath)

	c, err := config.UnmarshalFromFile(configFilePath)
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
		CombinedOutput()

	if err != nil {
		return out, fmt.Errorf("clone: %s", err)
	}

	return out, nil
}

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
