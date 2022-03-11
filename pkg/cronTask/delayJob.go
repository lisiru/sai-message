package cronTask

import (
	"sync"
	"time"
)

type DelayJob struct {
	Id    uint64
	Delay time.Duration
	Stat  int
	Func  func()
}

type Option func(*DelayJob)

var mux sync.Mutex

func NewDelayJob(options ...func(job *DelayJob)) *DelayJob {
	delayJob := &DelayJob{}
	for _, option := range options {

		option(delayJob)

	}
	return delayJob

}

func WithJobId(id uint64) Option {
	return func(job *DelayJob) {
		mux.Lock()
		defer mux.Unlock()
		job.Id = id
	}
}

func WithDelay(delayTime time.Duration) Option {
	return func(job *DelayJob) {
		mux.Lock()
		defer mux.Unlock()
		job.Delay = delayTime
	}
}

func WithStat(stat int) Option {
	return func(job *DelayJob) {
		mux.Lock()
		defer mux.Unlock()
		job.Stat = stat
	}
}
func WithFunc(fun func()) Option {
	return func(job *DelayJob) {
		mux.Lock()
		defer mux.Unlock()
		job.Func = fun
	}
}

func (delayJob *DelayJob) AddJob() {



}
