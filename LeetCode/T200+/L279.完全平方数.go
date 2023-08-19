package T200_

import (
	"math"
)

// https://leetcode.cn/problems/perfect-squares/

func numSquares(n int) int {
	nums := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		for _, v := range nums {
			if i < v {
				continue
			}
			if dp[i] > dp[i-v]+1 {
				dp[i] = dp[i-v] + 1
			}
		}
	}

	return dp[n]
}

func numSquaresR1(n int) int {
	nums := make([]int, 0, 100)
	for i := 1; i < 100; i++ {
		x := i * i
		if x > n {
			break
		}
		nums = append(nums, x)
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}

	for _, v := range nums {
		for i := v; i < len(dp); i++ {
			dp[i] = min(dp[i], dp[i-v]+1)
		}
	}

	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
