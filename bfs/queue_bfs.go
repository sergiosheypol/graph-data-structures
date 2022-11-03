package bfs

import (
	"github.com/rs/zerolog/log"
	"graph-data-structures/core"
	graphPkg "graph-data-structures/graph"
	"graph-data-structures/queue"
)

type Bfs[V core.Value] struct {
	toVisit queue.Queue[V]
	visited map[core.NodeId]bool
	graph   graphPkg.Graph[V]
	result  []core.NodeId
}

func NewBfs[V core.Value](graph graphPkg.Graph[V]) Bfs[V] {
	return Bfs[V]{
		toVisit: queue.NewQueue[V](),
		visited: make(map[core.NodeId]bool, graph.Length()),
		graph:   graph,
	}
}

func (bfs *Bfs[V]) Bfs(start core.Node[V]) (err error) {
	bfs.toVisit.Enqueue(start)

	for !bfs.toVisit.IsEmpty() {
		var nextNode core.Node[V]
		nextNode, _ = bfs.toVisit.Dequeue()

		if v := bfs.visited[nextNode.Id()]; v {
			continue
		}

		log.Debug().Int("id", nextNode.Id()).Msg("bfs current node")

		bfs.visited[nextNode.Id()] = true
		bfs.result = append(bfs.result, nextNode.Id())

		var neighbours []core.NodeId
		neighbours, err = bfs.graph.Edges(nextNode.Id())

		if err != nil {
			return
		}

		for _, neighbour := range neighbours {
			var node core.Node[V]
			node, err = bfs.graph.Node(neighbour)
			if err != nil {
				return
			}
			bfs.toVisit.Enqueue(node)
		}
	}

	return
}

func (bfs *Bfs[V]) Result() []core.NodeId {
	return bfs.result
}
