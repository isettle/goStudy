package scheduler

import "imooc.com/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (ss *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		ss.workerChan <- request
	}()
}

func (ss *SimpleScheduler) ConfigMasterChan(c chan engine.Request) {
	ss.workerChan = c
}
