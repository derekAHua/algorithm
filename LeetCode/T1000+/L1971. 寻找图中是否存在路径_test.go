package T1000_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/26 20:01
// @Version 1.0

func Test_validPath(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{
			n:           10,
			edges:       [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}},
			source:      7,
			destination: 5,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPath(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination); got != tt.want {
				t.Errorf("validPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
