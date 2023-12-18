package util

import (
	"fmt"
)

type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

func Insert(node *BSTNode, val int) {
	if node.value == 0 {
		node.value = val
	} else if val <= node.value {
		if node.left != nil {
			Insert(node.left, val)
		} else {
			node.left = &BSTNode{}
			Insert(node.left, val)
		}
	} else if val >= node.value {
		if node.right != nil {
			Insert(node.right, val)
		} else {
			node.right = &BSTNode{}
			Insert(node.right, val)
		}
	}
}

func TraverseInOrder(node *BSTNode) {
	if node.left != nil {
		TraverseInOrder(node.left)
	}

	fmt.Printf("%d ", node.value)

	if node.right != nil {
		TraverseInOrder(node.right)
	}
}

func TraversePreOrder(node *BSTNode) {
	fmt.Printf("%d ", node.value)
	if node.left != nil {
		TraversePreOrder(node.left)
	}

	if node.right != nil {
		TraversePreOrder(node.right)
	}
}

func TraversePostOrder(node *BSTNode) {
	if node.left != nil {
		TraversePostOrder(node.left)
	}

	if node.right != nil {
		TraversePostOrder(node.right)
	}

	fmt.Printf("%d ", node.value)
}
