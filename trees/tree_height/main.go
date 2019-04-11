package main

import "fmt"

// you can also use imports, for example:
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func treeHeight(subT *Tree) int{
	if subT == nil {
		return 0
	}

	return 1 + max (treeHeight(subT.L), treeHeight(subT.R))
}

func max (x, y int) int{
	if x > y{
		return x
	}

	return y
}

func Solution(T *Tree) int {
	// write your code in Go 1.4
	return max(treeHeight(T.L), treeHeight(T.R))

}

func main() {
	T := &Tree{}
	// (5, (3, (20, None, None), (21, None, None)), (10, (1, None, None), None))
	fmt.Println(Solution(T))
}
