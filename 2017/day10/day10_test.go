package day10

import "testing"

func TestGetResult(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"A", 12},
		// {"B", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file = "day10.test"
			listLength = 5
			if got := GetResult(tt.name); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
