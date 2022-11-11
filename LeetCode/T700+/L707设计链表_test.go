package T700_

import (
	"fmt"
	"testing"
)

func TestMyLinkedList1(t *testing.T) {
	l := Constructor()
	l.AddAtTail(1)
	l.Get(1)
}

func TestMyLinkedList(t *testing.T) {
	l := Constructor()
	l.AddAtHead(7)
	l.print()
	l.AddAtHead(2)
	l.print()
	l.AddAtHead(1)
	l.print()
	l.AddAtIndex(3, 0)
	l.print()
	l.DeleteAtIndex(2)
	l.print()
	l.AddAtHead(6)
	l.print()
	l.AddAtTail(4)
	l.print()
	fmt.Println(l.Get(4))
}
