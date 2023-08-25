package T100_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/19 08:55
// @Version 1.0

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{"t1", args{[]int{1, 2, 3, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := rob(tt.args.nums); gotRet != tt.wantRet {
				t.Errorf("rob() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
