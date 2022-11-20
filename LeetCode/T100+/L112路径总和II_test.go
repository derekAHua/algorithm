package T100_

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/19 15:18
// @Version 1.0

func Test_pathSum(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	sum2 := pathSum(tree, 22)
	t.Log(sum2)
}
