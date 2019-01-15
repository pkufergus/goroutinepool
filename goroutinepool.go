package goroutinepool

import "sync"

type RoutinePool struct {
	ch chan bool
	waitgroup sync.WaitGroup
}

type Job func()

func NewRoutinePool(max int) *RoutinePool {
	return &RoutinePool{
		ch: make(chan bool, max),
		waitgroup: sync.WaitGroup{},
	}
}

func RoutineRun(job Job, rp *RoutinePool) {
	defer rp.Done()
	job()
}

func (rp *RoutinePool) AddJob(job Job) {
	rp.ch<- false
	rp.waitgroup.Add(1)
	go RoutineRun(job, rp)
}

func (rp *RoutinePool) Done() {
	rp.waitgroup.Done()
	<-rp.ch
}

func (rp *RoutinePool) WaitAll() {
	rp.waitgroup.Wait()
}