package workerpool

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/ggary9424/go-worker-pool-practice/fifoqueue"
)

// WorkerPool is storage of worker pool
type WorkerPool struct {
	maxWorkersCount   int
	maxQueueSize      uint32
	queue             *fifoqueue.ConcurrentQueue
	jobsChan          chan func() error
	closeRoutinesChan chan bool
	routinesClosedWg  sync.WaitGroup
	isLogEnabled      bool
}

// CreateWorkerPool is for creating WorkerPool Object
func CreateWorkerPool(
	maxWorkersCount int,
	maxQueueSize uint32,
	isLogEnabled bool,
) *WorkerPool {
	var wg sync.WaitGroup
	wp := &WorkerPool{
		maxWorkersCount:   maxWorkersCount,
		maxQueueSize:      maxQueueSize,
		queue:             fifoqueue.CreateConcurrentQueue(maxQueueSize),
		jobsChan:          make(chan func() error, maxWorkersCount),
		closeRoutinesChan: make(chan bool),
		routinesClosedWg:  wg,
		isLogEnabled:      isLogEnabled,
	}

	for i := 0; i < maxWorkersCount; i++ {
		jobID := i
		go func() {
			wp.routinesClosedWg.Add(1)
			defer wp.routinesClosedWg.Done()

			wp.log("(Consumer) Worker[%d] is started\n", jobID)
			for job := range wp.jobsChan {
				wp.log("(Consumer) Worker[%d] received a job\n", jobID)
				job()
				wp.log("(Consumer) Worker[%d] finished a job\n", jobID)
			}
			wp.log("(Consumer) Worker[%d] is closed\n", jobID)
		}()
	}

	go func() {
		wp.routinesClosedWg.Add(1)
		defer wp.routinesClosedWg.Done()

		for {
			select {
			case <-wp.closeRoutinesChan:
				close(wp.jobsChan)
				close(wp.closeRoutinesChan)
				wp.log("(Producer) Dispatcher is closed\n")
				return
			default:
				job, _ := wp.queue.Pop()
				if f, ok := job.(func() error); ok {
					wp.log("(Producer) Dispatch a job\n")
					wp.jobsChan <- (func() error)(f)
				}
			}
			runtime.Gosched()
		}
	}()

	return wp
}

// Do is Do
func (wp *WorkerPool) Do(job func() error) {
	wp.queue.Push(job)
}

// Close is Close
func (wp *WorkerPool) Close() {
	wp.log("(Controller) Receive close workpool singal\n")
	wp.closeRoutinesChan <- true
	wp.log("(Controller) Wait for all routines closed\n")
	wp.routinesClosedWg.Wait()
	wp.log("(Controller) All routines are closed\n")
}

func (wp *WorkerPool) log(format string, a ...interface{}) {
	if wp.isLogEnabled {
		fmt.Printf(format, a...)
	}
}
