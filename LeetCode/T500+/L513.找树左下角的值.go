package T500_

import "container/list"

// https://leetcode.cn/problems/find-bottom-left-tree-value/

//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
//假设二叉树中至少有一个节点。

func findBottomLeftValue(root *TreeNode) (ret int) {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	min := 0

	var f func(node *TreeNode, depth int)
	f = func(node *TreeNode, depth int) {

		if depth > min && node.Left == nil && node.Right == nil {
			ret = node.Val
			min = depth
		}

		if node.Left != nil {
			f(node.Left, depth+1)
		}
		if node.Right != nil {
			f(node.Right, depth+1)
		}
	}

	f(root, 1)

	return
}

func findBottomLeftValue2(root *TreeNode) (ret int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			v := queue.Remove(queue.Front()).(*TreeNode)
			if v.Left != nil {
				queue.PushBack(v.Left)
			}
			if v.Right != nil {
				queue.PushBack(v.Right)
			}

			if i == 0 {
				ret = v.Val
			}
		}
	}
	return
}
