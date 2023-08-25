package T200_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/19 09:15
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
		{"t1", args{[]int{2, 3, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := rob(tt.args.nums); gotRet != tt.wantRet {
				t.Errorf("rob() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
