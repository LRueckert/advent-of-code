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
		{"A", 26},
		{"B", 56000011},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row = 10
			dimension = 20
			file = fmt.Sprintf("test%s.txt", tt.name)
			if got := getResult(tt.name); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		a    Point
		b    Point
		want int
	}{
		{"0", Point{}, Point{}, 0},
		{"1", Point{1, 1}, Point{2, 2}, 2},
		{"2", Point{1, 4}, Point{-2, 16}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.a, tt.b); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
