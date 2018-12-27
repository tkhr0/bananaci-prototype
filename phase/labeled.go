package phase

import (
	"github.com/tkhr0/bananaci-prototype/phase/phase_type"
)

type Labeled struct {
	BasePhase
}

func NewLabeled() *Labeled {
	return &Labeled{
		BasePhase: BasePhase{
			PhaseType: PhaseType.Labeled,
		},
	}
}
