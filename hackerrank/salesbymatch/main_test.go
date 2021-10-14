package main

import "testing"

func Test_SockMerchant(t *testing.T) {
	type args struct {
		n  int32
		ar []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"Debe ser un solo par", args{1, []int32{1, 1}}, 1},
		{"Debe ser dos pares", args{4, []int32{1, 1, 2, 2}}, 2},
		{"Debe ser tres pares", args{9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}}, 3},
		{"Debe ser cuatro pares", args{10, []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SockMerchant(tt.args.n, tt.args.ar); got != tt.want {
				t.Errorf("SockMerchant() = %v, want %v", got, tt.want)
			}
		})
	}
}
