package T600_

import "container/list"

// https://leetcode.cn/problems/average-of-levels-in-binary-tree/

// 给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) (ret []float64) {
	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) != 0 {
		var arr []*TreeNode
		r := make([]int, 0, len(stack))
		for i := 0; i < len(stack); i++ {
			if stack[i] == nil {
				continue
			}
			r = append(r, stack[i].Val)
			arr = append(arr, stack[i].Left, stack[i].Right)
		}
		if len(r) > 0 {
			sum := 0
			for _, v := range r {
				sum += v
			}
			ret = append(ret, float64(sum)/float64(len(r)))
		}
		stack = arr
	}

	return
}

func averageOfLevelsR1(root *TreeNode) (ret []float64) {
	if root == nil {
		return
	}

	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		length := q.Len()
		sum := 0
		for i := 0; i < length; i++ {
			node := q.Remove(q.Front()).(*TreeNode)
			sum += node.Val
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		ret = append(ret, float64(sum)/float64(length))
	}

	return
}
