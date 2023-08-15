package T400_

import "container/list"

// https://leetcode.cn/problems/sum-of-left-leaves/

//计算给定二叉树的所有左叶子之和。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) (ret int) {
	var f func(*TreeNode, bool)
	f = func(node *TreeNode, b bool) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil && b {
			ret += node.Val
		}

		f(node.Left, true)
		f(node.Right, false)
	}

	f(root, false)
	return
}

func sumOfLeftLeavesR1(root *TreeNode) (ret int) {
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

			if v.Left != nil && v.Left.Left == nil && v.Left.Right == nil {
				ret += v.Left.Val
			}
		}
	}

	return
}

func sumOfLeftLeaves2(root *TreeNode) (ret int) {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		ret += root.Left.Val
	}

	if root.Left != nil {
		ret += sumOfLeftLeaves(root.Left)
	}

	if root.Right != nil {
		ret += sumOfLeftLeaves(root.Right)
	}

	return
}
