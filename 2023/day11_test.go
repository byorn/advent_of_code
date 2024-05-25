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
	result, result2 := Day11(dir + "/input/scrap.txt")

	arr := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(arr)

	assert.Equal(t, result, 55834, "Is not 551834")
	assert.Equal(t, result2, 53221, "Is not 53221")
}
