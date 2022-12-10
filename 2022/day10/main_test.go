package main

import (
	"fmt"
	"testing"
)

func TestGetResult(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"A", "13140"},
		{"B", "##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file = fmt.Sprintf("test%s", tt.name)
			if got := getResult(tt.name); got != tt.want {
				t.Errorf("GetResult() =\n%vwant\n%v", got, tt.want)
			}
		})
	}
}
