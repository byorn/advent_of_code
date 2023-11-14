package util

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := Set{items: map[Item]bool{}}

	s.Add("A")

	s.Add("B")

	fmt.Println(s.items)
	fmt.Println(s.Get("C"))
	fmt.Println(s.Contains("B"))
}
