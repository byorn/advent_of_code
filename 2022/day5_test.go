package _2022

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("Dir %v does not exists", err)
	}
	result1 := Day5(dir + "/input/day5.txt")
	assert.Equal(t, result1, "MQSHJMWNH", "Is not MQSHJMWNH")
}
