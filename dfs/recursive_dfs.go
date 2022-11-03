package dfs

import (
	"github.com/rs/zerolog/log"
	"graph-data-structures/core"
	graphPkg "graph-data-structures/graph"
)

type RecursiveDfs[V core.Value] struct {
	visited map[core.NodeId]bool
	graph   graphPkg.Graph[V]
	result  []core.NodeId
}

func NewRecursiveDfs[V core.Value](graph graphPkg.Graph[V]) RecursiveDfs[V] {
	return RecursiveDfs[V]{
		visited: make(map[core.NodeId]bool, graph.Length()),
		graph:   graph,
	}
}

func (dfs *RecursiveDfs[V]) Dfs(start core.Node[V]) (err error) {
	log.Debug().Int("id", start.Id()).Msg("dfs current node")

	dfs.visited[start.Id()] = true
	dfs.result = append(dfs.result, start.Id())

	neighbours, err := dfs.graph.Edges(start.Id())

	if err != nil {
		return
	}

	for _, neighbour := range neighbours {
		if dfs.visited[neighbour] {
			continue
		}

		var nextNode core.Node[V]
		nextNode, err = dfs.graph.Node(neighbour)

		if err != nil {
			return
		}

		err = dfs.Dfs(nextNode)
	}

	return
}

func (dfs *RecursiveDfs[V]) Result() []core.NodeId {
	return dfs.result
}
