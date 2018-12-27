package banana

import (
	"github.com/tkhr0/bananaci-prototype/phase"
)

type Banana struct {
	phase.Phase
}

func NewBanana() *Banana {
	return &Banana{
		Phase: phase.NewOpened(),
	}
}

func (b *Banana) ToLabeled() {
	nextPhase := b.Phase.ToLabeled()
	b.Phase = nextPhase
}

func (b *Banana) ToClosed() {
	nextPhase := b.Phase.ToClosed()
	b.Phase = nextPhase
}
