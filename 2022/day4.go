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
	"regexp"
	"strconv"

	util "github.com/byorn/advent_of_code/util"
)

func Day4(inputFile string) int {
	inputText := util.ReadFile(inputFile)

	return getCountOfEntriesThatMatch(inputText)
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
