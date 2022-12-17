package T100_

// https://leetcode.cn/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s)+1; i++ {
		for _, v := range wordDict {
			l := len(v)
			if i-l < 0 {
				continue
			}
			cur := s[i-l : i]
			if dp[i-l] && cur == v {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}
