package T600_

import (
	"math"
)

// https://leetcode.cn/problems/maximum-binary-tree/

func constructMaximumBinaryTree(nums []int) (ret *TreeNode) {

	if len(nums) == 0 {
		return
	}

	value, index := max(nums)

	ret = &TreeNode{
		Val:   value,
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}

	return
}

func max(nums []int) (m, index int) {
	m = math.MinInt
	for i, v := range nums {
		if v > m {
			m = v
			index = i
		}
	}
	return
}
