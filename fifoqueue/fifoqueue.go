package fifoqueue

import "errors"

// Node is storage of queue data
type Node struct {
	data interface{}
	next *Node
}

// QueueBackend is storage of the queue
type QueueBackend struct {
	head *Node
	tail *Node

	// current size of queue
	size    uint32
	maxSize uint32
}

// CreateQueue is for creating a FIFO Queue
func CreateQueue(maxSize uint32) *QueueBackend {
	return &QueueBackend{
		maxSize: maxSize,
	}
}

func (q *QueueBackend) createNode(data interface{}) *Node {
	node := Node{}
	node.data = data
	node.next = nil

	return &node
}

// Push is for pushing data into the FIFO Queue
func (q *QueueBackend) Push(data interface{}) error {
	if q.size >= q.maxSize {
		err := errors.New("Queue is full")
		return err
	}

	if q.size == 0 {
		node := q.createNode(data)
		q.head = node
		q.tail = node

		q.size = 1

		return nil
	}

	newNode := q.createNode(data)
	q.tail.next = newNode
	q.tail = newNode

	q.size++
	return nil
}

// Pop is for popping data into the FIFO Queue
func (q *QueueBackend) Pop() (interface{}, error) {
	if q.size == 0 {
		err := errors.New("Queue is empty")
		return nil, err
	}

	currentHead := q.head
	newHead := currentHead.next
	currentHead.next = nil

	q.size--
	if q.size == 0 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = newHead
	}

	return currentHead.data, nil
}

// GetCurrentSize is for getting size of the FIFO Queue
func (q *QueueBackend) GetCurrentSize() uint32 {
	return q.size
}
