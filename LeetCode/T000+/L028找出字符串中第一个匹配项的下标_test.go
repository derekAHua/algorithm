package T000_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/11/13 18:46
// @Version 1.0

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"t1", args{"sadbutsad", "sad"}, 0},
		{"t2", args{"a", "a"}, 0},
		{"t2", args{"ababaabbbbababbaabaaabaabbaaaabbabaabbbbbbabbaabbabbbabbbbbaaabaababbbaabbbabbbaabbbbaaabbababbabbbabaaabbaabbabababbbaaaaaaababbabaababaabbbbaaabbbabb", "abbabbbabaa"}, 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
