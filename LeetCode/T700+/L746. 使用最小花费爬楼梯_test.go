package T700_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/15 23:55
// @Version 1.0

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"t1", args{[]int{10, 15, 20}}, 15},
		{"t1", args{[]int{10, 15}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := minCostClimbingStairs(tt.args.cost); gotRet != tt.wantRet {
				t.Errorf("minCostClimbingStairs() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
