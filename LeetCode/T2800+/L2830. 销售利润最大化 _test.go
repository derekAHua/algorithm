package LeetCode

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/26 10:11
// @Version 1.0

func Test_maximizeTheProfit(t *testing.T) {
	type args struct {
		n      int
		offers [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"t1", args{n: 5, offers: [][]int{{0, 0, 1}, {0, 2, 2}, {1, 3, 2}}}, 3},
		{"t1", args{n: 5, offers: [][]int{{0, 0, 1}, {0, 2, 10}, {1, 3, 2}}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeTheProfit(tt.args.n, tt.args.offers); got != tt.want {
				t.Errorf("maximizeTheProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
