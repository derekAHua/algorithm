package T000_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/9/24 10:05
// @Version 1.0

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{[]int{2, 0, 2, 1, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
		})
	}
}
