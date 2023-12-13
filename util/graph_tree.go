package util

type GraphTreeNode struct {
	value    string
	children []*GraphTreeNode
	visited  bool
}

// using recursion, or you can use a stack
func DFS(gtn *GraphTreeNode) {
	gtn.visited = true

	for _, c := range gtn.children {
		if !c.visited {
			DFS(c)
		}
	}
}

func BFS(gtn *GraphTreeNode) {
	queue := []*GraphTreeNode{gtn}

	popped := queue[0]

	popped.visited = true

	queue = queue[1:]

	for len(queue) != 0 {
		for _, tn := range gtn.children {
			queue = append(queue, tn)
		}
	}
}
