package T000_

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

func removeNthFromEnd(head *ListNode, n int) (ret *ListNode) {
	ret = new(ListNode)
	ret.Next = head

	cur := ret
	l := getLen(head) - n

	for i := l; i > 0 && cur != nil; i-- {
		cur = cur.Next
	}

	if cur != nil && cur.Next != nil {
		cur.Next = cur.Next.Next
	}

	return ret.Next
}

func removeNthFromEndR1(head *ListNode, n int) *ListNode {
	l := getLen(head)
	index := l - n
	if index < 0 {
		return head
	}

	t := new(ListNode)
	t.Next = head

	cur := t
	for index > 0 {
		cur = cur.Next
		index--
	}
	cur.Next = cur.Next.Next

	return t.Next
}

func getLen(head *ListNode) (ret int) {
	for head != nil {
		ret++
		head = head.Next
	}

	return
}
