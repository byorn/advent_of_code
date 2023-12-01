package _2022

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/byorn/advent_of_code/util"
)

func Day7(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)

	tree := buildTree(inputText)
	return part1Day7(*tree), part2Day7(*tree) // part1Day7(inputText)
}

func buildTree(entries []string) *util.TreeNode {
	root := util.InitRootTreeNode()

	activeNode := root

	for i, line := range entries {

		if i == 0 {
			continue
		}

		// is a command
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd") {
				folderName := strings.Replace(line, "$ cd", "", -1)
				cdFolder, error := activeNode.CD(strings.Trim(folderName, " "))

				if error != nil {
					fmt.Println(error)
				} else {
					activeNode = cdFolder
				}
			}
		} else {
			// is a result
			if strings.HasPrefix(line, "dir") {

				dirName := strings.Replace(line, "dir", "", -1)
				activeNode.Add(strings.Trim(dirName, " "), 0)
			} else {
				sizeAndFileName := strings.Split(line, " ")
				sizeInt, _ := strconv.Atoi(sizeAndFileName[0])
				activeNode.Add(sizeAndFileName[1], sizeInt)
			}
		}
	}
	return root
}

func part1Day7(treeNode util.TreeNode) int {
	return treeNode.FindTotalSumOfDirsHavingAtMost100000()
}

func part2Day7(tree util.TreeNode) int {
	return tree.FindLeastBiggestDirToDelete()
}
