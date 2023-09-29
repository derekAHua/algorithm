package T300_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/26 21:28
// @Version 1.0

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{"t1", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"t2", args{[]int{4, 10, 4, 3, 8, 9}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := lengthOfLIS(tt.args.nums); gotRet != tt.wantRet {
				t.Errorf("lengthOfLIS() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
