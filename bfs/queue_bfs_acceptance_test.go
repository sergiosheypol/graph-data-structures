package bfs

import (
	"github.com/stretchr/testify/assert"
	"graph-data-structures/core"
	"graph-data-structures/graph"
	"testing"
)

func TestBfs(t *testing.T) {
	// Setup
	g, n1, err := createGraph()
	assert.NoError(t, err)

	// Given
	bfs := NewBfs[int](g)

	// When
	err = bfs.Bfs(n1)
	actual := bfs.Result()

	// Then
	assert.NoError(t, err)
	assert.Exactly(t, []int{1, 2, 4, 3, 5, 6}, actual)
}

func createGraph() (g graph.Graph[int], n1 core.Node[int], err error) {
	n1 = core.NewNode(1, 1)
	n2 := core.NewNode(2, 2)
	n3 := core.NewNode(3, 3)
	n4 := core.NewNode(4, 4)
	n5 := core.NewNode(5, 5)
	n6 := core.NewNode(6, 6)

	g = graph.NewGraph[int](false)

	err = g.AddNode(n1)
	if err != nil {
		return
	}

	err = g.AddNode(n2)
	if err != nil {
		return
	}

	err = g.AddNode(n3)
	if err != nil {
		return
	}

	err = g.AddNode(n4)
	if err != nil {
		return
	}

	err = g.AddNode(n5)
	if err != nil {
		return
	}

	err = g.AddNode(n6)
	if err != nil {
		return
	}

	err = g.AddEdge(n1, n2)
	if err != nil {
		return
	}

	err = g.AddEdge(n2, n3)
	if err != nil {
		return
	}

	err = g.AddEdge(n1, n4)
	if err != nil {
		return
	}

	err = g.AddEdge(n4, n5)
	if err != nil {
		return
	}

	err = g.AddEdge(n5, n6)
	if err != nil {
		return
	}
	return
}
