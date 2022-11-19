package T600_

// https://leetcode.cn/problems/trim-a-binary-search-tree/

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low { //如果该节点值小于最小值，则该节点更换为该节点的右节点值，继续遍历
		return trimBST(root.Right, low, high)
	}
	if root.Val > high { //如果该节点的值大于最大值，则该节点更换为该节点的左节点值，继续遍历
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
