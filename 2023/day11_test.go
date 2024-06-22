package _2023

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result, result2 := Day11(dir + "/input/day11.txt")

	assert.Equal(t, result, 9274989, "Is not 9274989")
	assert.Equal(t, result2, 357134560737, "Is not 357134560737")
}
