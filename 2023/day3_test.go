package _2023

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result, result2 := Day3(dir + "/input/day3.txt")
	assert.Equal(t, result, 531932, "Is not 531932")
	assert.Equal(t, result2, 73646890, "Is not 73646890")

	str := "......#23...32.*1.45..8898.948"
	for _, num := range GetNumbersInLine(str) {
		fmt.Printf("- %d -", num)
	}

	fmt.Printf("\n")
	matrix := [][]string{
		{"4", "6", "7", "."},

		{".", ".", ".", "*"},
		{".", ".", "3", "5"},
	}

	fmt.Printf("%d search number", SearchNumber(matrix, 0, 2))
}
