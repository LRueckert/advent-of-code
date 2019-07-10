package day1

import "testing"

func Test_calculateResult(t *testing.T) {
	type args struct {
		part  string
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A-1", args{"A", []string{"+1", "+1", "+1"}}, 3},
		{"A-2", args{"A", []string{"+1", "+1", "-2"}}, 0},
		{"A-3", args{"A", []string{"-1", "-2", "-3"}}, -6},
		{"B-1", args{"B", []string{"+1", "-1"}}, 0},
		{"B-2", args{"B", []string{"+3", "+3", "+4", "-2", "-4"}}, 10},
		{"B-3", args{"B", []string{"-6", "+3", "+8", "+5", "-6"}}, 5},
		{"B-4", args{"B", []string{"+7", "+7", "-2", "-7", "-4"}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateResult(tt.args.part, tt.args.input); got != tt.want {
				t.Errorf("calculateResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
