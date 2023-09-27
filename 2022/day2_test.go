package _2022

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDay2(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result := Day2_Part1(dir + "/input/day2.txt")
	assert.Equal(t, result, 8392, "Is not 8392")

	result = Day2_Part2(dir + "/input/day2.txt")
	assert.Equal(t, result, 10116, "Is not 10116")
}
