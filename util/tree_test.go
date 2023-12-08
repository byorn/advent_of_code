package util

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := TreeNode{
		name: "root",
		size: 0,
	}

	root.Add("testfile", 131232)
	root.Add("Projects", 0)
	d, err := root.CD("Projects")
	if err == nil {
		d.Add("cv.doc", 122)
		d.Add("KUBERNETES", 0)
	}
	k, err := d.CD("KUBERNETES")
	if err == nil {
		k.Add("deployment.yaml", 44)
	}
	root.Add("expenses.xlsx", 300)
	root.Print()

	fmt.Println("--------")

	fmt.Println(root)
}
