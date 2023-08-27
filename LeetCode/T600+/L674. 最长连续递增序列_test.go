package T600_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/27 20:12
// @Version 1.0

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{"t1", args{nums: []int{1, 3, 5, 4, 7}}, 3},
		{"t1", args{nums: []int{1, 3, 5, 7}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := findLengthOfLCIS(tt.args.nums); gotRet != tt.wantRet {
				t.Errorf("findLengthOfLCIS() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
