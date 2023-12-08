package _2023

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
	result, result2 := Day1(dir + "/input/day1.txt")
	assert.Equal(t, result, 55834, "Is not 55834")
	assert.Equal(t, result2, 53221, "Is not 53221")
}
