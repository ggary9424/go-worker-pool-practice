package main

import (
	"time"

	"github.com/ggary9424/go-worker-pool-practice/fifoqueue"
	"github.com/ggary9424/go-worker-pool-practice/workerpool"
)

func main() {
	// Try Concurrent FIFO Queue
	println("=========== Try Concurrent FIFO Queue ===========")
	q := fifoqueue.CreateConcurrentQueue(100000)

	var funcs []func()
	for i := 0; i < 100; i++ {
		jobID := i
		q.Push(jobID)
		f := func() {
			q.Pop()
		}
		funcs = append(funcs, f)
	}
	for _, f := range funcs {
		go f()
	}
	time.Sleep(time.Second * 3)

	// Try Worker Pool
	println("=========== Try Worker Pool ===========")
	wp := workerpool.CreateWorkerPool(5, 1000000000, true)
	for i := 0; i < 100; i++ {
		c := i
		wp.Do(func() error {
			println(c)
			time.Sleep(time.Second * 3)

			return nil
		})
	}

	time.Sleep(time.Second * 3)
	wp.Close()
}
