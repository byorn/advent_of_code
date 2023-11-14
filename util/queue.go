package util

import (
	"errors"
)

type Queue struct {
	Items []Item
}

func (queue *Queue) Push(item Item) {
	Items := append(queue.Items, item)
	queue.Items = Items
}

func (queue *Queue) Shift() (Item, error) {
	if len(queue.Items) == 0 {
		return nil, errors.New("Queue is empty")
	}
	removed := queue.Items[0]
	Items := queue.Items[1:]
	queue.Items = Items
	return removed, nil
}
