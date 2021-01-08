package main

import (
	"GoCrawler/engine"
	"GoCrawler/persist"
	"GoCrawler/scheduler"
	"GoCrawler/ygdy8/parser"
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
