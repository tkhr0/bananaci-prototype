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
			State:     false,
		},
	}
}

func (p Opened) Run() {
	fmt.Println("Opened.Run")
}

func (p Opened) ToLabeled() Labeled {
	println("Opened.ToLabeled")
	return newLabeled()
}

func (p Opened) ToClose() Closed {
	println("Opened.ToClose")
	return newClosed()
}
