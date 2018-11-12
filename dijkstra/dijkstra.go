package dijkstra

import (
	"fmt"
)

type Graph struct {
	Edges map[string]map[string]int
}

func (g *Graph) dijkstra(start string, end string) (map[string]int, map[string]string) {
	if _, exist := g.Edges[start]; !exist {
		fmt.Println("root not found")
	}
	if _, exist := g.Edges[end]; !exist {
		fmt.Println("dest not found")
	}

	startDist := map[string]int{start: 0}
	queue := map[string]int{start: 0}
	pred := map[string]string{}

	unvisited := map[string]bool{}
	for key := range g.Edges {
		unvisited[key] = true
	}

	for len(unvisited) > 0 {
		var currStr string
		currVal := -1

		for key := range queue {
			if currVal == -1 || queue[key] <= currVal {
				currStr = key
				currVal = queue[key]
			}
		}

		startDist[currStr] = currVal
		delete(queue, currStr)
		delete(unvisited, currStr)

		// if currStr == end {
		// 	break
		// }

		for edge := range g.Edges[currStr] {
			if _, exists := unvisited[edge]; exists {
				tempDist := startDist[currStr] + g.Edges[currStr][edge]
				if _, exists := queue[edge]; !exists {
					queue[edge] = tempDist
					pred[edge] = currStr
				} else if tempDist < queue[edge] {
					queue[edge] = tempDist
					pred[edge] = currStr
				}
			}
		}
	}

	return startDist, pred
}

func (g *Graph) shortestPath(start, end string) []string {
	_, pred := g.dijkstra(start, end)
	v := end
	path := []string{v}
	for v != start {
		v = pred[v]
		path = append(path, v)
	}

	return path

}

func Test() {
	var g Graph

	distances := map[string]map[string]int{
		"b": map[string]int{"a": 5, "d": 1, "g": 2},
		"a": map[string]int{"b": 5, "d": 3, "e": 12, "f": 5},
		"d": map[string]int{"b": 1, "g": 1, "e": 1, "a": 3},
		"g": map[string]int{"b": 2, "d": 1, "c": 2},
		"c": map[string]int{"g": 2, "e": 1, "f": 16},
		"e": map[string]int{"a": 12, "d": 1, "c": 1, "f": 2},
		"f": map[string]int{"a": 5, "e": 2, "c": 16},
	}

	g.Edges = distances

	fmt.Println(g.dijkstra("a", "b"))
	fmt.Println(g.shortestPath("a", "b"))

}
