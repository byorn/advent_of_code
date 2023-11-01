package _2022

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("Dir %v does not exists", err)
	}
	result1, result2 := Day4(dir + "/input/day4.txt")
	assert.Equal(t, result1, 524, "Is not 524")
	assert.Equal(t, result2, 798, "Is not 798")
}
