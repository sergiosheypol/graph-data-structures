package queue

import (
	"graph-data-structures/core"
)

type Queue[V core.Value] struct {
	items []core.Node[V]
}

func NewQueue[V core.Value](nodes ...core.Node[V]) (q Queue[V]) {
	if nodes == nil || len(nodes) == 0 {
		q.items = make([]core.Node[V], 0)
		return
	}

	q.items = nodes
	return
}

func (q *Queue[V]) Enqueue(n core.Node[V]) {
	q.items = append(q.items, n)
}

func (q *Queue[V]) Dequeue() (n core.Node[V], err error) {
	if len(q.items) < 1 {
		err = core.ErrPopEmpty
		return
	}

	n = q.items[0]

	q.items = q.items[1:]
	return
}

func (q *Queue[V]) IsEmpty() bool {
	return len(q.items) == 0
}
