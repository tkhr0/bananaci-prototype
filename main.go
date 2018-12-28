package main

import (
	"github.com/tkhr0/bananaci-prototype/runtime"
)

func main() {
	ob := runtime.NewRuntime()
	ob.Run()

	ob.ToLabeled()
	ob.Run()

	ob.ToClosed()
	ob.Run()

	ob.ToLabeled() // panic
	ob.Run()
}
