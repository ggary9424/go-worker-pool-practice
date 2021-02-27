package fifoqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentQueueFunctionalityWorks(t *testing.T) {
	q := CreateConcurrentQueue(100000)

	for i := 0; i < 100; i++ {
		go func() {
			q.Push(1)
			_, err := q.Pop()

			assert.Equal(t, err, nil)
		}()
	}
}
