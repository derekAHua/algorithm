package T300_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/16 09:18
// @Version 1.0

func Test_integerBreak(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"t1", args{10}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := integerBreak(tt.args.n); gotRet != tt.wantRet {
				t.Errorf("integerBreak() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
