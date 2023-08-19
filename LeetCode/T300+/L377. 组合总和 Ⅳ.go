package T300_

// https://leetcode.cn/problems/combination-sum-iv/

func combinationSum4(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		for _, num := range nums {
			if i < num {
				continue
			}
			dp[i] += dp[i-num]
		}
	}

	return dp[target]
}

func combinationSum4R1(nums []int, target int) int {

	// 总和为 target 的元素组合的个数
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 0; i <= target; i++ {
		for _, v := range nums {
			if i >= v {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}
