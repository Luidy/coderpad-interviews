package main

import (
	"fmt"
	"math"
)

type NodeWithSum struct {
	Node *Node
	Sum  int
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func minSum(root *Node) int {
	var queue []*NodeWithSum
	queue = append(queue, &NodeWithSum{Node: root, Sum: root.Value})
	min := math.MaxInt64

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		curSum := top.Sum

		if top.Node.Left == nil && top.Node.Right == nil {
			if curSum < min {
				min = curSum
				continue
			}
		}

		if top.Node.Left != nil {
			queue = append(queue, &NodeWithSum{Node: top.Node.Left, Sum: curSum + top.Node.Left.Value})
		}
		if top.Node.Right != nil {
			queue = append(queue, &NodeWithSum{Node: top.Node.Right, Sum: curSum + top.Node.Right.Value})
		}
	}
	return min
}

func main() {
	root := &Node{Value: 10, Left: &Node{Value: 5, Left: nil, Right: &Node{Value: 2, Left: nil, Right: nil}},
		Right: &Node{Value: 5, Left: nil, Right: &Node{Value: 1, Left: &Node{Value: -1, Left: nil, Right: nil}, Right: nil}}}
	fmt.Println(minSum(root))
}
