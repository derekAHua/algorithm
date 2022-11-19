package T100_

import "container/list"

// https://leetcode.cn/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (ret [][]int) {
	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(ret) == depth {
			ret = append(ret, []int{})
		}
		ret[depth] = append(ret[depth], root.Val)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, 0)

	return
}

func levelOrder3(root *TreeNode) (ret [][]int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		var r []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			v := queue.Remove(queue.Front()).(*TreeNode)
			r = append(r, v.Val)
			if v.Left != nil {
				queue.PushBack(v.Left)
			}
			if v.Right != nil {
				queue.PushBack(v.Right)
			}
		}
		ret = append(ret, r)
	}

	return
}

func levelOrder2(root *TreeNode) (ret [][]int) {
	if root == nil {
		return
	}

	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) != 0 {
		var r []int
		var ins []*TreeNode
		for _, v := range stack {
			if v == nil {
				continue
			}
			r = append(r, v.Val)
			if v.Left != nil {
				ins = append(ins, v.Left)
			}
			if v.Right != nil {
				ins = append(ins, v.Right)
			}
		}
		ret = append(ret, r)
		stack = ins
	}

	return
}
