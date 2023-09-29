package T000_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/9/27 22:00
// @Version 1.0

func Test_exist(t *testing.T) {
	type args struct {
		arr    [][]byte
		target string
	}
	tests := []struct {
		name    string
		args    args
		wantRet bool
	}{
		{"t1", args{
			arr:    [][]byte{{'a'}},
			target: "a",
		}, true},
		{"t1", args{
			arr: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'E', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			target: "ABCESEEEFS",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := exist(tt.args.arr, tt.args.target); gotRet != tt.wantRet {
				t.Errorf("exist() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
