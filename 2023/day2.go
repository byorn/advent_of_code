package _2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	util "github.com/byorn/advent_of_code/util"
)

func Day2(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)
	return day2Part1(inputText), day2Part2(inputText)
}

/*
Return the max total from all group totals
*/
func day2Part1(inputLines []string) int {
	total := 0
	max_red := 12
	max_green := 13
	max_blue := 14
	for _, line := range inputLines {

		values := strings.Split(line, ";")
		gameNum := extractGameNumber(values[0])
		notPossible := false
		for _, v := range values {
			if max_blue < extract("blue", v) ||
				max_red < extract("red", v) ||
				max_green < extract("green", v) {
				notPossible = true
			}
		}

		if !notPossible {
			total += gameNum
		}

	}
	return total
}

func day2Part2(inputLines []string) int {
	total := 0
	for _, line := range inputLines {
		max_red := 0
		max_green := 0
		max_blue := 0

		values := strings.Split(line, ";")
		for _, v := range values {
			if max_blue < extract("blue", v) {
				max_blue = extract("blue", v)
			}

			if max_red < extract("red", v) {
				max_red = extract("red", v)
			}

			if max_green < extract("green", v) {
				max_green = extract("green", v)
			}
		}
		power := max_red * max_green * max_blue
		total += power

	}
	return total
}

func extractGameNumber(line string) int {
	pattern := regexp.MustCompile(`Game (\d+)`)
	matches := pattern.FindStringSubmatch(line)

	if len(matches) >= 2 {
		val, _ := strconv.Atoi(matches[1])

		return val
	}
	return 0
}

func extract(color string, line string) int {
	patternString := fmt.Sprintf(`(\d+) %s`, color)
	pattern := regexp.MustCompile(patternString)

	// Find all matches in the input line
	matches := pattern.FindStringSubmatch(line)

	if len(matches) >= 2 {
		val, _ := strconv.Atoi(matches[1])

		return val
	}
	return 0
}
