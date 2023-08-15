package T200_

// https://leetcode.cn/problems/invert-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) (ret *TreeNode) {
	var f func(root *TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}

		f(root.Left)
		f(root.Right)
		root.Left, root.Right = root.Right, root.Left
	}
	f(root)
	return root
}

func invertTreeR1(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	if root.Left != nil {
		invertTree(root.Left)
	}

	if root.Right != nil {
		invertTree(root.Right)
	}

	return root
}
