package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"A", 95437},
		{"B", 24933642},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file = fmt.Sprintf("test%s", tt.name)
			if got := getResult(tt.name); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalSizes(t *testing.T) {
	tests := []struct {
		name string
		want map[string]int
	}{
		{"A", map[string]int{"/": 48381165, "/a/": 94853, "/d/": 24933642, "/a/e/": 584}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			file = fmt.Sprintf("test%s", tt.name)
			input := getInput()
			root := ConstructFilesystem(input)
			dirSizes := root.SizeMap()
			assert.Equal(t, tt.want, dirSizes)
		})
	}
}
