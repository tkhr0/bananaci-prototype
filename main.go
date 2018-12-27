package main

import (
	"github.com/tkhr0/bananaci-prototype/banana"
)

func main() {
	ob := banana.NewBanana()
	ob.Run()

	ob.ToLabeled()
	ob.Run()

	ob.ToClosed()
	ob.Run()

	ob.ToLabeled() // panic
	ob.Run()
}
