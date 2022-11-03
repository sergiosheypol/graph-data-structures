package core

import (
	"testing"
)

func TestNode_Equals_Int(t *testing.T) {
	tests := []struct {
		name string
		n    Node[int]
		n2   Node[int]
		want bool
	}{
		{
			name: "True Happy Path",
			n:    NewNode(1, 2),
			n2:   NewNode(1, 2),
			want: true,
		},
		{
			name: "False Wrong Value",
			n:    NewNode(1, 2),
			n2:   NewNode(1, 3),
			want: false,
		},
		{
			name: "False Wrong Id",
			n:    NewNode(1, 2),
			n2:   NewNode(2, 2),
			want: false,
		},
		{
			name: "False All Wrong",
			n:    NewNode(1, 2),
			n2:   NewNode(2, 15),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Equals(tt.n2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Equals_String(t *testing.T) {
	tests := []struct {
		name string
		n    Node[string]
		n2   Node[string]
		want bool
	}{
		{
			name: "True Happy Path",
			n:    NewNode(1, "blue"),
			n2:   NewNode(1, "blue"),
			want: true,
		},
		{
			name: "False Wrong Value",
			n:    NewNode(1, "blue"),
			n2:   NewNode(1, "red"),
			want: false,
		},
		{
			name: "False Wrong Id",
			n:    NewNode(1, "blue"),
			n2:   NewNode(2, "blue"),
			want: false,
		},
		{
			name: "False All Wrong",
			n:    NewNode(1, "blue"),
			n2:   NewNode(2, "red"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Equals(tt.n2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
