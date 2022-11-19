package T500_

import "container/list"

// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}

	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		length := q.Len()
		max := q.Front().Value.(*TreeNode).Val
		for i := 0; i < length; i++ {
			node := q.Remove(q.Front()).(*TreeNode)
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		ret = append(ret, max)
	}

	return
}
