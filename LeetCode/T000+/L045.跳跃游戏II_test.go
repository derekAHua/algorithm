package T000_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/12/3 11:32
// @Version 1.0

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"", args{[]int{2, 3, 1, 1, 4}}, 2},
		{"", args{[]int{1, 2, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := jump(tt.args.nums); gotRet != tt.wantRet {
				t.Errorf("jump() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
