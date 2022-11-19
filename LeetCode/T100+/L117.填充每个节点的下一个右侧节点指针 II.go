package T100_

import "container/list"

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/

func connect2(root *Node) (ret *Node) {
	if root == nil {
		return root
	}

	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		length := q.Len()
		var l []*Node
		for i := 0; i < length; i++ {
			node := q.Remove(q.Front()).(*Node)
			l = append(l, node)
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}

		for i := 0; i < len(l)-1; i++ {
			l[i].Next = l[i+1]
		}
	}

	return root
}
