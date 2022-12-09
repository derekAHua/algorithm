package T400_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/12/4 12:29
// @Version 1.0

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := eraseOverlapIntervals(tt.args.intervals); gotRet != tt.wantRet {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
