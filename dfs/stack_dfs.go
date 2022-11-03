package dfs

import (
	"github.com/rs/zerolog/log"
	"graph-data-structures/core"
	graphPkg "graph-data-structures/graph"
	"graph-data-structures/stack"
)

type StackDfs[V core.Value] struct {
	toVisit stack.Stack[V]
	visited map[core.NodeId]bool
	graph   graphPkg.Graph[V]
	result  []core.NodeId
}

func NewStackDfs[V core.Value](graph graphPkg.Graph[V]) StackDfs[V] {
	return StackDfs[V]{
		toVisit: stack.NewStack[V](),
		visited: make(map[core.NodeId]bool, graph.Length()),
		graph:   graph,
	}
}

func (dfs *StackDfs[V]) Dfs(start core.Node[V]) (err error) {
	dfs.toVisit.Push(start)

	for !dfs.toVisit.IsEmpty() {
		var nextNode core.Node[V]
		nextNode, _ = dfs.toVisit.Pop()

		if v := dfs.visited[nextNode.Id()]; v {
			continue
		}

		log.Debug().Int("id", nextNode.Id()).Msg("dfs current node")

		dfs.visited[nextNode.Id()] = true
		dfs.result = append(dfs.result, nextNode.Id())

		var neighbours []core.NodeId
		neighbours, err = dfs.graph.Edges(nextNode.Id())

		if err != nil {
			return
		}

		for _, neighbour := range neighbours {
			var node core.Node[V]
			node, err = dfs.graph.Node(neighbour)
			if err != nil {
				return
			}
			dfs.toVisit.Push(node)
		}
	}

	return
}

func (dfs *StackDfs[V]) Result() []core.NodeId {
	return dfs.result
}
