package T000_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/9/9 20:35
// @Version 1.0

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{"t1", args{grid: [][]int{
			{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
		}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := minPathSum(tt.args.grid); gotRet != tt.wantRet {
				t.Errorf("minPathSum() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
