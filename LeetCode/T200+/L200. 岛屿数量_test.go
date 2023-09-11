package T200_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/9/10 16:14
// @Version 1.0

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{"t1", args{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
		}, 1},
		{"t2", args{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
		}, 3},
		{"t3", args{
			[][]byte{
				{'1', '1', '1'},
				{'0', '1', '0'},
				{'1', '1', '1'},
			},
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := numIslands(tt.args.grid); gotRet != tt.wantRet {
				t.Errorf("numIslands() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
