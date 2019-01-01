package main

import (
	"github.com/tkhr0/bananaci-prototype/server"
)

func main() {
	maxWorkers := 3
	maxQueues := 10000

	server.Call(maxWorkers, maxQueues)
}
