package T400_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/11/20 22:16
// @Version 1.0

func Test_findContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"", args{
			g: []int{1, 2, 3},
			s: []int{1, 1},
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := findContentChildren(tt.args.g, tt.args.s); gotRet != tt.wantRet {
				t.Errorf("findContentChildren() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
