package stack

import (
	"github.com/stretchr/testify/assert"
	"graph-data-structures/core"
	"testing"
)

func TestStack_FullFlow(t *testing.T) {
	// Given
	s := NewStack[int]()
	n1 := core.NewNode(1, 1)
	n2 := core.NewNode(2, 2)

	// When
	s.Push(n1)
	s.Push(n2)
	nActual, err := s.Pop()

	// Then
	assert.NoError(t, err)
	assert.EqualValues(t, n2, nActual)
}
