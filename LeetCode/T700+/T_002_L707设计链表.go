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

type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

// Constructor 仅保存哑节点，pre-> rear, next-> head
func Constructor() MyLinkedList {
	rear := &Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	rear.Next = rear
	rear.Pre = rear
	return MyLinkedList{rear}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.dummy.Next
	for cur != this.dummy && index > 0 {
		cur = cur.Next
		index--
	}

	if 0 != index {
		return -1
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Next: dummy.Next,
		Pre:  dummy,
	}

	dummy.Next.Pre = node
	dummy.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Next: dummy,
		Pre:  dummy.Pre,
	}

	dummy.Pre.Next = node
	dummy.Pre = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this.dummy.Next
	for cur != this.dummy && index > 0 {
		cur = cur.Next
		index--
	}

	if index <= 0 {
		node := &Node{
			Val:  val,
			Next: cur,
			Pre:  cur.Pre,
		}

		cur.Pre.Next = node
		cur.Pre = node
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	//链表为空
	if this.dummy.Next == this.dummy {
		return
	}

	cur := this.dummy.Next
	for cur.Next != this.dummy && index > 0 {
		cur = cur.Next
		index--
	}

	//验证index有效
	if index == 0 {
		cur.Next.Pre = cur.Pre
		cur.Pre.Next = cur.Next
		//以上两步顺序无所谓
	}
}

func (this MyLinkedList) print() {
	head := this.dummy.Next
	for head != this.dummy {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
