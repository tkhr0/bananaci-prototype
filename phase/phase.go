package phase

import (
	"github.com/tkhr0/bananaci-prototype/phase/phase_type"
)

type Phase interface {
	Run()
	Done()
	GetName() string

	ToLabeled() Labeled
	ToClosed() Closed
}

type BasePhase struct {
	phaseType.PhaseType
	State bool
}

func newBasePhase(t phaseType.PhaseType) BasePhase {
	return BasePhase{
		PhaseType: t,
		State:     false,
	}
}

func (p BasePhase) Run() {
	panic("not implement")
}

func (p BasePhase) Done() {
	panic("not implement")
}

func (p BasePhase) GetName() string {
	return p.PhaseType.String()
}

func (p BasePhase) ToLabeled() Labeled {
	panic("not implement")
}

func (p BasePhase) ToClosed() Closed {
	panic("not implement")
}
