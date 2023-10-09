package T700_

import "sort"

// @Author: Derek
// @Description:
// @Date: 2023/9/29 09:58
// @Version 1.0

//给你一个整数数组 nums ，你可以对它进行一些操作。
//每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。
//开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。

// todo -》打家劫舍 need to see
func deleteAndEarn(nums []int) (ans int) {
	sort.Ints(nums)

	sum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		val := nums[i]

		if val == nums[i-1] {
			sum[len(sum)-1] += val
		} else if val == nums[i-1]+1 {
			sum = append(sum, val)
		} else {
			ans += rob(sum)
			sum = []int{val}
		}
	}
	ans += rob(sum)
	return
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(dp)-1]
}
