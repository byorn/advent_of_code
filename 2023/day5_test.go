package _2023

import (
	"fmt"
	"math"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Dir %v does not exists", err)
	}
	result, result2 := Day5(dir + "/input/day5.txt")
	assert.Equal(t, result, 157211394, "Is not 157211394")
	assert.Equal(t, result2, 50855035, "Is not 50855035")
	assert.Equal(t, calculate(), 35, "Is not 35")
}

func calculate() int {
	sort.Sort(soilToSeeds1)
	sort.Sort(soilToFertilizer1)
	sort.Sort(fertilizerToWater1)
	sort.Sort(waterToLight1)
	sort.Sort(lightToTemperarure1)
	sort.Sort(tempToHumidity1)
	sort.Sort(humidityToLocation1)

	minVal := math.MaxInt

	for _, seed := range seeds1 {
		soil, err := soilToSeeds1.GetValue(seed)
		if err != nil {
			soil = seed
		}

		fertilzer, err := soilToFertilizer1.GetValue(soil)
		if err != nil {
			fertilzer = soil
		}

		water, err := fertilizerToWater1.GetValue(fertilzer)
		if err != nil {
			water = fertilzer
		}

		light, err := waterToLight1.GetValue(water)
		if err != nil {
			light = water
		}

		temp, err := lightToTemperarure1.GetValue(light)
		if err != nil {
			temp = light
		}

		humidity, err := tempToHumidity1.GetValue(temp)
		if err != nil {
			humidity = temp
		}

		location, err := humidityToLocation1.GetValue(humidity)
		if err != nil {
			location = humidity
		}

		if location < minVal {
			minVal = location
		}

	}

	return minVal
}

var seeds1 = []int{79, 14, 55, 13}

var soilToSeeds1 = Ranges{
	{50, 98, 2},
	{52, 50, 48},
}

var soilToFertilizer1 = Ranges{
	{0, 15, 37},
	{37, 52, 2},
	{39, 0, 15},
}

var fertilizerToWater1 = Ranges{
	{49, 53, 8},
	{0, 11, 42},
	{42, 0, 7},
	{57, 7, 4},
}

var waterToLight1 = Ranges{
	{88, 18, 7},
	{18, 25, 70},
}

var lightToTemperarure1 = Ranges{
	{45, 77, 23},
	{81, 45, 19},
	{68, 64, 13},
}

var tempToHumidity1 = Ranges{
	{0, 69, 1},
	{1, 0, 69},
}

var humidityToLocation1 = Ranges{
	{60, 56, 37},
	{56, 93, 4},
}

/*
Return the max total from all group totals
*/
