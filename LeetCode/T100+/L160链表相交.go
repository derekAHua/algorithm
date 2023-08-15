package T100_

// https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/

//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

func getIntersectionNode(headA, headB *ListNode) (ret *ListNode) {

	h1, h2 := headA, headB
	for h1 != nil {

		for h2 != nil {
			if h1 == h2 {
				return h1
			}
			h2 = h2.Next
		}

		h2 = headB
		h1 = h1.Next
	}

	return
}

func getIntersectionNodeR1(headA, headB *ListNode) *ListNode {
	ha := headA
	hb := headB

	for hb != nil {
		for ha != nil {
			if ha == hb {
				return ha
			}
			ha = ha.Next
		}
		hb = hb.Next
		ha = headA
	}

	return nil
}
