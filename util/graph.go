package util

import (
	"fmt"
)

type Graph struct {
	vertices map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		vertices: map[string][]string{},
	}
}

func (g *Graph) Add(from string, to string) {
	g.vertices[from] = append(g.vertices[from], to)
}

func (g *Graph) PrintDFS(start string) {
	visited := make(map[string]bool)

	stack := []string{start}

	for len(stack) != 0 {
		popped := stack[len(stack)-1]
		visited[popped] = true
		fmt.Printf("%s", popped)
		stack = stack[:len(stack)-1]
		for _, v := range g.vertices[popped] {
			//	if !visited[v] {
			stack = append(stack, v)
			//}
		}
	}
}

func (g *Graph) PrintBFS(start string) {
	visited := make(map[string]bool)
	q := []string{start}

	for len(q) != 0 {
		shifted := q[0]
		visited[shifted] = true
		fmt.Printf("%s", shifted)
		q = q[1:]
		for _, neighbr := range g.vertices[shifted] {
			// if !visited[neighbr] {
			q = append(q, neighbr)
			//	}
		}
	}
}
