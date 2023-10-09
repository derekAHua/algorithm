package T100_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/11/15 22:46
// @Version 1.0

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"t1", args{tokens: []string{"4", "13", "5", "/", "+"}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := evalRPN(tt.args.tokens); gotRet != tt.wantRet {
				t.Errorf("evalRPN() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
