package util

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	leaves []TreeNode
	parent *TreeNode
	size   int
	name   string
}

func InitRootTreeNode() *TreeNode {
	root := TreeNode{
		size: 0,
		name: "root",
	}
	return &root
}

func (treeNode *TreeNode) Add(name string, size int) {
	node := TreeNode{
		size:   size,
		name:   name,
		leaves: []TreeNode{},
		parent: treeNode,
	}

	treeNode.leaves = append(treeNode.leaves, node)
	treeNode.size += size

	parent := treeNode.parent
	for parent != nil {
		parent.size += size
		parent = parent.parent
	}
}

func (treeNode *TreeNode) CD(name string) (*TreeNode, error) {
	if name == ".." {
		return treeNode.parent, nil
	}

	for i, n := range treeNode.leaves {
		if n.name == name {
			return &treeNode.leaves[i], nil
			// return &n, nil
		}
	}
	return nil, errors.New("did not find :" + name)
}

func (treeNode *TreeNode) Print() {
	numberOfParents := 0
	parent := treeNode.parent
	for parent != nil {
		numberOfParents++
		parent = parent.parent
	}
	for i := 0; i < numberOfParents; i++ {
		fmt.Printf(" - ")
	}

	fmt.Printf("name: %s, size %d \n", treeNode.name, treeNode.size)
	for i := range treeNode.leaves {
		treeNode.leaves[i].Print()
	}
}

func isDirAndGetTotalIfLessThan100000(treeNode TreeNode) int {
	if len(treeNode.leaves) > 0 {
		if treeNode.size <= 100000 {
			return treeNode.size
		}
	}
	return 0
}

func (treeNode *TreeNode) FindTotalSumOfDirsHavingAtMost100000() int {
	sum := 0
	sum = isDirAndGetTotalIfLessThan100000(*treeNode)

	for i := range treeNode.leaves {
		sum += treeNode.leaves[i].FindTotalSumOfDirsHavingAtMost100000()
	}
	return sum
}

func (treeNode *TreeNode) FindLeastBiggestDirToDelete() int {
	available := 70000000 - treeNode.size
	required := 30000000 - available

	leastBiggest := 30000000
	currentNode := treeNode

	updateLeastBiggest(*currentNode, required, &leastBiggest)

	return leastBiggest
}

func updateLeastBiggest(treeNode TreeNode, required int, leastBiggest *int) {
	if len(treeNode.leaves) > 0 && treeNode.size > required && treeNode.size <= *leastBiggest {
		*leastBiggest = treeNode.size
	}

	for i := range treeNode.leaves {
		updateLeastBiggest(treeNode.leaves[i], required, leastBiggest)
	}
}
