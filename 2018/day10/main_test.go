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
		{"A", 0},
		{"B", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file = fmt.Sprintf("day10%s.test", tt.name)
			if got := getResult(tt.name); got != tt.want {
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
