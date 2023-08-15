package T700_

import (
	"reflect"
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2023/8/15 23:02
// @Version 1.0

func Test_dailyTemperatures(t *testing.T) {
	type args struct {
		temperatures []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		{"t1", args{temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73}}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := dailyTemperatures(tt.args.temperatures); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("dailyTemperatures() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
