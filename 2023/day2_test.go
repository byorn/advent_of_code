package _2023

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result, result2 := Day2(dir + "/input/day2.txt")
	assert.Equal(t, result, 2283, "Is not 2283")
	assert.Equal(t, result2, 53221, "Is not 53221")
}
