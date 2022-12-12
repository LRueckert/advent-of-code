package main

import (
	"fmt"
	"testing"
)

func TestGetResult(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"A", 31},
		{"B", 29},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file = fmt.Sprintf("test%s.txt", tt.name)
			if got := getResult(tt.name); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Connects(t *testing.T) {
	tests := []struct {
		name  string
		node  Node
		other Node
		want  bool
	}{
		{"one up", Node{Elevation: "a"}, Node{Elevation: "b"}, true},
		{"two up", Node{Elevation: "a"}, Node{Elevation: "c"}, false},
		{"max down", Node{Elevation: "z"}, Node{Elevation: "a"}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.Connects(&tt.other); got != tt.want {
				t.Errorf("Node.Connects() = %v, want %v", got, tt.want)
			}
		})
	}
}
