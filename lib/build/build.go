package build

import (
	"fmt"

	"github.com/tkhr0/bananaci-prototype/lib"
	localBuilder "github.com/tkhr0/bananaci-prototype/lib/build/provider/local"
)

type Builder interface {
	Build() error
	Up() error
}

func NewBuilder(repo lib.Repository) (Builder, error) {
	builder, err := localBuilder.NewLocalBuilder(repo)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return builder, nil
}
