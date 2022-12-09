package T800_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/12/3 14:51
// @Version 1.0

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet bool
	}{
		{"", args{bills: []int{5, 5, 10, 10, 20}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := lemonadeChange(tt.args.bills); gotRet != tt.wantRet {
				t.Errorf("lemonadeChange() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
