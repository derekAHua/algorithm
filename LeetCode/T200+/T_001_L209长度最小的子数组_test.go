package T200_

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{target: 7, nums: []int{2, 3, 1, 2, 4, 3}}, 2},
		{"t2", args{target: 4, nums: []int{1, 4, 4}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSubArrayLenRepeat1(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"t1", args{target: 7, nums: []int{2, 3, 1, 2, 4, 3}}, 2},
		{"t2", args{target: 4, nums: []int{1, 4, 4}}, 1},
		{"t3", args{target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := minSubArrayLenRepeat1(tt.args.target, tt.args.nums); gotRet != tt.wantRet {
				t.Errorf("minSubArrayLenRepeat1() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
