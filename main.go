package main

import (
	"fmt"

	"github.com/tkhr0/bananaci-prototype/phase"
)

type Banana struct {
	phase.Phase
}

func main() {
	ob := Banana{
		Phase: phase.NewOpened(),
	}

	fmt.Println(ob.GetName())
	ob.Run()

	lb := Banana{
		Phase: phase.NewLabeled(),
	}

	fmt.Println(lb.GetName())
	lb.Run() // panic
}
