package T000_

// https://leetcode.cn/problems/swap-nodes-in-pairs/

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	t := new(ListNode)
	t.Next = head

	cur := t
	for cur.Next != nil && cur.Next.Next != nil {
		l := cur.Next
		r := cur.Next.Next

		l.Next = r.Next
		r.Next = l
		cur.Next = r

		cur = cur.Next.Next
	}

	return t.Next
}
