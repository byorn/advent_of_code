package util

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()
	g.Add("a", "b")

	g.Add("a", "e")

	g.Add("a", "d")
	g.Add("a", "f")
	g.Add("b", "c")
	g.Add("c", "x")

	g.Add("e", "k")
	g.PrintDFS("a")
	fmt.Println("--------")
	g.PrintBFS("a")
}
