package T100_

// https://leetcode.cn/problems/linked-list-cycle-ii/

// 题意： 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

func detectCycle(head *ListNode) (ret *ListNode) {
	m := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}

		m[head] = struct{}{}
		head = head.Next
	}

	return
}

func detectCycle2(head *ListNode) *ListNode {
	l, r := head, head
	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next.Next

		if l == r {
			for l != head {
				l = l.Next
				head = head.Next
			}
			return head
		}
	}

	return nil
}
