package T000_

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t1", args{n: 3}, [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		}},
		{"t2", args{n: 1}, [][]int{
			{1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateMatrixRepeat1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRet [][]int
	}{
		{"t1", args{n: 3}, [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		}},
		{"t2", args{n: 1}, [][]int{
			{1},
		}},
		{"t3", args{n: 2}, [][]int{
			{1, 2},
			{4, 3},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := generateMatrixRepeat3(tt.args.n); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("generateMatrixRepeat1() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
