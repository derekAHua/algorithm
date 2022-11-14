package T000_

import (
	"reflect"
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/12 17:49
// @Version 1.0

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"t2", args{[]int{0, 0, 0, 0}}, [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})

	}
}
