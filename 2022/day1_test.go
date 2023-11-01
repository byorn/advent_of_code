package _2022

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result := Day1(dir + "/input/day1.txt")
	assert.Equal(t, result, 72478, "Is not 72478")
}
