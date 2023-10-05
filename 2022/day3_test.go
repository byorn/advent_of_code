package _2022

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDay3(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result := Day3_Part1(dir + "/input/day3.txt")
	assert.Equal(t, 7889, result, "Is not 7889")
	result = Day3_Part2(dir + "/input/day3.txt")
	assert.Equal(t, 2825, result, "Is not 2825")
}
