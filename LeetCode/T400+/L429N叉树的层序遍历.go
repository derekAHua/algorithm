package T200_

import "container/list"

// https://leetcode.cn/problems/n-ary-tree-level-order-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (ret [][]int) {
	if root == nil {
		return
	}

	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		length := q.Len()
		var r []int
		for i := 0; i < length; i++ {
			node := q.Remove(q.Front()).(*Node)
			r = append(r, node.Val)
			for _, v := range node.Children {
				q.PushBack(v)
			}
		}
		ret = append(ret, r)
	}

	return
}
