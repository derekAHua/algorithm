package T700_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/27 21:20
// @Version 1.0

func Test_findLength(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{"t1", args{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}}, 3},
		{"t1", args{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}}, 5},
		{"t1", args{[]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := findLength(tt.args.nums1, tt.args.nums2); gotRet != tt.wantRet {
				t.Errorf("findLength() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
