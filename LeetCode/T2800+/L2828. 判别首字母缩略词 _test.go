package LeetCode

import "testing"

// @Author: Derek
// @Description:
// @Date: 2023/8/26 09:39
// @Version 1.0

func Test_isAcronym(t *testing.T) {
	type args struct {
		words []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{words: []string{"alice", "bob", "charlie"}, s: "abc"}, true},
		{"t2", args{words: []string{"a", "b", "c"}, s: "abcd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAcronym(tt.args.words, tt.args.s); got != tt.want {
				t.Errorf("isAcronym() = %v, want %v", got, tt.want)
			}
		})
	}
}
