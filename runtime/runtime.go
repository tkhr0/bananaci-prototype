package runtime

import (
	"github.com/tkhr0/bananaci-prototype/phase"
)

type Runtime struct {
	phase.Phase
}

func NewRuntime() *Runtime {
	return &Runtime{
		Phase: phase.NewOpened(),
	}
}

func (b *Runtime) ToLabeled() {
	nextPhase := b.Phase.ToLabeled()
	b.Phase = nextPhase
}

func (b *Runtime) ToClosed() {
	nextPhase := b.Phase.ToClosed()
	b.Phase = nextPhase
}
