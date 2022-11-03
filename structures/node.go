package structures

type Value interface {
	int | string
}

type Node[V Value] struct {
	id    string
	value V
}

func NewNode[V Value](id string, value V) Node[V] {
	return Node[V]{
		id:    id,
		value: value,
	}
}

func (n Node[V]) Id() string {
	return n.id
}

func (n Node[V]) Value() V {
	return n.value
}

func (n Node[V]) Equals(n2 Node[V]) bool {
	return n.id == n2.id && n.value == n2.value
}
