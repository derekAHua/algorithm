package T400_

import "math"

// https://leetcode.cn/problems/target-sum/

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if int(math.Abs(float64(target))) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}

	bag := (sum + target) / 2
	dp := make([]int, bag+1)
	dp[0] = 1
	for _, v := range nums {
		for i := bag; i >= v; i-- {
			dp[i] += dp[i-v]
		}
	}

	return dp[bag]
}
