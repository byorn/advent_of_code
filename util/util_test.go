package util

import (
	"fmt"
	"testing"
)

func TestUtil(t *testing.T) {
	items := []Item{"a", "b", "c"}

	fmt.Println(Contains(items, "c"))

	fmt.Println(InsertIntToIndex([]int{1, 2, 3}, 1, 9))

	fmt.Println(InsertStringArrayToIndex([][]string{{"a", "a", "a"}, {"b", "b", "b"}, {"c", "c", "c"}}, 1, []string{"#"}))
	fmt.Println(InsertStringToIndex([]string{"a", "a", "a"}, 1, "#"))
}
