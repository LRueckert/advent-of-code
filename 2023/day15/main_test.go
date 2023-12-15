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
		{"A", 1320},
		{"B", 145},
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

func TestHashFunction(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"rn=1", 30},
		{"cm-", 253},
		{"qp=3", 97},
		{"cm=2", 47},
		{"qp-", 14},
		{"pc=4", 180},
		{"ot=9", 9},
		{"ab=5", 197},
		{"pc-", 48},
		{"pc=6", 214},
		{"ot=7", 231},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashFunction(tt.name); got != tt.want {
				t.Errorf("HashFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
