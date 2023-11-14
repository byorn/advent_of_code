package _2022

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("Dir %v does not exists", err)
	}
	result1, result2 := Day6(dir + "/input/day6.txt")
	assert.Equal(t, result1, 1538, "Is not 1538")
	assert.Equal(t, result2, 2315, "Is not 2315")
}
