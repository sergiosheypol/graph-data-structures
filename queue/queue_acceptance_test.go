package queue

import (
	"github.com/stretchr/testify/assert"
	"graph-data-structures/core"
	"testing"
)

func TestQueue_FullFlow(t *testing.T) {
	// Given
	q := NewQueue[int]()
	n1 := core.NewNode(1, 1)
	n2 := core.NewNode(2, 2)

	// When
	q.Enqueue(n1)
	q.Enqueue(n2)
	nActual, err := q.Dequeue()

	// Then
	assert.NoError(t, err)
	assert.EqualValues(t, n1, nActual)
}
