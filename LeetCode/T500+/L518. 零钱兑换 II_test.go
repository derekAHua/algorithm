package T500_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/12/11 13:14
// @Version 1.0

func Test_change(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{
			amount: 5,
			coins:  []int{1, 2, 5},
		}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
		})
	}
}
