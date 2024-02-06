package _2023

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result, result2 := Day7(dir + "/input/day7.txt")
	assert.Equal(t, result, 246912307, "Is not 246912307")
	assert.Equal(t, result2, 246894760, "Is not 246894760")
}

func TestDay7CompareHands(t *testing.T) {
	assert.Equal(t, 7, GetHandType("KKKKK"), "Is not 7")
	assert.Equal(t, 6, GetHandType("KKKKQ"), "Is not 6")
	assert.Equal(t, 5, GetHandType("KKKQQ"), "Is not 5")
	assert.Equal(t, 4, GetHandType("KKKQJ"), "Is not 4")
	assert.Equal(t, 3, GetHandType("KKQQ2"), "Is not 3")
	assert.Equal(t, 2, GetHandType("KKQJ2"), "Is not 2")
	assert.Equal(t, 1, GetHandType("AKQJ2"), "Is not 1")
	assert.Equal(t, 1, CompareHands("KKKKK", "QQQQQ"), "KKKKK is not greater than QQQQQ")
	assert.Equal(t, -1, CompareHands("2JKKK", "2QKKK"), "is not -1")
}

func TestBSTLeaf(t *testing.T) {
	root := BSTLeaf{value: "KKKQJ", bid: "2323"}
	root.Insert("KKKKK", "23423")
	root.Insert("AKQJ2", "23423")
	root.Insert("KKKKQ", "23423")
	root.Insert("QQQQQ", "23423")

	fmt.Println(root.TraverseInOrder())
}
