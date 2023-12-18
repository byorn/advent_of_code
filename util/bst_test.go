package util

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	unsortedArr := []int{25, 20, 15, 27, 30, 29, 26, 22, 32}
	root := BSTNode{}

	for _, i := range unsortedArr {
		Insert(&root, i)
	}

	fmt.Println(root)

	fmt.Println("In Order : ")
	fmt.Printf("\n")
	TraverseInOrder(&root)

	fmt.Printf("\n")
	fmt.Println("Pre Order : ")
	TraversePreOrder(&root)

	fmt.Printf("\n")
	fmt.Println("Post Order : ")
	TraversePostOrder(&root)
}
