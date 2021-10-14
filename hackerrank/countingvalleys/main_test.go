package main

import "testing"

func Test_countingValleys(t *testing.T) {
	type args struct {
		steps int32
		path  string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"Debe retornar 1 valle", args{8, "UDDDUDUU"}, 1},
		{"Debe retornar 2 valles", args{12, "DDUUDDUDUUUD"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingValleys(tt.args.steps, tt.args.path); got != tt.want {
				t.Errorf("countingValleys() = %v, want %v", got, tt.want)
			}
		})
	}
}
