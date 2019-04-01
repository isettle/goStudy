package engine

import (
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	WorkerNum int
	Scheduler Scheduler
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterChan(chan Request)
}

func (ce *ConcurrentEngine) Run(seeds ...Request) {
	// 通道并发
	in := make(chan Request)
	out := make(chan ParserResult)
	ce.Scheduler.ConfigMasterChan(in)
	// 创建worker
	for i := 0; i < ce.WorkerNum; i++ {
		createWorker(in, out)
	}

	// 入通道
	for _, seed := range seeds {
		ce.Scheduler.Submit(seed)
	}
	itemsNum := 0
	for {
		parserResult := <-out
		// 结果日志打印
		for _, item := range parserResult.Items {
			log.Printf("+got item#%d: %+v", itemsNum, item)
			itemsNum++
		}
		// 入任务列
		for _, r := range parserResult.Requests {
			ce.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult) {
	fmt.Println("creating worker")
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()

}
