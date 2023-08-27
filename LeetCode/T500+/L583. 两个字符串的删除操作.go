package T500_

// https://leetcode.cn/problems/delete-operation-for-two-strings/

func minDistance(word1 string, word2 string) (ret int) {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return len(word1) + len(word2) - 2*dp[len(word1)][len(word2)]
}

func minDistanceR1(word1 string, word2 string) int {

	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return len(word1) - dp[len(word1)][len(word2)]*2 + len(word2)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
