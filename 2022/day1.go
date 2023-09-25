/*
For example, suppose the Elves finish writing their items' Calories and end up with the following list:

1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
This list represents the Calories of the food carried by five Elves:

The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
The second Elf is carrying one food item with 4000 Calories.
The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
The fifth Elf is carrying one food item with 10000 Calories.
In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the most Calories. In the example above, this is 24000
*/
package _2022

import (
	util "github.com/byorn/advent_of_code/util"
	"strconv"
)

func Day1(inputFile string) int {
	inputText := util.ReadFile(inputFile)
	return getGroupMax(inputText)
}

/*
Return the max total from all group totals
*/
func getGroupMax(inputLines []string) int {

	var maximumCalories int
	var currentTotalReindeerCalories int
	currentReindeer := 1

	for index, line := range inputLines {
		if line == "" {
			if currentTotalReindeerCalories > maximumCalories {
				maximumCalories = currentTotalReindeerCalories
			}
			currentTotalReindeerCalories = 0
			currentReindeer++
		} else {
			calories, _ := strconv.Atoi(line)
			currentTotalReindeerCalories += calories
			if index == len(inputLines)-1 {
				if currentTotalReindeerCalories > maximumCalories {
					maximumCalories = currentTotalReindeerCalories
				}
			}
		}
	}
	return maximumCalories
}
