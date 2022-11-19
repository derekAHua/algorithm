package T500_

// https://leetcode.cn/problems/convert-bst-to-greater-tree/

func convertBST(root *TreeNode) *TreeNode {
	var sum int

	var f func(root *TreeNode) *TreeNode
	f = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		f(root.Right)
		sum += root.Val
		root.Val = sum
		f(root.Left)
		return root
	}

	f(root)

	return root
}
