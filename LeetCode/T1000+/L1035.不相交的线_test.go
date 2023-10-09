package T1000_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/27 21:59
// @Version 1.0

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"t1", args{[]int{1}, []int{1, 3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := maxUncrossedLines(tt.args.nums1, tt.args.nums2); gotRet != tt.wantRet {
				t.Errorf("maxUncrossedLines() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
