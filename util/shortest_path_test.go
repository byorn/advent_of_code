package util

import (
	"fmt"
	"testing"
)

func TestShortestPath(t *testing.T) {
	grid := [][]string{
		{".", ".", ".", ".", "1", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "2", ".", ".", "."},
		{"3", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "4", ".", ".", ".", ".", "."},
		{".", "5", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "7", ".", ".", "."},
		{"8", ".", ".", ".", ".", "9", ".", ".", ".", ".", ".", ".", "."},
	}
	count := ShortestPath(grid, "1", "9")
	fmt.Printf("count is %d", count)

	count1 := CalculateDistance(grid, "1", "9")
	fmt.Printf("count is %d", count1)
}
