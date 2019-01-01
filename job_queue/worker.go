package jobQueue

type worker struct {
	dispatcher *Dispatcher
	job        chan Job
	quit       chan struct{}
}

func (w *worker) start() {
	go func() {
		for {
			w.dispatcher.pool <- w

			select {
			case j := <-w.job:
				w.run(j)
				w.dispatcher.wg.Done()

			case <-w.quit:
				return
			}
		}
	}()
}

func (w *worker) run(j Job) {
	next := j.Run()

	if next {
		w.dispatcher.Add(j)
	}
}
