package T200_

// https://leetcode.cn/problems/reverse-linked-list/

//题意：反转一个单链表。
//
//示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL

func reverseList(head *ListNode) (ret *ListNode) {
	ret = new(ListNode)

	cur := head

	for cur != nil {
		node := cur
		cur = cur.Next

		node.Next = ret.Next
		ret.Next = node
	}

	return ret.Next
}

func reverseListR1(head *ListNode) *ListNode {
	t := new(ListNode)

	var cur *ListNode
	for head != nil {
		cur = head
		head = head.Next

		cur.Next = t.Next
		t.Next = cur
	}

	return t.Next
}
