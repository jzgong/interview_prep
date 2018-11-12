package bst

import (
	"errors"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int) error {
	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	if value == n.Value {
		return nil
	} else if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
			return nil
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
			return nil
		} else {
			n.Right.Insert(value)
		}
	}

	return nil
}

func (n *Node) Find(value int) bool {
	if n == nil {
		return false
	}

	if n.Value == value {
		return true
	} else if n.Value < value {
		return n.Right.Find(value)
	} else {
		return n.Left.Find(value)
	}
}

func (n *Node) BFSDisplay() []int {
	var ans []int
	queue := []*Node{n}

	for len(queue) > 0 {
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		ans = append(ans, queue[0].Value)
		queue = queue[1:]
	}
	return ans
}

func (n *Node) DFSDisplay() []int {
	if n == nil {
		return []int{}
	}

	ans := []int{}
	stack := []*Node{n}

	for len(stack) != 0 {
		if stack[0].Right != nil {
			stack = append(stack[:1], append([]*Node{stack[0].Right}, stack[1:]...)...)
		}
		if stack[0].Left != nil {
			stack = append(stack[:1], append([]*Node{stack[0].Left}, stack[1:]...)...)
		}
		ans = append(ans, stack[0].Value)
		stack = stack[1:]
	}

	return ans
}

func (n *Node) MaxHeightv2() int {
	var lvlQueue [][]*Node

	lvlQueue = append(lvlQueue, []*Node{n})

	height := 0
	for len(lvlQueue) > 0 {
		var nextLvl []*Node
		height++
		for _, tempNode := range lvlQueue[0] {
			if tempNode.Left != nil {
				nextLvl = append(nextLvl, tempNode.Left)
			}
			if tempNode.Right != nil {
				nextLvl = append(nextLvl, tempNode.Right)
			}
		}
		lvlQueue = lvlQueue[1:]
		if len(nextLvl) == 0 {
			return height
		} else {
			lvlQueue = append(lvlQueue, nextLvl)
		}
	}

	return height
}

func (n *Node) MinHeightv2() int {
	var nodeQueue [][]*Node

	height := 0

	nodeQueue = append(nodeQueue, []*Node{n})

	for len(nodeQueue) > 0 {
		var nodeLevel []*Node
		height++
		for _, ele := range nodeQueue[0] {
			if ele.Left == nil && ele.Right == nil {
				return height
			}
			if ele.Left != nil {
				nodeLevel = append(nodeLevel, ele.Left)
			}
			if ele.Right != nil {
				nodeLevel = append(nodeLevel, ele.Right)
			}

			nodeQueue = nodeQueue[1:]
			nodeQueue = append(nodeQueue, nodeLevel)
		}
	}

	return height
}

func intMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func intMin(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func (n *Node) MaxHeight() int {
	if n == nil {
		return 0
	}

	return 1 + intMax(n.Left.MaxHeight(), n.Right.MaxHeight())
}

func (n *Node) MinHeight() int {
	if n == nil {
		return 0
	}

	return 1 + intMin(n.Left.MinHeight(), n.Right.MinHeight())
}

func (n *Node) AllPaths(currPath []*Node, pathArr *[][]*Node) {
	currPath = append(currPath, n)

	if n.Left == nil && n.Right == nil {
		*pathArr = append(*pathArr, currPath)
		return
	}

	if n.Left != nil {
		n.Left.AllPaths(currPath, pathArr)
	}
	if n.Right != nil {
		n.Right.AllPaths(currPath, pathArr)
	}
}
