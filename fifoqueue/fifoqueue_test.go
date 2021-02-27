package fifoqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentSize(t *testing.T) {
	q := CreateQueue(100000)
	pushTimes, popTimes := 30, 18
	for i := 0; i < pushTimes; i++ {
		q.Push(i)
	}
	for i := 0; i < popTimes; i++ {
		q.Pop()
		assert.Equal(t, q.GetCurrentSize(), uint32(pushTimes-i-1))
	}
}

func TestFunctionalityWorks(t *testing.T) {
	q := CreateQueue(100000)
	pushTimes, popTimes := 30, 18
	for i := 0; i < pushTimes; i++ {
		q.Push(i)
	}
	for i := 0; i < popTimes; i++ {
		data, _ := q.Pop()
		assert.Equal(t, data, i)
	}
}
