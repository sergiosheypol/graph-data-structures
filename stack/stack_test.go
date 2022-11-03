package stack

import (
	"testing"

	"graph-data-structures/core"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	type args struct {
		nodes []core.Node[int]
	}
	tests := []struct {
		name  string
		args  args
		wantS Stack[int]
	}{
		{
			name: "Create empty stack when no nodes (nil)",
			wantS: Stack[int]{
				items: []core.Node[int]{},
			},
		},
		{
			name: "Create empty stack from empty slice",
			args: args{
				nodes: []core.Node[int]{},
			},
			wantS: Stack[int]{
				items: []core.Node[int]{},
			},
		},
		{
			name: "Create empty stack from nil",
			args: args{
				nodes: nil,
			},
			wantS: Stack[int]{
				items: []core.Node[int]{},
			},
		},
		{
			name: "Create empty stack with nodes",
			args: args{
				nodes: []core.Node[int]{
					core.NewNode(1, 1),
					core.NewNode(2, 14),
				},
			},
			wantS: Stack[int]{
				items: []core.Node[int]{
					core.NewNode(1, 1),
					core.NewNode(2, 14),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := NewStack(tt.args.nodes...); !assert.EqualValues(t, tt.wantS, gotS) {
				t.Errorf("NewStack() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		n core.Node[int]
	}
	tests := []struct {
		name  string
		stack Stack[int]
		args  args
		want  []core.Node[int]
	}{
		{
			name:  "Ok",
			stack: NewStack[int](),
			args: args{
				n: core.NewNode[int](1, 67),
			},
			want: []core.Node[int]{
				core.NewNode[int](1, 67),
			},
		},
		{
			name:  "Ok With More Elements",
			stack: NewStack[int](core.NewNode[int](1, 1)),
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
			tt.stack.Push(tt.args.n)
			assert.Exactly(t, tt.stack.items, tt.want)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name           string
		stack          Stack[string]
		wantNode       core.Node[string]
		wantStackNodes []core.Node[string]
		wantErr        bool
	}{
		{
			name:     "Ok",
			stack:    NewStack[string](core.NewNode[string](1, "blue"), core.NewNode[string](2, "red"), core.NewNode[string](3, "green")),
			wantNode: core.NewNode[string](3, "green"),
			wantStackNodes: []core.Node[string]{
				core.NewNode[string](1, "blue"),
				core.NewNode[string](2, "red"),
			},
			wantErr: false,
		},
		{
			name:    "Err Empty Stack",
			stack:   NewStack[string](),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.stack.Pop()
			if tt.wantErr {
				assert.ErrorIs(t, err, core.ErrPopEmpty, "Dequeue() - Did not return an error")
				return
			}
			assert.EqualValues(t, tt.wantNode, gotN, "Dequeue() - Different node expected")
			assert.Exactly(t, tt.wantStackNodes, tt.stack.items, "Dequeue() - Remaining nodes do not match")
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		stack Stack[int]
		want  bool
	}{
		{
			name:  "Empty - true",
			stack: NewStack[int](),
			want:  true,
		},
		{
			name:  "Empty - false",
			stack: NewStack[int](core.NewNode[int](1, 12432432)),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equalf(t, tt.want, tt.stack.IsEmpty(), "IsEmpty()")
		})
	}
}
