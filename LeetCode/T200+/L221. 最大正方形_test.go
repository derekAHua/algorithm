package T200_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/9/14 23:10
// @Version 1.0

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			"t1",
			args{matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := maximalSquare(tt.args.matrix); gotRet != tt.wantRet {
				t.Errorf("maximalSquare() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
