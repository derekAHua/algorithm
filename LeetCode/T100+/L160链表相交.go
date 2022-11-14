package T100_

// https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/

//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

func getIntersectionNode(headA, headB *ListNode) *ListNode {
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
