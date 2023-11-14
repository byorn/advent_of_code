package util

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	myQueue := Queue{
		Items: []Item{},
	}
	myQueue.Push(3333)
	myQueue.Push("abc")

	fmt.Println(myQueue.Items)
	myQueue.Shift()

	fmt.Println(myQueue.Items)
}
