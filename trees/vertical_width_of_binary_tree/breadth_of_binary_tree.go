package main

import (
	"fmt"
	"sort"
)

type Node struct {
	X int
	left *Node
	right *Node
}

//        1
//     2     3
//  4

func WidthOfTree(tree *Node, level int) int {
	if tree == nil {
		return 0
	}

	if tree.left != nil {
		WidthOfTree(tree *Node, level int)
	}

	sort.Sort()
	return 0
}

func main(){
	/*
        Constructed bunary tree is:
              1
            /  \
           2    3
          / \    \
         4   5    8
                 / \
                6   7 */
	tree := &Node{1,nil,nil}
	tree.left = &Node{2,nil,nil}
	tree.right = &Node{3,nil,nil}
	tree.left.left = &Node{4,nil,nil}
	tree.left.right = &Node{5,nil,nil}
	tree.right.right = &Node{8,nil,nil}
	tree.right.right.left = &Node{6,nil,nil}
	tree.right.right.right = &Node{7,nil,nil}
	fmt.Println(WidthOfTree(node, 0))
}
