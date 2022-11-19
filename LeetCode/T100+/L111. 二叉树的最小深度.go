package T100_

import "container/list"

// https://leetcode.cn/problems/minimum-depth-of-binary-tree/

func minDepth(root *TreeNode) (ret int) {
	if root == nil {
		return
	}

	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		length := q.Len()
		for i := 0; i < length; i++ {
			node := q.Remove(q.Front()).(*TreeNode)

			if node.Left == nil && node.Right == nil {
				return ret + 1
			}

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
