package T700_

import (
	"testing"
)

func TestSearch(t *testing.T) {
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
				t.Errorf("Search() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
