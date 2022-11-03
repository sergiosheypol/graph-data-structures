package queue

import (
	"testing"

	"graph-data-structures/core"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	type args struct {
		nodes []core.Node[int]
	}
	tests := []struct {
		name  string
		args  args
		wantS Queue[int]
	}{
		{
			name: "Create empty queue when no nodes (nil)",
			wantS: Queue[int]{
				items: []core.Node[int]{},
			},
		},
		{
			name: "Create empty queue from empty slice",
			args: args{
				nodes: []core.Node[int]{},
			},
			wantS: Queue[int]{
				items: []core.Node[int]{},
			},
		},
		{
			name: "Create empty queue from nil",
			args: args{
				nodes: nil,
			},
			wantS: Queue[int]{
				items: []core.Node[int]{},
			},
		},
		{
			name: "Create empty queue with nodes",
			args: args{
				nodes: []core.Node[int]{
					core.NewNode(1, 1),
					core.NewNode(2, 14),
				},
			},
			wantS: Queue[int]{
				items: []core.Node[int]{
					core.NewNode(1, 1),
					core.NewNode(2, 14),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := NewQueue(tt.args.nodes...); !assert.EqualValues(t, tt.wantS, gotS) {
				t.Errorf("NewQueue() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type args struct {
		n core.Node[int]
	}
	tests := []struct {
		name  string
		queue Queue[int]
		args  args
		want  []core.Node[int]
	}{
		{
			name:  "Ok",
			queue: NewQueue[int](),
			args: args{
				n: core.NewNode[int](1, 67),
			},
			want: []core.Node[int]{
				core.NewNode[int](1, 67),
			},
		},
		{
			name:  "Ok With More Elements",
			queue: NewQueue[int](core.NewNode[int](1, 1)),
			args: args{
				n: core.NewNode[int](2, 44),
			},
			want: []core.Node[int]{
				core.NewNode[int](1, 1),
				core.NewNode[int](2, 44),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.queue.Enqueue(tt.args.n)
			assert.Exactly(t, tt.queue.items, tt.want)
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	tests := []struct {
		name           string
		queue          Queue[string]
		wantNode       core.Node[string]
		wantQueueNodes []core.Node[string]
		wantErr        bool
	}{
		{
			name:     "Ok",
			queue:    NewQueue[string](core.NewNode[string](1, "blue"), core.NewNode[string](2, "red"), core.NewNode[string](3, "green")),
			wantNode: core.NewNode[string](1, "blue"),
			wantQueueNodes: []core.Node[string]{
				core.NewNode[string](2, "red"),
				core.NewNode[string](3, "green"),
			},
			wantErr: false,
		},
		{
			name:    "Err Empty Queue",
			queue:   NewQueue[string](),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.queue.Dequeue()
			if tt.wantErr {
				assert.ErrorIs(t, err, core.ErrPopEmpty, "Dequeue() - Did not return an error")
				return
			}
			assert.EqualValues(t, tt.wantNode, gotN, "Dequeue() - Different node expected")
			assert.Exactly(t, tt.wantQueueNodes, tt.queue.items, "Dequeue() - Remaining nodes do not match")
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		queue Queue[int]
		want  bool
	}{
		{
			name:  "Empty - true",
			queue: NewQueue[int](),
			want:  true,
		},
		{
			name:  "Empty - false",
			queue: NewQueue[int](core.NewNode[int](1, 12432432)),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equalf(t, tt.want, tt.queue.IsEmpty(), "IsEmpty()")
		})
	}
}
