package day6

import (
	"fmt"
	"testing"
)

func TestGetResult(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"A", 17},
		{"B", 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file = fmt.Sprintf("day6%s.test", tt.name)
			distanceLimit = 32
			if got := GetResult(tt.name); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
