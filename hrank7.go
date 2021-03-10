package main

import (
	"fmt"
	"math"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func inorder(root *Node) {
	if root == nil {
		return
	}

	inorder(root.Left)
	fmt.Println(root.Data)
	inorder(root.Right)
}

func postorder(root *Node) {
	if root == nil {
		return
	}

	postorder(root.Right)
	fmt.Println(root.Data)
	postorder(root.Left)
}

func main() {

	var i uint64
	i = 1<<64 - 1
	fmt.Println(i)
	fmt.Println(math.MaxUint64)

	r := Node{
		Data: 0,
	}

	r.Left = &Node{
		Data: 1,
		Left: &Node{
			Data:  5,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

	r.Right = &Node{
		Data:  2,
		Left:  nil,
		Right: nil,
	}

	postorder(&r)
}
