package util

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	i := 1

	ll := LinkedList{}
	ll.insert(i)
	ll.insert(2)
	ll.insert(3)

	ll.print()
}
