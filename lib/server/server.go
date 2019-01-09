package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	jobQueue "github.com/tkhr0/bananaci-prototype/lib/job_queue"
	"github.com/tkhr0/bananaci-prototype/lib/runtime"
)

func Call(d *jobQueue.Dispatcher) {

	sigHandler(d)

	go func() {
		for i := 0; i < 100; i++ {
			d.Add(generator())
		}
	}()

	go func() {
		for {
			fmt.Printf("pool: %d queue: %d\n", d.Pool(), d.Queue())
			time.Sleep(1 * time.Second)
		}
	}()

	println("start")
	d.Start()
	println("wait")
	d.Wait()
	println("exit")
}

func sigHandler(d *jobQueue.Dispatcher) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGUSR1, syscall.SIGINT)

	go func() {
		for {
			switch s := <-sigCh; s {
			case syscall.SIGUSR1:
				d.Add(generator())
			case syscall.SIGINT:
				d.Stop(true)
				return
			}
		}
	}()
}

func generator() jobQueue.Job {
	return *jobQueue.NewJob(*runtime.NewRuntime())
}
