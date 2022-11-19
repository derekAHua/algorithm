package T100_

// https://leetcode.cn/problems/symmetric-tree/

func isSymmetric(root *TreeNode) bool {

	return check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return check(left.Left, right.Right) && check(left.Right, right.Left)
}
