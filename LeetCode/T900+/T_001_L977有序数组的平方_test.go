package T900_

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{nums: []int{-4, -1, 0, 3, 10}}, []int{0, 1, 9, 16, 100}},
		{"t2", args{nums: []int{-7, -3, 2, 3, 11}}, []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedSquaresRepeat1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		{"t1", args{nums: []int{-4, -1, 0, 3, 10}}, []int{0, 1, 9, 16, 100}},
		{"t2", args{nums: []int{-7, -3, 2, 3, 11}}, []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := sortedSquaresRepeat1(tt.args.nums); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("sortedSquaresRepeat1() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
