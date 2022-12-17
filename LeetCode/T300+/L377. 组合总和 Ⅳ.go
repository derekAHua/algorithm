package T300_

// https://leetcode.cn/problems/combination-sum-iv/

// 1,2,3 target=4

// 0 1 2 3 4

// 1 1 1 1 1
// 1 1 2 2 3

func combinationSum4(nums []int, target int) int {

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
