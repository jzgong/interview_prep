package graphs

import (
	"fmt"
)

type Node struct {
	Val string
}

type Graph struct {
	Nodes []*Node
	Edges map[Node][]*Node
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}

	g.Edges[*n1] = append(g.Edges[*n1], n2)
	// // for undirect graph
	// g.Edges[*n2] = append(g.Edges[*n2], n1)
}

func (g *Graph) BFS() {
	nodeQueue := []Node{}
	nodeQueue = append(nodeQueue, *g.Nodes[0])
	visited := map[Node]bool{}

	for len(nodeQueue) > 0 {
		currNode := nodeQueue[0]
		visited[currNode] = true
		near := g.Edges[currNode]

		for i := range near {
			if _, exists := visited[*near[i]]; !exists {
				nodeQueue = append(nodeQueue, *near[i])
				visited[*near[i]] = true
			}
		}

		fmt.Println(currNode.Val)
		nodeQueue = nodeQueue[1:]
	}
}

func (g *Graph) DFSstack() {
	nodeQueue := []Node{}
	nodeQueue = append(nodeQueue, *g.Nodes[0])
	visited := map[Node]bool{*g.Nodes[0]: true}

	for len(nodeQueue) > 0 {

		near := g.Edges[nodeQueue[0]]

		for i := range near {
			if _, exists := visited[*near[i]]; !exists {
				nodeQueue = append(nodeQueue[:1], append([]Node{*near[i]}, nodeQueue[1:]...)...)
				visited[*near[i]] = true
			}
		}

		fmt.Println(nodeQueue[0].Val)
		nodeQueue = nodeQueue[1:]
	}
}

func (g *Graph) CycleCheck(curr *Node, visited map[*Node]bool) bool {

	near := g.Edges[*curr]
	visited[curr] = true

	for i := range near {
		if _, exist := visited[near[i]]; exist {
			return true
		} else {
			if g.CycleCheck(near[i], visited) {
				return true
			}
		}
	}

	return false

}

func FillGraph(g *Graph) {
	nA := Node{Val: "A"}
	nB := Node{Val: "B"}
	nC := Node{Val: "C"}
	nD := Node{Val: "D"}
	nE := Node{Val: "E"}
	nF := Node{Val: "F"}

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nB, &nC)
	g.AddEdge(&nA, &nD)
	g.AddEdge(&nD, &nA)
}
