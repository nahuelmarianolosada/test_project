package main

import (
	"reflect"
	"testing"
)

func Test_breakingRecords(t *testing.T) {
	type args struct {
		scores []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{"Debería retornar una vez para el maximo, y una para el minimo", args{[]int32{12, 24, 10, 24}}, []int32{1, 1}},
		{"Debería retornar 2 veces para el maximo, y 4 para el minimo", args{[]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}}, []int32{2, 4}},
		{"Debería retornar 4 veces para el maximo, y 0 para el minimo", args{[]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}}, []int32{4, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakingRecords(tt.args.scores); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("breakingRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
