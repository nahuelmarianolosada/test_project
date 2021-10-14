package main

import "testing"

func Test_getTotalX(t *testing.T) {
	type args struct {
		a []int32
		b []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"Debe retornar 2", args{[]int32{2, 6}, []int32{24, 36}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTotalX(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getTotalX() = %v, want %v", got, tt.want)
			}
		})
	}
}
