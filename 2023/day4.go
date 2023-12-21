package _2023

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/byorn/advent_of_code/util"
)

func Day4(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)
	return day4Part1(inputText), day4Part2(inputText)
}

/*
Return the max total from all group totals
*/
func day4Part1(inputLines []string) int {
	sum := 0
	for _, str := range inputLines {

		unique := map[int]bool{}

		x, y := GetTwoArrays(str)

		for _, i := range x {
			unique[i] = false
		}

		count := 0
		for _, j := range y {
			_, exists := unique[j]
			if exists {
				if count == 0 {
					count += 1
				} else {
					count *= 2
				}
			}
		}

		sum += count

	}

	return sum
}

func day4Part2(inputLines []string) int {
	cardNumberCount := map[int]int{}
	cardNumberMatches := map[int]int{}

	queue := [][]int{{}}

	for cardNum, str := range inputLines {

		unique := map[int]bool{}

		x, y := GetTwoArrays(str)

		for _, i := range x {
			unique[i] = false
		}

		count := 0
		for _, j := range y {
			_, exists := unique[j]
			if exists {
				count += 1
			}
		}

		cardNumberCount[cardNum] += 1
		cardNumberMatches[cardNum] = count

		cardsToAddToQ := []int{}
		for x := 1; x <= count; x++ {
			cardsToAddToQ = append(cardsToAddToQ, cardNum+x)
		}

		queue = append(queue, cardsToAddToQ)
	}

	for len(queue) > 0 {

		shifted := queue[0]

		for _, cardNum := range shifted {
			cardNumberCount[cardNum] += 1

			appendToQueue := []int{}
			for x := 1; x <= cardNumberMatches[cardNum]; x++ {
				appendToQueue = append(appendToQueue, cardNum+x)
			}
			queue = append(queue, appendToQueue)
		}

		queue = queue[1:]
	}

	sum := 0
	for _, v := range cardNumberCount {
		sum += v
	}

	return sum
}

func GetTwoArrays(str string) ([]int, []int) {
	strWithouGame := str[10:]

	wAndh := strings.Split(strWithouGame, "|")

	result1 := strings.Fields(wAndh[0])
	result2 := strings.Fields(wAndh[1])

	r1 := []int{}
	for _, x := range result1 {

		num, err := strconv.Atoi(strings.Trim(x, " "))
		if err == nil {
			r1 = append(r1, num)
		} else {
			fmt.Println(x)
			panic("could not convert str to a number" + x)
		}
	}

	r2 := []int{}
	for _, x := range result2 {
		num, err := strconv.Atoi(x)
		if err == nil {
			r2 = append(r2, num)
		} else {
			panic("could not convert str to a number" + x)
		}
	}

	return r1, r2
}
