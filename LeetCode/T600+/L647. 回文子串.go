package T600_

// https://leetcode.cn/problems/palindromic-substrings/

func countSubstrings(s string) (ret int) {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for l := len(s) - 1; l >= 0; l-- {
		for r := l; r < len(s); r++ {
			if s[l] != s[r] {
				continue
			}

			if r-l <= 1 {
				ret++
				dp[l][r] = true
				continue
			}

			if dp[l+1][r-1] {
				ret++
				dp[l][r] = true
			}
		}
	}

	return
}

func countSubstringsR1(s string) (ret int) {
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
