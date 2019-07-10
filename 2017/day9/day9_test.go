package day9

import "testing"

func Test_calculateResult(t *testing.T) {
	type args struct {
		part  string
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A-1", args{"A", "{}"}, 1},
		{"A-2", args{"A", "{{{}}}"}, 6},
		{"A-3", args{"A", "{{},{}}"}, 5},
		{"A-4", args{"A", "{{{},{},{{}}}}"}, 16},
		{"A-5", args{"A", "{<a>,<a>,<a>,<a>}"}, 1},
		{"A-6", args{"A", "{{<ab>},{<ab>},{<ab>},{<ab>}}"}, 9},
		{"A-7", args{"A", "{{<!!>},{<!!>},{<!!>},{<!!>}}"}, 9},
		{"A-8", args{"A", "{{<a!>},{<a!>},{<a!>},{<ab>}}"}, 3},
		{"B-1", args{"B", "<>"}, 0},
		{"B-2", args{"B", "<random characters>"}, 17},
		{"B-3", args{"B", "<<<<>"}, 3},
		{"B-4", args{"B", "<{!>}>"}, 2},
		{"B-5", args{"B", "<!!>"}, 0},
		{"B-6", args{"B", "<!!!>>"}, 0},
		{"B-7", args{"B", "<{o'i!a,<{i<a>}"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateResult(tt.args.part, tt.args.input); got != tt.want {
				t.Errorf("calculateResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
