package T100_

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

func _buildTree(inorder []int, postorder []int) (ret *TreeNode) {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}

	nodeValue := postorder[len(postorder)-1]
	left := getIndex(inorder, nodeValue)
	ret = &TreeNode{
		Val:   nodeValue,
		Left:  _buildTree(inorder[:left], postorder[:left]),
		Right: _buildTree(inorder[left+1:], postorder[left:len(postorder)-1]),
	}

	return
}

func getIndex(arr []int, tar int) int {
	for i, v := range arr {
		if v == tar {
			return i
		}
	}

	return -1
}
