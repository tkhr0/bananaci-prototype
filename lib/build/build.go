package build

import (
	"fmt"

	"github.com/tkhr0/bananaci-prototype/lib"
	localBuilder "github.com/tkhr0/bananaci-prototype/lib/build/provider/local"
)

type Builder interface {
	Build() error
}

func NewBuilder(repo lib.Repository, cloneDir string) (Builder, error) {
	builder, err := localBuilder.NewLocalBuilder(repo, cloneDir)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return builder, nil
}
