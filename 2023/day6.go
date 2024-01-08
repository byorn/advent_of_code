package _2023

import (
	"fmt"

	util "github.com/byorn/advent_of_code/util"
)

func Day6(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)
	fmt.Printf("%d \n", len(inputText))
	return day6Part1(), day6Part2()
}

/*
Return the max total from all group totals
*/
func day6Part1() int {
	answ1 := calNumOfWaysToBeatDis(49, 356)

	answ2 := calNumOfWaysToBeatDis(87, 1378)

	answ3 := calNumOfWaysToBeatDis(78, 1502)

	answ4 := calNumOfWaysToBeatDis(95, 1882)
	return answ1 * answ2 * answ3 * answ4
}

func calNumOfWaysToBeatDis(time, distance int) int {
	counter := 0

	for i := 1; i < time; i++ {

		travelled := i * (time - i)

		if travelled > distance {
			counter++
		}

	}

	return counter
}

func day6Part2() int {
	return calNumOfWaysToBeatDis(49877895, 356137815021882)
}
