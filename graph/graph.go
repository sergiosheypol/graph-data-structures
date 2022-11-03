package graph

import (
	"github.com/rs/zerolog/log"
	"graph-data-structures/core"
)

type Graph[V core.Value] struct {
	adjacencyList map[core.NodeId][]core.NodeId
	nodes         map[core.NodeId]core.Node[V]
	directed      bool
}

func NewGraph[V core.Value](directed bool) Graph[V] {
	return Graph[V]{
		adjacencyList: make(map[core.NodeId][]core.NodeId),
		nodes:         make(map[core.NodeId]core.Node[V]),
		directed:      directed,
	}
}

func (g *Graph[V]) AddNode(node core.Node[V]) (err error) {
	if n, ok := g.nodes[node.Id()]; ok {
		err = core.ErrDuplicatedNode
		log.Err(err).Int("node", n.Id()).Msg("node already exists")
		return
	}

	g.nodes[node.Id()] = node
	log.Info().Int("node", node.Id()).Msg("added node")
	return
}

func (g *Graph[V]) AddEdge(from core.Node[V], to core.Node[V]) (err error) {
	if _, ok := g.nodes[from.Id()]; !ok {
		err = core.ErrNotFoundNode
		log.Err(err).Int("node", from.Id()).Msgf("node does not exist in graph")
		return
	}

	if _, ok := g.nodes[to.Id()]; !ok {
		err = core.ErrNotFoundNode
		log.Err(err).Int("node", to.Id()).Msgf("node does not exist in graph")
		return
	}

	fromVertex := g.adjacencyList[from.Id()]
	fromVertex = append(fromVertex, to.Id())
	g.adjacencyList[from.Id()] = fromVertex

	log.Info().Int("node", from.Id()).Int("edge", to.Id()).Msg("new edge added")

	if !g.directed {
		toVertex := g.adjacencyList[to.Id()]
		toVertex = append(toVertex, from.Id())
		g.adjacencyList[to.Id()] = toVertex
		log.Info().Int("node", to.Id()).Int("edge", from.Id()).Msg("new edge added")
	}

	return
}

func (g *Graph[V]) Edges(id core.NodeId) (edges []core.NodeId, err error) {
	edges, ok := g.adjacencyList[id]

	if !ok {
		err = core.ErrNotFoundNode
		log.Err(err).Int("node", id).Msgf("node does not exist in graph")
		return
	}

	return
}

func (g *Graph[V]) Node(id core.NodeId) (node core.Node[V], err error) {
	node, ok := g.nodes[id]

	if !ok {
		err = core.ErrNotFoundNode
		log.Err(err).Int("node", id).Msgf("node does not exist in graph")
		return
	}

	return
}

func (g *Graph[V]) Length() int {
	return len(g.nodes)
}

func (g *Graph[V]) RemoveNode(_ core.Node[V]) {
	log.Panic().Msg("unsupported feature")
}
