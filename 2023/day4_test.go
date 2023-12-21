package _2023

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result, result2 := Day4(dir + "/input/day4.txt")
	assert.Equal(t, result, 23673, "Is not 23673")
	assert.Equal(t, result2, 12263631, "Is not 12263631")

	str := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	x, y := GetTwoArrays(str)
	fmt.Printf("%v get two arrays %v", x, y)
}
