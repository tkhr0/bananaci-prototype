package jobQueue

import (
	"sync"
)

type Dispatcher struct {
	pool    chan *worker
	queue   chan Job
	workers []*worker
	wg      sync.WaitGroup
	quit    chan struct{}
}

func NewDispatcher(maxWorkers int, maxQueues int) *Dispatcher {
	d := &Dispatcher{
		pool:  make(chan *worker, maxWorkers),
		queue: make(chan Job, maxQueues),
		quit:  make(chan struct{}),
	}

	d.workers = make([]*worker, cap(d.pool))

	for i := 0; i < cap(d.pool); i++ {
		w := worker{
			dispatcher: d,
			job:        make(chan Job),
			quit:       make(chan struct{}),
		}
		d.workers[i] = &w
	}

	return d
}

func (d *Dispatcher) Add(j Job) {
	d.wg.Add(1)
	d.queue <- j
	println("added")
}

func (d *Dispatcher) Start() {
	for _, w := range d.workers {
		w.start()
	}

	d.wg.Add(1) // Queue がなくても維持する
	go func() {
		for {
			select {
			case j := <-d.queue:
				(<-d.pool).job <- j
			case <-d.quit:
				return
			}
		}
	}()
}

func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

func (d *Dispatcher) Stop(immediately bool) {
	defer d.wg.Done()

	if !immediately {
		d.Wait()
	}

	d.quit <- struct{}{}
	for _, w := range d.workers {
		w.quit <- struct{}{}
	}
}

func (d *Dispatcher) Pool() int {
	return len(d.pool)
}

func (d *Dispatcher) Queue() int {
	return len(d.queue)
}
