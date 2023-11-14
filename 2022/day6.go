package _2022

import (
	util "github.com/byorn/advent_of_code/util"
)

func Day6(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)

	return part1Day6(inputText), part2Day6(inputText)
}

func part1Day6(entries []string) int {
	input := entries[0]

	q := util.Queue{
		Items: []util.Item{},
	}

	for index, inputChar := range input {
		if len(q.Items) == 3 && !util.Contains(q.Items, string(inputChar)) {
			return index + 1
		} else if len(q.Items) <= 3 {
			if util.Contains(q.Items, string(inputChar)) {
				for _, qi := range q.Items {
					q.Shift()

					if qi == string(inputChar) {
						break
					}
				}
			}
			q.Push(string(inputChar))
		}
		//
	}
	return 0
}

func part2Day6(entries []string) int {
	input := entries[0]

	q := util.Queue{
		Items: []util.Item{},
	}

	for index, inputChar := range input {
		if len(q.Items) == 13 && !util.Contains(q.Items, string(inputChar)) {
			return index + 1
		} else if len(q.Items) <= 13 {
			if util.Contains(q.Items, string(inputChar)) {
				for _, qi := range q.Items {
					q.Shift()

					if qi == string(inputChar) {
						break
					}
				}
			}
			q.Push(string(inputChar))
		}
		//
	}
	return 0
}
