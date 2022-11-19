package T000_

import "math"

// https://leetcode.cn/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var check func(node *TreeNode, min, max int) bool

	check = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}

		if node.Val <= min || node.Val >= max {
			return false
		}

		return check(node.Right, node.Val, max) && check(node.Left, min, node.Val)
	}

	return check(root, math.MinInt64, math.MaxInt64)
}
