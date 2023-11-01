package _2022

import (
	"regexp"
	"strconv"

	util "github.com/byorn/advent_of_code/util"
)

func Day4(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)

	part1 := getCountOfEntriesThatMatch(inputText)
	part2 := getCountOfEntriesThatMatchPart2(inputText)
	return part1, part2
}

func getCountOfEntriesThatMatch(entries []string) int {
	count := 0
	for _, entry := range entries {

		nums := extractNumbersFromString(entry)

		if isWithinRange(nums[0], nums[1], nums[2], nums[3]) ||
			isWithinRange(nums[2], nums[3], nums[0], nums[1]) {
			count++
		}
	}
	return count
}

func getCountOfEntriesThatMatchPart2(entries []string) int {
	count := 0
	for _, entry := range entries {

		nums := extractNumbersFromString(entry)

		if isWithinAtleastOneRange(nums[0], nums[1], nums[2], nums[3]) {
			count++
		}
	}
	return count
}

func extractNumbersFromString(text string) []int {
	re := regexp.MustCompile(`\d+`)

	strNums := re.FindAllString(text, -1)

	intNums := make([]int, len(strNums))

	for i, str := range strNums {
		intNums[i], _ = strconv.Atoi(str)
	}
	return intNums
}

func isWithinRange(a1, a2, b1, b2 int) bool {
	if a1 >= b1 && a2 <= b2 {
		return true
	}
	return false
}

func isWithinAtleastOneRange(a1, a2, b1, b2 int) bool {
	if (a1 <= b1 && a2 >= b1) ||
		(a1 >= b1 && a2 <= b2) ||
		(a1 <= b2 && a2 >= b2) {
		return true
	}
	return false
}
