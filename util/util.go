package util

import (
	"sort"
)

type Predicate func(int) bool

func Filter(arr []int, pre Predicate) []int {
	var result []int
	for _, v := range arr {
		if pre(v) {
			result = append(result, v)
		}
	}
	return result
}

func InsertIntToIndex(arr []int, index int, number int) []int {
	slice1 := arr[:index]
	slice2 := arr[index:]
	return append(slice1, append([]int{number}, slice2...)...)
}

func InsertStringToIndex(arr []string, index int, value string) []string {
	slice1 := arr[:index]
	slice2 := arr[index:]
	return append(slice1, append([]string{value}, slice2...)...)
}

func InsertStringArrayToIndex(arr [][]string, index int, valueToInsert []string) [][]string {
	slice1 := arr[:index]
	slice2 := arr[index:]
	return append(slice1, append([][]string{valueToInsert}, slice2...)...)
}

func Contains(items []Item, target string) bool {
	for _, item := range items {
		if str, ok := item.(string); ok && target == str {
			return true
		}
	}
	return false
}

type KeyValue struct {
	Key   string
	Value int
}

// 10,9,8,7 desc order
func SortMapByValues(mymap map[string]int) []KeyValue {
	keyValues := []KeyValue{}

	for key, value := range mymap {
		keyValues = append(keyValues, KeyValue{key, value})
	}

	sort.Slice(keyValues, func(i, j int) bool {
		return keyValues[i].Value > keyValues[j].Value
	})
	return keyValues
}
