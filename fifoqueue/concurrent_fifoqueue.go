package fifoqueue

import (
	"sync"
)

// ConcurrentQueue is storage of current queue data
type ConcurrentQueue struct {
	lock    *sync.Mutex
	backend *QueueBackend
}

// CreateConcurrentQueue is for creating a concurrent FIFO Queue
func CreateConcurrentQueue(maxSize uint32) *ConcurrentQueue {
	return &ConcurrentQueue{
		lock:    &sync.Mutex{},
		backend: CreateQueue(maxSize),
	}
}

// Push is for pushing data into the current FIFO Queue
func (cq *ConcurrentQueue) Push(data interface{}) error {
	cq.lock.Lock()
	defer cq.lock.Unlock()

	err := cq.backend.Push(data)

	if err != nil {
		return err
	}

	return nil
}

// Pop is for popping data into the current FIFO Queue
func (cq *ConcurrentQueue) Pop() (interface{}, error) {
	cq.lock.Lock()
	defer cq.lock.Unlock()

	data, err := cq.backend.Pop()

	if err != nil {
		return data, err
	}

	return data, nil
}

// GetCurrentSize is for getting size of the current FIFO Queue
func (cq *ConcurrentQueue) GetCurrentSize() uint32 {
	cq.lock.Lock()

	size := cq.backend.GetCurrentSize()

	cq.lock.Unlock()

	return size
}
