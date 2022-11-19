package LeetCode

import "container/list"

// 确定递归函数的参数和返回值：
// 确定哪些参数是递归的过程中需要处理的，那么就在递归函数里加上这个参数， 并且还要明确每次递归的返回值是什么进而确定递归函数的返回类型。
//
// 确定终止条件：
// 写完了递归算法, 运行的时候，经常会遇到栈溢出的错误，就是没写终止条件或者终止条件写的不对，
// 操作系统也是用一个栈的结构来保存每一层递归的信息，如果递归没有终止，操作系统的内存栈必然就会溢出。
//
// 确定单层递归的逻辑：
// 确定每一层递归需要处理的信息。在这里也就会重复调用自己来实现递归的过程。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	res = append(res, root.Val)
	res = append(res, preTraversal(root.Left)...)
	res = append(res, preTraversal(root.Right)...)

	return
}

// 中序遍历
func inTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	res = append(res, inTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inTraversal(root.Right)...)

	return
}

// 后续遍历
func postTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	res = append(res, postTraversal(root.Left)...)
	res = append(res, postTraversal(root.Right)...)
	res = append(res, root.Val)

	return
}

// 迭代法前序遍历
func preorderTraversal(root *TreeNode) []int {
	var ans []int

	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}