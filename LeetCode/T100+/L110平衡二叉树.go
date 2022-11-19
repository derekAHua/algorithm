package T100_

import "math"

// https://leetcode.cn/problems/balanced-binary-tree/

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
//本题中，一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	return math.Abs(float64(mDepth(root.Left)-mDepth(root.Right))) <= 1
}

func mDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(mDepth(root.Left), mDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
