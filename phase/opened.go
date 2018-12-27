package phase

import (
	"fmt"

	"github.com/tkhr0/bananaci-prototype/phase/phase_type"
)

type Opened struct {
	BasePhase
}

func NewOpened() *Opened {
	return &Opened{
		BasePhase: BasePhase{
			PhaseType: phaseType.Opened,
		},
	}
}

func (p Opened) Run() {
	fmt.Println("Opened.Run")
}
