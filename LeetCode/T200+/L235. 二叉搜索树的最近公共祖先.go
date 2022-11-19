package T200_

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func lowestCommonAncestorA(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestorA(root.Left, p, q)
	r := lowestCommonAncestorA(root.Right, p, q)
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
