package T200_

// https://leetcode-cn.com/problems/remove-linked-list-elements/

/*
题意：删除链表中等于给定值 val 的所有节点。

示例 1：
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

示例 2：
输入：head = [], val = 1
输出：[]

示例 3：
输入：head = [7,7,7,7], val = 7
输出：[]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) (ret *ListNode) {
	cur := new(ListNode)
	cur.Next = head
	ret = cur
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}

	return ret.Next
}

func removeElementsR1(head *ListNode, val int) (ret *ListNode) {
	top := &ListNode{Next: head}

	cur := top
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			continue
		}

		cur = cur.Next
	}

	return top.Next
}
