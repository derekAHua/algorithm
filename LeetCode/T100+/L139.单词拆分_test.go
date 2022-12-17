package T100_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/12/11 16:25
// @Version 1.0

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{s: "leetcode", wordDict: []string{"leet", "code"}}, true},
		{"", args{s: "applepenapple", wordDict: []string{"apple", "pen"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
