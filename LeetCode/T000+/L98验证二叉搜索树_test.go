package T000_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/11/19 17:22
// @Version 1.0

func Test_isValidBST(t *testing.T) {

	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val:   11,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}

	t.Log(isValidBST(root))
}
