/*
For example, suppose you have the following list of contents from six rucksacks:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp. The only item type that appears in both compartments is lowercase p.
The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
The fourth rucksack's compartments only share item type v.
The fifth rucksack's compartments only share item type t.
The sixth rucksack's compartments only share item type s.
To help prioritize item rearrangement, every item type can be converted to a priority:

Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52.
In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.
*/

package _2022

import (
	util "github.com/byorn/advent_of_code/util"
	"strings"
)

func Day3_Part1(inputFile string) int {
	inputText := util.ReadFile(inputFile)
	return getSumOfCharsCommonToBothGroups(inputText)
}

func Day3_Part2(inputFile string) int {
	inputText := util.ReadFile(inputFile)
	return getSumOfCharsCommonToThreeGroups(inputText)
}

func getSumOfCharsCommonToBothGroups(charLines []string) int {
	var commonChars []string
	for _, charLine := range charLines {

		halfIndex := len(charLine) / 2
		group1 := charLine[0:halfIndex]
		group2 := charLine[halfIndex:]

		commonInLine := findCommon(group1, group2)
		commonChars = append(commonChars, commonInLine[:]...)

	}
	return getTotal(commonChars)
}

func getSumOfCharsCommonToThreeGroups(charLines []string) int {

	var commonChars []string
	counter := 1
	group1 := ""
	group2 := ""
	group3 := ""
	for _, charLine := range charLines {

		switch counter {
		case 1:
			{
				group1 = charLine
				break
			}
		case 2:
			{
				group2 = charLine
				break
			}
		case 3:
			{
				group3 = charLine
				break
			}
		}

		if counter == 3 {
			commonInLine := findCommon(group1, group2, group3)
			commonChars = append(commonChars, commonInLine[:]...)
			counter = 1
		} else {
			counter++
		}
	}

	return getTotal(commonChars)
}

func findCommon(group1 string, groups ...string) []string {
	var uniqueCommon = make(map[string]string)

	for _, g1 := range group1 {
		notCommon := false
		for _, group := range groups {
			if !strings.Contains(group, string(g1)) {
				notCommon = true
			}
		}
		if !notCommon {
			uniqueCommon[string(g1)] = string(g1)
		}
	}
	var stringCommon []string
	for key := range uniqueCommon {
		stringCommon = append(stringCommon, key)
	}
	return stringCommon
}

func getTotal(commonChars []string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	points := 0
	for _, cc := range commonChars {

		if strings.ToLower(cc) == cc {
			points += strings.Index(alphabet, cc) + 1
		}
		if strings.ToUpper(cc) == cc {
			points += strings.Index(alphabet, strings.ToLower(cc)) + 1 + 26
		}
	}
	return points
}
