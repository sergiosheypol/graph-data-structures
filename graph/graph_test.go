package graph

import (
	"testing"

	"graph-data-structures/core"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	type args struct {
		directed bool
	}
	tests := []struct {
		name string
		args args
		want Graph[string]
	}{
		{
			name: "happy path",
			args: args{directed: true},
			want: Graph[string]{
				adjacencyList: map[core.NodeId][]core.NodeId{},
				nodes:         map[core.NodeId]core.Node[string]{},
				directed:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewGraph[string](tt.args.directed), "NewGraph(%v)", tt.args.directed)
		})
	}
}

func TestGraph_AddNode(t *testing.T) {
	type args struct {
		node core.Node[string]
	}
	tests := []struct {
		name    string
		graph   Graph[string]
		args    args
		wantErr bool
	}{
		{
			name: "Ok",
			graph: Graph[string]{
				adjacencyList: map[core.NodeId][]core.NodeId{},
				nodes:         map[core.NodeId]core.Node[string]{},
				directed:      false,
			},
			args: args{
				node: core.NewNode[string](1, "blue"),
			},
			wantErr: false,
		},
		{
			name: "Duplicated node",
			graph: Graph[string]{
				adjacencyList: map[core.NodeId][]core.NodeId{},
				nodes: map[core.NodeId]core.Node[string]{
					1: core.NewNode[string](1, "blue"),
				},
				directed: true,
			},
			args: args{
				node: core.NewNode[string](1, "blue"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.graph.AddNode(tt.args.node)
			if tt.wantErr {
				assert.ErrorIs(t, err, core.ErrDuplicatedNode, "No Error was returned")
				return
			}
			assert.EqualValues(t, tt.graph.nodes[tt.args.node.Id()], tt.args.node)
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {

	n1 := core.NewNode[int](1, 1)
	n2 := core.NewNode[int](2, 2)

	type fields struct {
		adjacencyList map[core.NodeId][]core.NodeId
		nodes         map[core.NodeId]core.Node[int]
		directed      bool
	}
	type args struct {
		from core.Node[int]
		to   core.Node[int]
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Ok - undirected",
			fields: fields{
				adjacencyList: make(map[core.NodeId][]core.NodeId),
				nodes: map[core.NodeId]core.Node[int]{
					1: n1,
					2: n2,
				},
				directed: false,
			},
			args: args{
				from: n1,
				to:   n2,
			},
			wantErr: false,
		},
		{
			name: "Ok - directed",
			fields: fields{
				adjacencyList: make(map[core.NodeId][]core.NodeId),
				nodes: map[core.NodeId]core.Node[int]{
					1: n1,
					2: n2,
				},
				directed: true,
			},
			args: args{
				from: n1,
				to:   n2,
			},
			wantErr: false,
		},
		{
			name: "Err - non existing nodes",
			fields: fields{
				adjacencyList: make(map[core.NodeId][]core.NodeId),
				nodes:         make(map[core.NodeId]core.Node[int]),
				directed:      false,
			},
			args: args{
				from: n1,
				to:   n2,
			},
			wantErr: true,
		},
		{
			name: "Err - n2 does not exist",
			fields: fields{
				adjacencyList: make(map[core.NodeId][]core.NodeId),
				nodes: map[core.NodeId]core.Node[int]{
					1: n1,
				},
				directed: false,
			},
			args: args{
				from: n1,
				to:   n2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph[int]{
				adjacencyList: tt.fields.adjacencyList,
				nodes:         tt.fields.nodes,
				directed:      tt.fields.directed,
			}

			err := g.AddEdge(tt.args.from, tt.args.to)

			if tt.wantErr {
				assert.Error(t, err, "No Error was returned")
				return
			}

			assert.Contains(t, tt.fields.adjacencyList[n1.Id()], n2.Id())

			if tt.fields.directed {
				assert.NotContains(t, tt.fields.adjacencyList[n2.Id()], n1.Id())
				return
			}

			assert.Contains(t, tt.fields.adjacencyList[n2.Id()], n1.Id())

		})
	}
}
