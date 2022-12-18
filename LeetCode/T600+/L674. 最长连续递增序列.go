package T600_

// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/

func findLengthOfLCIS(nums []int) (ret int) {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ret = 1

	for i := 1; i < len(dp); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}

		if dp[i] > ret {
			ret = dp[i]
		}
	}

	return
}
