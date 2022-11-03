package dfs

import (
	"github.com/stretchr/testify/assert"
	"graph-data-structures/core"
	"graph-data-structures/graph"
	"testing"
)

func TestRecursiveDfs(t *testing.T) {
	// Setup
	g, n1, err := createRecursiveDfsTestGraph()
	assert.NoError(t, err)

	// Given
	dfs := NewRecursiveDfs[int](g)

	// When
	err = dfs.Dfs(n1)
	actual := dfs.Result()

	// Then
	assert.NoError(t, err)
	assert.Exactly(t, []int{1, 2, 3, 4}, actual)
}

func createRecursiveDfsTestGraph() (g graph.Graph[int], n1 core.Node[int], err error) {
	n1 = core.NewNode(1, 1)
	n2 := core.NewNode(2, 2)
	n3 := core.NewNode(3, 3)
	n4 := core.NewNode(4, 4)

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

	err = g.AddEdge(n1, n2)
	if err != nil {
		return
	}

	err = g.AddEdge(n1, n3)
	if err != nil {
		return
	}

	err = g.AddEdge(n3, n4)
	if err != nil {
		return
	}
	return
}
