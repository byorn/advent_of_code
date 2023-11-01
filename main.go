package main

import (
	"fmt"
	"os"

	_2022 "github.com/byorn/advent_of_code/2022"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("Dir %v does not exists", err)
	}
	result := _2022.Day4(dir + "/2022/input/day4.txt")

	fmt.Println(result, err)
}
