package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 创建ParserResult chan,用来接收Fetcher的结果(Requests和Items)
	out := make(chan ParserResult)
	// 创建调度器管理的Request chan 和WorkerChan chan
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		// 创建worker chan，
		createWork(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		// 将Request放入Request chan
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			item := item
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWork(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			// 将worker chan 放到 workerChan chan里
			ready.WorkerReady(in)
			// 消费worker  chan中的Request
			request := <-in
			// 将Request 交给Fetcher,得到ParserResult
			result, err := worker(request)
			if err != nil {
				continue
			}
			// 将ParserResult放入ParserResult chan中
			out <- result
		}
	}()
}
