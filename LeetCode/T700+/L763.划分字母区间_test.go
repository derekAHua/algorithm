package T700_

import (
	"reflect"
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/12/4 12:53
// @Version 1.0

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		{"", args{"eaaaabaaec"}, []int{9, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := partitionLabels(tt.args.s); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("partitionLabels() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
