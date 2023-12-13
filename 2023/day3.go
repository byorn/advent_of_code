package _2023

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	util "github.com/byorn/advent_of_code/util"
)

func Day3(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)
	return day3Part1(inputText), day3Part2(inputText)
}

/*
Return the max total from all group totals
*/
func day3Part1(inputLines []string) int {
	fmt.Println(inputLines[0])

	rows := [][]string{}
	rowsAndNumbers := [][]int{}

	for _, line := range inputLines {

		cols := []string{}

		for _, c := range line {
			cols = append(cols, string(c))
		}

		rows = append(rows, cols)
		numbersInLine := GetNumbersInLine(line)
		rowsAndNumbers = append(rowsAndNumbers, numbersInLine)
	}

	fmt.Print(rows)

	sum := 0

	for i, r := range rows {
		for j := 0; j < len(r)-1; {

			c := rows[i][j]
			if unicode.IsDigit([]rune(c)[0]) {

				isMachinePart := false
				numsInLine := rowsAndNumbers[i]
				numLen := len(strconv.Itoa(numsInLine[0]))

				// top horizontal
				if i > 0 {
					for k := 0; k < numLen-1; k++ {
						v := rows[i-1][j+k]
						if v != "." && unicode.IsSymbol([]rune(v)[0]) {
							isMachinePart = true
						}
					}

					if j > 0 {

						v := rows[i-1][j-1]
						if v != "." && unicode.IsSymbol([]rune(v)[0]) {
							isMachinePart = true
						}

					}

					if j+numLen < len(r) {
						v := rows[i-1][j+numLen]
						if v != "." && unicode.IsSymbol([]rune(v)[0]) {
							isMachinePart = true
						}

					}

				}

				// bottom horizontal
				if i < len(rows)-1 {
					for k := 0; k < numLen-1; k++ {
						v := rows[i+1][j+k]
						if v != "." && unicode.IsSymbol([]rune(v)[0]) {
							isMachinePart = true
						}
					}

					if j > 0 {

						v := rows[i+1][j-1]
						if v != "." && unicode.IsSymbol([]rune(v)[0]) {
							isMachinePart = true
						}

					}

					if j+numLen < len(r) {
						v := rows[i+1][j+numLen]
						if v != "." && unicode.IsSymbol([]rune(v)[0]) {
							isMachinePart = true
						}

					}

				}

				// behind
				if j > 0 {
					v := rows[i][j-1]
					if v != "." && unicode.IsSymbol([]rune(v)[0]) {
						isMachinePart = true
					}

				}

				// front
				if j+numLen < len(r) {
					v := rows[i][j+numLen]
					if v != "." && unicode.IsSymbol([]rune(v)[0]) {
						isMachinePart = true
					}

				}

				if isMachinePart {
					sum += numsInLine[0]
				}

				j = j + numLen
			} else {
				j += 1
			}

		}
	}

	return sum
}

func GetNumbersInLine(line string) []int {
	newLine := strings.Replace(line, ".", " ", -1)

	words := strings.Fields(newLine)

	var nums []int

	for _, w := range words {
		i, err := strconv.Atoi(w)
		if err == nil {
			nums = append(nums, i)
		}
	}
	return nums
}

func day3Part2(inputLines []string) int {
	return 0
}
