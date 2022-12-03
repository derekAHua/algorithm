package T000_

// https://leetcode.cn/problems/two-sum/

func canJump(nums []int) bool {
	cover := 0
	n := len(nums) - 1

	for i := 0; i <= cover; i++ {
		if i+nums[i] > cover {
			cover = nums[i] + i
		}
		if cover >= n {
			return true
		}
	}

	return false
}
