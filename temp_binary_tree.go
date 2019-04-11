package main

import (
	"fmt"
	"errors"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

var root *Node

func main() {
	var data = []int{5, 4, 3, 6, 7, 2, 9}
	root = &Node{6, nil, nil}
	for _, i := range data {
		root.insert(i)
		//root.insert1(i)
	}

	root.Display()
}

func (node *Node) insert1(data int){

	for node != nil {
		if node.data >= data {
			node = node.left
		} else {
			node = node.right
		}
	}

	fmt.Print(data)
	node = &Node{data, nil, nil}
}

func (node *Node) insert(data int) {
	if node == nil {
		node = &Node{data, nil, nil}
		fmt.Println("root", data)
		return
	}

	if node.data >= data {
		if node.left == nil {
			node.left = &Node{data, nil, nil}
			fmt.Println("left", data)
			return
		} else {
			node.left.insert(data)
		}
	}

	if node.data < data {
		if node.right == nil {
			node.right = &Node{data, nil, nil}
			fmt.Println("right", data)
			return
		} else {
			node.right.insert(data)
		}
	}
}

type Stack struct {
	node []*Node
}

func (stack *Stack) Push(x *Node) {
	stack.node = append(stack.node, x)
}

func (stack *Stack) Pop() (*Node, error) {
	if stack.node == nil || len(stack.node) == 0 {
		return nil, errors.New("empty")
	} else {
		node := stack.node[len(stack.node)-1]
		stack.node = stack.node[:len(stack.node)-1]
		return node, nil
	}
}

func (node *Node) Display(){
	var stack Stack
	stack.Push(node)
	for len(stack.node) > 0 {
		printNode, err := stack.Pop()

		if err == nil {
			//fmt.Println("in", len(stack.node))
			if printNode.right != nil {
				stack.Push(printNode.right)
				//fmt.Println("right", printNode.right)
			}
			if printNode.left != nil {
				stack.Push(printNode.left)
				//fmt.Println("left", printNode.left)
			}


		}

		//if printNode.left == nil || printNode.right == nil {
		fmt.Println(printNode.data)
		//}
	}
}
