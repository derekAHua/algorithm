package T600_

// https://leetcode.cn/problems/merge-two-binary-trees/

func mergeTrees(root1 *TreeNode, root2 *TreeNode) (ret *TreeNode) {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	ret = &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}

	return
}
