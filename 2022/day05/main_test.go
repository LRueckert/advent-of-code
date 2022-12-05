package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetResult(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"A", "CMZ"},
		{"B", "MCD"},
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

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name string
		want Command
	}{
		{"move 19 from 8 to 6", Command{19, 8, 6}},
		{"move 3 from 5 to 8", Command{3, 5, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCommand(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
