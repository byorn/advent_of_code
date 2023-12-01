package _2022

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("Dir %v does not exists", err)
	}
	result1, result2 := Day7(dir + "/input/day7.txt")
	assert.Equal(t, result1, 1443806, "Is not 1443806")
	assert.Equal(t, result2, 942298, "Is not 942298")
}
