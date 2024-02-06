package util

import (
	"sort"
)

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
