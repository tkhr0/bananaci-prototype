package phase

import (
	"fmt"
	"time"

	phaseType "github.com/tkhr0/bananaci-prototype/lib/phase/phase_type"
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
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%p Labeled.Run: running...#%d\n", &p, i+1)
	}
}
