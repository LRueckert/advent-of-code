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
		{"A", 142},
		{"B", 281},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file = fmt.Sprintf("test%s.txt", tt.name)
			if got := getResult(tt.name); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLineValue(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{"rkzlnmzgnk91zckqprrptnthreefourtwo"}, 92},
		{"overlapping", args{"vnjz8onemdnjfzmqcgxqonemspchrxlggoneightptg"}, 88},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLineValue(tt.args.line, true); got != tt.want {
				t.Errorf("GetLineValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
