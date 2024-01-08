package _2023

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result, result2 := Day6(dir + "/input/day6.txt")
	assert.Equal(t, result, 503424, "Is not 503424")
	assert.Equal(t, result2, 32607562, "Is not 32607562")
}
