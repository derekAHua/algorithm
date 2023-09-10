package T700_

import (
	"reflect"
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2023/9/3 15:04
// @Version 1.0

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRet [][]int
	}{
		// TODO: Add test cases.
		{"t1", args{[][]int{
			{1, 2},
			{3},
			{3},
			{},
		}}, [][]int{
			{0, 1, 3},
			{0, 2, 3},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
