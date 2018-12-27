package phase

import (
	"github.com/tkhr0/bananaci-prototype/phase/phase_type"
)

type Phase interface {
	Run()
	Done()
	GetName() string
}

type BasePhase struct {
	PhaseType PhaseType.PhaseType
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
