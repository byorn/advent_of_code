package util

import (
	"errors"
	"fmt"
)

type Item interface{}

type Stack struct {
	items []Item
}

func NewStack(items []Item) Stack {
	return Stack{
		items: items,
	}
}

func (stack *Stack) Print() {
	fmt.Println(stack.items)
}

func (stack *Stack) Push(item Item) Item {
	items := append(stack.items, item)

	stack.items = items

	return items
}

func (stack *Stack) Pop() (Item, error) {
	if len(stack.items) == 0 {
		return nil, errors.New("The stack is empty")
	}

	popped := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]

	return popped, nil
}

func (stack *Stack) Peek() (Item, error) {
	if len(stack.items) == 0 {
		return nil, errors.New("The stack is empty")
	}
	return stack.items[len(stack.items)-1], nil
}
