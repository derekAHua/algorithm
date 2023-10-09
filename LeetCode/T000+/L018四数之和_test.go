package T000_

import (
	"reflect"
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/12 22:56
// @Version 1.0

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t1", args{[]int{2, 2, 2, 2, 2}, 8}, [][]int{{2, 2, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
