package _2023

import (
	"regexp"
	"strconv"
	"strings"

	util "github.com/byorn/advent_of_code/util"
)

func Day1(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)
	return getTotal(inputText), GetTotalPart2(inputText)
}

/*
Return the max total from all group totals
*/
func getTotal(inputLines []string) int {
	total := 0
	for _, line := range inputLines {
		total += getNumber(line)
	}
	return total
}

func GetTotalPart2(inputLines []string) int {
	total := 0
	for _, line := range inputLines {
		replacedWordWithNum := replaceWordNumberToDigit(line)

		total += getNumber(replacedWordWithNum)
	}
	return total
}

func replaceWordNumberToDigit(line string) string {
	charsFrontToBack := ""
	for _, char := range line {

		charsFrontToBack += string(char)
		s1 := strings.Replace(charsFrontToBack, "one", "1", -1)
		s2 := strings.Replace(s1, "two", "2", -1)
		s3 := strings.Replace(s2, "three", "3", -1)
		s4 := strings.Replace(s3, "four", "4", -1)
		s5 := strings.Replace(s4, "five", "5", -1)
		s6 := strings.Replace(s5, "six", "6", -1)
		s7 := strings.Replace(s6, "seven", "7", -1)
		s8 := strings.Replace(s7, "eight", "8", -1)
		s9 := strings.Replace(s8, "nine", "9", -1)

		charsFrontToBack = s9

	}

	charsBackToFront := ""
	lineRuneArr := []rune(line)
	for i := len(lineRuneArr) - 1; i >= 0; i-- {
		charsBackToFront = string(lineRuneArr[i]) + charsBackToFront
		s1 := strings.Replace(charsBackToFront, "one", "1", -1)
		s2 := strings.Replace(s1, "two", "2", -1)
		s3 := strings.Replace(s2, "three", "3", -1)
		s4 := strings.Replace(s3, "four", "4", -1)
		s5 := strings.Replace(s4, "five", "5", -1)
		s6 := strings.Replace(s5, "six", "6", -1)
		s7 := strings.Replace(s6, "seven", "7", -1)
		s8 := strings.Replace(s7, "eight", "8", -1)
		s9 := strings.Replace(s8, "nine", "9", -1)

		charsBackToFront = s9

	}

	return charsFrontToBack + charsBackToFront
}

func getNumber(line string) int {
	// Create a regular expression pattern to match non-numeric characters
	re := regexp.MustCompile("[^0-9]")

	// Use the regular expression to replace non-numeric characters with an empty string
	result := re.ReplaceAllString(line, "")

	num, error := strconv.Atoi(string(result[0]) +
		string(result[len(result)-1]))

	if error == nil {
		return num
	}

	return num
}
