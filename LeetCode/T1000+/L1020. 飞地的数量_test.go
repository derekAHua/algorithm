package T1000_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/10/12 21:09
// @Version 1.0

func Test_numEnclaves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			name: "t1",
			args: args{grid: [][]int{
				{0, 0, 0, 0},
				{1, 0, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			}},
			wantRet: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := numEnclaves(tt.args.grid); gotRet != tt.wantRet {
				t.Errorf("numEnclaves() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
