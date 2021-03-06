package T000_

import (
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{nums: []int{3, 2, 2, 3}, val: 3}, 2},
		{"t2", args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeElementRepeat1(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{nums: []int{3, 2, 2, 3}, val: 3}, 2},
		{"t2", args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElementRepeat1(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElementRepeat1() = %v, want %v", got, tt.want)
			}
		})
	}
}
