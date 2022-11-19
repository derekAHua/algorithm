package T200_

import "strconv"

// https://leetcode.cn/problems/binary-tree-paths/

func binaryTreePaths(root *TreeNode) (ret []string) {

	var f func(root *TreeNode, s string)

	f = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			ret = append(ret, s+strconv.Itoa(node.Val))
			return
		}

		if node.Left != nil {
			f(node.Left, s+strconv.Itoa(node.Val)+"->")
		}
		if node.Right != nil {
			f(node.Right, s+strconv.Itoa(node.Val)+"->")
		}
	}

	f(root, "")

	return
}

func binaryTreePaths2(root *TreeNode) (ret []string) {
	if root == nil {
		return
	}
	var travel func(node *TreeNode, s string)

	travel = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			ret = append(ret, v)
			return
		}

		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}

	travel(root, "")
	return
}
