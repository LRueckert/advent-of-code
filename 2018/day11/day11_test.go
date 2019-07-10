package day11

import (
	"testing"
)

func TestGetResult(t *testing.T) {
	tests := []struct {
		name         string
		part         string
		serialNumber int
		want         string
	}{
		{"A-1", "A", 18, "33,45"},
		// {"A-2", "A", 42, "21,61"},
		// {"B-1", "B", 18, "90,269,16"},
		// {"B-2", "B", 42, "232,251,12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serialNumber = tt.serialNumber
			if got := GetResult(tt.part); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateCellPower(t *testing.T) {
	type args struct {
		x            int
		y            int
		serialNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{122, 79, 57}, -5},
		{"2", args{217, 196, 39}, 0},
		{"3", args{101, 153, 71}, 4},
		{"4", args{3, 5, 8}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serialNumber = tt.args.serialNumber
			if got := calculateCellPower(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("calculateCellPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
