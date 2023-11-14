package util

import (
	"fmt"
	"testing"
)

func TestUtil(t *testing.T) {
	items := []Item{"a", "b", "c"}

	fmt.Println(Contains(items, "c"))
}
