package limitedWaitGroup

import "sync"

type LimitedWaitGroup struct {
	waitGroup sync.WaitGroup
	limiter   chan bool
}

func NewLimitedWaitGroup(Limiter uint) *LimitedWaitGroup {
	tc := LimitedWaitGroup{
		waitGroup: sync.WaitGroup{},
		limiter:   make(chan bool, Limiter),
	}
	return &tc
}

func (tc *LimitedWaitGroup) AddOne() {
	tc.waitGroup.Add(1)
	tc.limiter <- true
}

func (tc *LimitedWaitGroup) Done() {
	tc.waitGroup.Done()
	<-tc.limiter
}

func (tc *LimitedWaitGroup) Wait() {
	tc.waitGroup.Wait()
}
