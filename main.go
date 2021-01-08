package main

import (
	"gowork/crawler/engine"
	"gowork/crawler/persist"
	"gowork/crawler/scheduler"
	"gowork/crawler/ygdy8/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        parser.PrefixUrl,
		ParserFunc: parser.ParseTypeList,
	})
}
