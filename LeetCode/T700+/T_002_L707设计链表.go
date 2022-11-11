package T700_

import "fmt"

// https://leetcode-cn.com/problems/design-linked-list/

/*
在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
*/

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head   *Node
	Length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: new(Node),
	}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.Length {
		return -1
	}

	cur := l.Head
	for index >= 0 {
		cur = cur.Next
		index--
	}

	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	ins := &Node{
		Val:  val,
		Next: nil,
	}

	next := l.Head.Next
	l.Head.Next = ins
	ins.Next = next

	l.Length++
}

func (l *MyLinkedList) AddAtTail(val int) {
	ins := &Node{
		Val:  val,
		Next: nil,
	}

	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = ins

	l.Length++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		l.AddAtHead(val)
		return
	}
	if index > l.Length {
		return
	}

	ins := &Node{
		Val:  val,
		Next: nil,
	}

	cur := l.Head
	for index > 0 {
		cur = cur.Next
		index--
	}

	next := cur.Next
	cur.Next = ins
	ins.Next = next

	l.Length++

	return
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.Length {
		return
	}

	cur := l.Head
	for index > 0 {
		cur = cur.Next
		index--
	}

	next := cur.Next
	cur.Next = next.Next

	l.Length--
}

func (l MyLinkedList) print() {
	head := l.Head.Next
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}
