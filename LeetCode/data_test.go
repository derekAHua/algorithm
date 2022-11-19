package LeetCode

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/16 23:49
// @Version 1.0

func Test_Traversal(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 8},
		},
	}

	t.Log(preTraversal(tree))
	t.Log(preorderTraversal(tree))
	t.Log(inTraversal(tree))
	t.Log(postTraversal(tree))
}
