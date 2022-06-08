package T700_

import (
	"testing"
)

func Test_searchV2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "t1", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, want: 4},
		{name: "t2", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchV2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{name: "t1", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, wantRet: 4},
		{name: "t2", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2}, wantRet: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := search(tt.args.nums, tt.args.target); gotRet != tt.wantRet {
				t.Errorf("search() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_searchRepeat(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{name: "t1", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, wantRet: 4},
		{name: "t2", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2}, wantRet: -1},
		{name: "t2", args: args{nums: []int{5}, target: 5}, wantRet: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := searchRepeat2(tt.args.nums, tt.args.target); gotRet != tt.wantRet {
				t.Errorf("searchRepeat1() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
