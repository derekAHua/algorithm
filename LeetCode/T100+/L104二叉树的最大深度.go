package T100_

import "container/list"

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

func maxDepth(root *TreeNode) (ret int) {
	var f func(root *TreeNode, deep int)
	f = func(root *TreeNode, deep int) {
		if root == nil {
			if deep > ret {
				ret = deep - 1
			}
			return
		}

		f(root.Left, deep+1)
		f(root.Right, deep+1)
	}

	f(root, 1)
	return
}

func maxDepthR1(root *TreeNode) (ret int) {
	if root == nil {
		return
	}

	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		length := q.Len()
		for i := 0; i < length; i++ {
			node := q.Remove(q.Front()).(*TreeNode)
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		ret++
	}

	return
}
