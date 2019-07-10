package day8

import "testing"

func TestGetResult(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"A", 1},
		{"B", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file = "day8.test"
			if got := GetResult(tt.name); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
