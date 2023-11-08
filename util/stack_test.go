package util

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	myStack := NewStack([]Item{1, 33, 213})
	myStack.Print()
	myStack.Push(3333)
	myStack.Print()
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Peek())
	myStack.Push("abc")

	fmt.Println(myStack.Peek())
}
