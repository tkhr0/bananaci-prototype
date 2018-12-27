package phase

import (
	"github.com/tkhr0/bananaci-prototype/phase/phase_type"
)

type Labeled struct {
	BasePhase
}

func newLabeled() Labeled {
	return Labeled{
		BasePhase: newBasePhase(phaseType.Labeled),
	}
}

func (p Labeled) ToClosed() Closed {
	println("Labeled.ToClose")

	return newClosed()
}

func (p Labeled) Run() {
	println("Labeled.Run")
}
