package _2023

import (
	"fmt"

	util "github.com/byorn/advent_of_code/util"
)

func Day11(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)
	fmt.Printf("%d \n", len(inputText))
	return day11Part1(inputText), day11Part2(inputText)
}

func day11Part1(input []string) int {
	var universe [][]string
	for _, line := range input {

		var universeRow []string
		for _, lineVal := range line {
			universeRow = append(universeRow, string(lineVal))
		}
		universe = append(universe, universeRow)
	}
	expandUniverse(universe)
	return 0
}

func expandUniverse(universe [][]string) {
	universe = expandRows(universe)
	universe = expandCols(universe)

	printUniverse(universe)
}

func printUniverse(universe [][]string) {
	for _, r := range universe {
		fmt.Println(r)
	}
}

func expandCols(universe [][]string) [][]string {
	var colIndexesToExpand []int
	for colIndex := range universe[0] {

		colHasHash := false
		for rowIndex := range universe {
			if universe[rowIndex][colIndex] == "#" {
				colHasHash = true
			}
		}
		if !colHasHash {
			colIndexesToExpand = append(colIndexesToExpand, colIndex)
		}

	}

	for i, row := range universe {
		afterInserting := 0
		for _, colIndex := range colIndexesToExpand {
			row = util.InsertStringToIndex(row, colIndex+afterInserting, row[colIndex+afterInserting])
			universe[i] = row
			afterInserting++
		}
	}
	return universe
}

func expandRows(universe [][]string) [][]string {
	rowsTocopy := []int{}

	for index, r := range universe {
		rowHasHash := false
		for _, c := range r {
			if c == "#" {
				rowHasHash = true
			} else {
			}
		}
		if !rowHasHash {
			rowsTocopy = append(rowsTocopy, index)
		}
	}

	afterInserting := 0
	for _, rowIndexToCopy := range rowsTocopy {
		rowToCopy := universe[rowIndexToCopy+afterInserting]
		universe = util.InsertStringArrayToIndex(universe, rowIndexToCopy+afterInserting, rowToCopy)
		afterInserting++
	}
	return universe
}

func day11Part2(input []string) int {
	//	fmt.Println(input)
	return 0
}
