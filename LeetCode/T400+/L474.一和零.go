package T400_

// https://leetcode.cn/problems/ones-and-zeroes/

func findMaxForm(strs []string, m int, n int) int {

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, v := range strs {
		num1, num0 := 0, 0
		for _, vv := range v {
			switch vv {
			case '0':
				num0++
			default:
				num1++
			}
		}
		for i := m; i >= num0; i-- {
			for j := n; j >= num1; j-- {
				if dp[i][j] < dp[i-num0][j-num1]+1 {
					dp[i][j] = dp[i-num0][j-num1] + 1
				}
			}
		}
	}

	return dp[m][n]
}

func findMaxFormR1(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {

		zeroNum, oneNum := 0, 0
		for _, v := range strs[i] {
			if v == '0' {
				zeroNum++
			}
		}
		oneNum = len(strs[i]) - zeroNum

		// 从后往前 遍历背包容量
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				// 推导公式
				dp[j][k] = max(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
