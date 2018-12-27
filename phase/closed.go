package phase

import (
	"github.com/tkhr0/bananaci-prototype/phase/phase_type"
)

type Closed struct {
	BasePhase
}

func newClosed() Closed {
	return Closed{
		BasePhase: newBasePhase(phaseType.Closed),
	}
}

func (p Closed) Run() {
	println("Closed.Run")
}
