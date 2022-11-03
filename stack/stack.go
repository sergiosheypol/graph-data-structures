package stack

import (
	"graph-data-structures/core"
)

type Stack[V core.Value] struct {
	items []core.Node[V]
}

func NewStack[V core.Value](nodes ...core.Node[V]) (s Stack[V]) {
	if nodes == nil || len(nodes) == 0 {
		s.items = make([]core.Node[V], 0)
		return
	}

	s.items = nodes
	return
}

func (s *Stack[V]) Push(n core.Node[V]) {
	s.items = append(s.items, n)
}

func (s *Stack[V]) Pop() (n core.Node[V], err error) {
	if len(s.items) < 1 {
		err = core.ErrPopEmpty
		return
	}

	n = s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]
	return
}
func (s *Stack[V]) IsEmpty() bool {
	return len(s.items) == 0
}
