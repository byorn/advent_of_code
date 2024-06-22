package util

import (
	"math"
)

type Point struct {
	Row   int
	Col   int
	Steps int
}

func CalculateDistance(array [][]string, from, to string) int {
	fromPoint := FindPoint(array, from)
	toPoint := FindPoint(array, to)

	dis1 := fromPoint.Row - toPoint.Row
	dis2 := fromPoint.Col - toPoint.Col
	return int(math.Abs(float64(dis1)) + math.Abs(float64(dis2)))
}

func ShortestPath(array [][]string, from string, to string) int {
	fromPoint := FindPoint(array, from)
	toPoint := FindPoint(array, to)

	var queue []Point
	queue = append(queue, *fromPoint)
	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	rows, cols := len(array), len(array[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	for len(queue) > 0 {
		currentPoint := queue[0]
		queue = queue[1:]

		if currentPoint.Col == toPoint.Col &&
			currentPoint.Row == toPoint.Row {
			return currentPoint.Steps
		}

		for _, d := range directions {
			nextCol := currentPoint.Col + d[1]
			nextRow := currentPoint.Row + d[0]
			if nextCol >= 0 && nextCol < len(array[0]) && nextRow >= 0 && nextRow < len(array) && !visited[nextRow][nextCol] {
				p := &Point{
					Row:   nextRow,
					Col:   nextCol,
					Steps: currentPoint.Steps + 1,
				}
				visited[nextRow][nextCol] = true
				queue = append(queue, *p)
			}
		}
	}

	return -9999
}

func FindPoint(array [][]string, target string) *Point {
	for ri, r := range array {
		for ci, c := range r {
			if c == target {
				return &Point{
					Row: ri,
					Col: ci,
				}
			}
		}
	}
	return nil
}
