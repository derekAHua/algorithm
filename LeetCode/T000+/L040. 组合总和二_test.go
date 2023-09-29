package T000_

import (
	"reflect"
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2023/9/16 11:22
// @Version 1.0

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name    string
		args    args
		wantRet [][]int
	}{
		{"t1", args{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
		}, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("combinationSum2() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
