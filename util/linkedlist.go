package util

import (
	"fmt"
)

type LinkedList struct {
	head *LinkedListNode
}

type LinkedListNode struct {
	next  *LinkedListNode
	value int
}

func (headNode *LinkedListNode) insert(val int) *LinkedListNode {
	return &LinkedListNode{next: headNode, value: val}
}

func (ll *LinkedList) insert(val int) {
	if ll.head == nil {
		ll.head = &LinkedListNode{value: val, next: nil}
	} else {
		ll.head = ll.head.insert(val)
	}
}

func (ll *LinkedList) print() {
	for lln := ll.head; lln != nil; lln = lln.next {
		fmt.Printf("<-  %d ", lln.value)
	}
}
