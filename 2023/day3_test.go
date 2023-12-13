package _2023

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result, result2 := Day3(dir + "/input/scrap.txt")
	assert.Equal(t, result, 2283, "Is not 2283")
	assert.Equal(t, result2, 53221, "Is not 53221")

	str := "......23...32..1.45..8898.948"
	for _, num := range GetNumbersInLine(str) {
		fmt.Printf("- %d -", num)
	}
}
