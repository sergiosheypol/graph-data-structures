package core

type Value interface {
	int | string
}

type NodeId = int

type Node[V Value] struct {
	id    NodeId
	value V
}

func NewNode[V Value](id int, value V) Node[V] {
	return Node[V]{
		id:    id,
		value: value,
	}
}

func (n Node[V]) Id() int {
	return n.id
}

func (n Node[V]) Value() V {
	return n.value
}

// Equals TODO change equal to be more resilient
func (n Node[V]) Equals(n2 Node[V]) bool {
	return n.id == n2.id && n.value == n2.value
}
