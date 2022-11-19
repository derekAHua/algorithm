package T100_

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

func buildTree(preorder []int, inorder []int) (ret *TreeNode) {
	if len(inorder) < 1 || len(preorder) < 1 {
		return nil
	}

	nodeValue := preorder[0]
	left := getIndex(inorder, nodeValue)
	ret = &TreeNode{
		Val:   nodeValue,
		Left:  buildTree(preorder[1:left+1], inorder[:left]),
		Right: buildTree(preorder[left+1:], inorder[left+1:]),
	}

	return
}
