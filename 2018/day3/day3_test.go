package day3

import (
	"fmt"
	"testing"
)

func TestGetResult(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"A", 4},
		{"B", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file = fmt.Sprintf("day3%s.test", tt.name)
			if got := GetResult(tt.name); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
