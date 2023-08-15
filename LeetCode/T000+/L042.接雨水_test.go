package T000_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/15 23:30
// @Version 1.0

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"t1", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := trap(tt.args.height); gotRet != tt.wantRet {
				t.Errorf("trap() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
