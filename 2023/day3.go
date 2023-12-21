package _2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	util "github.com/byorn/advent_of_code/util"
)

func Day3(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)
	return day3Part1(inputText), day3Part2(inputText)
}

func isSymbol(str string) bool {
	if str == "!" || str == "@" || str == "#" || str == "$" || str == "%" || str == "^" || str == "&" || str == "*" || str == "(" || str == ")" || str == "-" || str == "+" || str == "=" || str == "/" {
		return true
	}
	return false
}

func isDigit(str string) bool {
	return unicode.IsDigit([]rune(str)[0])
}

/*
Return the max total from all group totals
*/
func day3Part1(inputLines []string) int {
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
					for k := 0; k < numLen; k++ {
						v := rows[i-1][j+k]
						if numsInLine[0] == 434 {
							fmt.Printf("%s", v)
						}
						if v != "." && isSymbol(v) {
							isMachinePart = true
						}
					}

					if j > 0 {

						v := rows[i-1][j-1]
						if v != "." && isSymbol(v) {
							isMachinePart = true
						}

					}

					if j+numLen < len(r) {
						v := rows[i-1][j+numLen]

						if v != "." && isSymbol(v) {
							isMachinePart = true
						}

					}

				}

				// bottom horizontal
				if i < len(rows)-1 {
					for k := 0; k < numLen; k++ {
						v := rows[i+1][j+k]
						if v != "." && isSymbol(v) {
							isMachinePart = true
						}
					}

					if j > 0 {

						v := rows[i+1][j-1]
						if v != "." && isSymbol(v) {
							isMachinePart = true
						}

					}

					if j+numLen < len(r) {
						v := rows[i+1][j+numLen]
						if v != "." && isSymbol(v) {
							isMachinePart = true
						}

					}

				}

				// behind
				if j > 0 {
					v := rows[i][j-1]
					if v != "." && isSymbol(v) {
						isMachinePart = true
					}

				}

				// front
				if j+numLen < len(r) {
					v := rows[i][j+numLen]
					if v != "." && isSymbol(v) {
						isMachinePart = true
					}

				}

				if isMachinePart {
					sum += numsInLine[0]
				}

				j = j + numLen
				rowsAndNumbers[i] = rowsAndNumbers[i][1:]
			} else {
				j += 1
			}

		}
	}

	return sum
}

func GetNumbersInLine(line string) []int {
	newLine := strings.Replace(line, ".", " ", -1)
	re := regexp.MustCompile(`[^0-9]`)

	// Replace all non-alphabetic characters with a space
	result := re.ReplaceAllString(newLine, " ")
	words := strings.Fields(result)

	var nums []int

	for _, w := range words {
		i, err := strconv.Atoi(w)
		if err == nil {
			nums = append(nums, i)
		}
	}
	return nums
}

func SearchNumber(matrix [][]string, x, y int) int {
	originalString := matrix[x][y]
	prefix := ""
	suffix := ""

	for i := y - 1; i >= 0; i-- {
		if isDigit(matrix[x][i]) {
			prefix = matrix[x][i] + prefix
		} else {
			break
		}
	}

	for i := y + 1; i < len(matrix[x]); i++ {
		if isDigit(matrix[x][i]) {
			suffix += matrix[x][i]
		} else {
			break
		}
	}

	resultString := fmt.Sprintf("%s%s%s", prefix, originalString, suffix)
	num, err := strconv.Atoi(resultString)
	if err == nil {
		return num
	} else {
		panic("the number i got wasnt a number")
	}
}

func day3Part2(inputLines []string) int {
	rows := [][]string{}

	for _, line := range inputLines {

		cols := []string{}

		for _, c := range line {
			cols = append(cols, string(c))
		}

		rows = append(rows, cols)
	}

	sum := 0

	for i, r := range rows {
		for j := 0; j < len(r)-1; {

			c := rows[i][j]
			if isSymbol(c) {

				foundDigit := []int{}

				// top horizontal
				if i > 0 {
					v := rows[i-1][j]
					top_hor_found := false
					if v != "." && isDigit(v) {
						foundDigit = append(foundDigit, SearchNumber(rows, i-1, j))
						top_hor_found = true
					}

					if !top_hor_found {
						if j > 0 {

							v := rows[i-1][j-1]
							if v != "." && isDigit(v) {
								foundDigit = append(foundDigit, SearchNumber(rows, i-1, j-1))
							}

						}

						if j+1 < len(r) {
							v := rows[i-1][j+1]

							if v != "." && isDigit(v) {
								foundDigit = append(foundDigit, SearchNumber(rows, i-1, j+1))
							}

						}
					}

				}

				// bottom horizontal
				if i < len(rows)-1 {
					v := rows[i+1][j]
					bottom_hor_found := false
					if v != "." && isDigit(v) {
						foundDigit = append(foundDigit, SearchNumber(rows, i+1, j))
						bottom_hor_found = true
					}

					if !bottom_hor_found {
						if j > 0 {

							v := rows[i+1][j-1]
							if v != "." && isDigit(v) {
								foundDigit = append(foundDigit, SearchNumber(rows, i+1, j-1))
							}

						}

						if j+1 < len(r) {
							v := rows[i+1][j+1]
							if v != "." && isDigit(v) {
								foundDigit = append(foundDigit, SearchNumber(rows, i+1, j+1))
							}

						}
					}

				}

				// behind
				if j > 0 {
					v := rows[i][j-1]
					if v != "." && isDigit(v) {
						foundDigit = append(foundDigit, SearchNumber(rows, i, j-1))
					}

				}

				// front
				if j+1 < len(r) {
					v := rows[i][j+1]
					if v != "." && isDigit(v) {
						foundDigit = append(foundDigit, SearchNumber(rows, i, j+1))
					}

				}

				if len(foundDigit) == 2 {
					// fmt.Printf("num1: %d num2 : %d", foundDigit[0], foundDigit[1])
					sum += (foundDigit[0] * foundDigit[1])
				}

				j = j + 1
			} else {
				j += 1
			}

		}
	}

	return sum
}
