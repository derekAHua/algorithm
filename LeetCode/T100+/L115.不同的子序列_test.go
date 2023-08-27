package T100_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/27 22:21
// @Version 1.0

func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{"t1", args{"rabbbit", "rabit"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := numDistinct(tt.args.s, tt.args.t); gotRet != tt.wantRet {
				t.Errorf("numDistinct() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
