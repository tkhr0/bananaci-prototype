package jobQueue

import (
	"github.com/tkhr0/bananaci-prototype/phase/phase_type"
	"github.com/tkhr0/bananaci-prototype/runtime"
)

type Job struct {
	runtime.Runtime
}

func NewJob(r runtime.Runtime) *Job {
	return &Job{
		Runtime: r,
	}
}

func (j *Job) Run() bool {
	switch j.GetPhaseType() {
	case phaseType.Opened:
		j.Runtime.ToLabeled()
		j.Runtime.Run()
		return true
	case phaseType.Labeled:
		j.Runtime.ToClosed()
		j.Runtime.Run()
		return true
	default:
		return false
	}
}