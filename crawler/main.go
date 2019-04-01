package main

import (
	"imooc.com/crawler/engine"
	"imooc.com/crawler/scheduler"
	"imooc.com/crawler/zhenai/parser"
)

func main() {
	// 简单引擎
	// simpleEngine := engine.SimpleEngine{}
	// 并发引擎
	ce := engine.ConcurrentEngine{
		WorkerNum: 100,
		Scheduler: &scheduler.SimpleScheduler{},
	}
	// 获取城市列表
	ce.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFun: parser.ParseCityList,
	})
}
