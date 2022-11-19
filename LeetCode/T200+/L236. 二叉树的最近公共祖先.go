package T200_

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == q || root == p {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	if r != nil {
		return r
	}

	return nil
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}

	return nil
}
