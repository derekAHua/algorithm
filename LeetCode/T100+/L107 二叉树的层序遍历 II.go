package T100_

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/

// 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

func levelOrderBottom(root *TreeNode) (ret [][]int) {
	if root == nil {
		return
	}

	var q []*TreeNode
	q = append(q, root)

	for len(q) != 0 {
		r := make([]int, 0)
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[i]
			r = append(r, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[length:]
		ret = append(ret, r)
	}

	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}

	return
}
