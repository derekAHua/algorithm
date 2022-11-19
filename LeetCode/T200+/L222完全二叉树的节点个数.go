package T200_

import "container/list"

// https://leetcode.cn/problems/count-complete-tree-nodes/

func countNodes(root *TreeNode) (ret int) {
	if root == nil {
		return
	}

	ret = countNodes(root.Left) + countNodes(root.Right) + 1

	return
}

func countNodes2(root *TreeNode) (ret int) {
	if root == nil {
		return
	}

	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		length := q.Len()
		for i := 0; i < length; i++ {
			node := q.Remove(q.Front()).(*TreeNode)
			ret++
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
	}

	return
}
