package T400_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/16 23:48
// @Version 1.0

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"t1", args{[]int{1, 1, 1, 1, 1}, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := findTargetSumWays(tt.args.nums, tt.args.target); gotRet != tt.wantRet {
				t.Errorf("findTargetSumWays() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
