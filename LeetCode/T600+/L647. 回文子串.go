package T600_

// https://leetcode.cn/problems/palindromic-substrings/

func countSubstrings(s string) (ret int) {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] != s[j] {
				continue
			}

			if j-i <= 1 {
				ret++
				dp[i][j] = true
				continue
			}

			if dp[i+1][j-1] {
				ret++
				dp[i][j] = true
			}
		}
	}

	return
}
