package T100_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/12/17 10:57
// @Version 1.0

func Test_maxProfit_188(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit_188(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit_188() = %v, want %v", got, tt.want)
			}
		})
	}
}
