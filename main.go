package main

import (
	"sync"

	jobQueue "github.com/tkhr0/bananaci-prototype/job_queue"
	"github.com/tkhr0/bananaci-prototype/server"
)

func main() {
	maxWorkers := 3
	maxQueues := 10000
	d := jobQueue.NewDispatcher(maxWorkers, maxQueues)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		server.Call(d)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		server.CallHTTP(d)
		wg.Done()
	}()

	wg.Wait()
}
