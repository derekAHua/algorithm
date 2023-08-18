package T400_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/12/10 13:32
// @Version 1.0

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet bool
	}{
		{"", args{[]int{1, 2, 5}}, false},
		{"", args{[]int{1, 5, 11, 5}}, true},
		{"", args{[]int{2, 2, 1, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := canPartition(tt.args.nums); gotRet != tt.wantRet {
				t.Errorf("canPartition() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
