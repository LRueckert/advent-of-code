package day7

import (
	"fmt"
	"testing"
)

func TestGetResult(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"A", "CABDFE"},
		{"B", "15"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file = fmt.Sprintf("day7%s.test", tt.name)
			baseDuration = 0
			numWorkers = 2
			if got := GetResult(tt.name); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
