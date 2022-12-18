package T300_

// https://leetcode.cn/problems/longest-increasing-subsequence/

func lengthOfLIS(nums []int) (ret int) {
	dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = 1
	}

	ret = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return
}
