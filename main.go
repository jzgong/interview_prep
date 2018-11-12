package main

import (
	"fmt"
	"goog_prep/bst"
)

func main() {
	intArr := []int{170, 45, 75, 90, 2, 1802, 66}
	// // strArr := []string{"a", "g", "b", "brew", "agave", "go", "goes"}

	// // intArr = sorting.RadixInt(intArr)
	// // fmt.Println(intArr)
	// // fmt.Println(sorting.RadixStr(strArr, 0))

	// // fmt.Println(search.BinarySearch([]int{2, 3}, 3))

	root := bst.Node{Value: intArr[0]}

	for _, ele := range intArr[1:] {
		root.Insert(ele)
		fmt.Println(root.BFSDisplay())
	}

	var currPath []*bst.Node
	var pathArr [][]*bst.Node

	root.AllPaths(currPath, &pathArr)

	fmt.Println(pathArr)

	// var g graphs.Graph

	// graphs.FillGraph(&g)

	// visited := map[*graphs.Node]bool{}
	// fmt.Println(g.CycleCheck(g.Nodes[0], visited))

}
