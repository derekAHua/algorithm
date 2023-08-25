package T200_

// https://leetcode.cn/problems/house-robber-ii/

func rob(nums []int) (ret int) {
	if len(nums) == 1 {
		return nums[0]
	}

	v1 := robRR(nums[:len(nums)-1])
	v2 := robRR(nums[1:len(nums)])
	return max(v1, v2)
}

func robRR(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])

	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}

func robR1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	result1 := robRange(nums, 0)
	result2 := robRange(nums, 1)
	return max(result1, result2)
}

// 偷盗指定的范围
func robRange(nums []int, start int) int {
	dp := make([]int, len(nums))
	dp[1] = nums[start]

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i-1+start], dp[i-1])
	}

	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
