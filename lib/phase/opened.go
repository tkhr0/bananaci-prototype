package phase

import (
	"fmt"
	"time"

	phaseType "github.com/tkhr0/bananaci-prototype/lib/phase/phase_type"
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
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%p Opened.Run: running...#%d\n", &p, i+1)
	}
}

func (p Opened) ToLabeled() Labeled {
	println("Opened.ToLabeled")
	return newLabeled()
}

func (p Opened) ToClose() Closed {
	println("Opened.ToClose")
	return newClosed()
}
