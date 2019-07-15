package main

import (
	"testing"
)

func TestGetResult(t *testing.T) {
	tests := []struct {
		name    string
		part    string
		players int
		marbles int
		want    int
	}{
		{"A1", "A", 9, 25, 32},
		{"A2", "A", 10, 1618, 8317},
		{"A3", "A", 13, 7999, 146373},
		{"A3", "A", 17, 1104, 2764},
		{"A4", "A", 21, 6111, 54718},
		{"A4", "A", 30, 5807, 37305},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numPlayers = tt.players
			lastMarble = tt.marbles
			if got := getResult(tt.part); got != tt.want {
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
