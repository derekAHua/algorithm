package T100_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/9/14 22:31
// @Version 1.0

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{"t1", args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}}, 9},
		{"t1", args{nums: []int{1, 2, 0, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := longestConsecutive(tt.args.nums); gotRet != tt.wantRet {
				t.Errorf("longestConsecutive() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
