package main

import (
	"fmt"
	"math"
)

type Node struct {
	top  int
	left int
}

func matrixProduct(m [][]int, node *[3][3]Node) int {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if i == 0 && j == 0 {
				node[i][j].top = m[i][j]
				node[i][j].left = m[i][j]
			} else if i == 0 && j > 0 {
				node[i][j].top = node[i][j-1].left * m[i][j]
				node[i][j].left = m[i][j]
			} else if i > 0 && j == 0 {
				node[i][j].top = m[i][j]
				node[i][j].left = node[i-1][j].top * m[i][j]
			} else {
				node[i][j].top = int(math.Max(float64(node[i][j-1].top), float64(node[i][j-1].left))) * m[i][j]
				node[i][j].left = int(math.Max(float64(node[i-1][j].top), float64(node[i-1][j].left))) * m[i][j]
			}
		}
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			fmt.Print(node[i][j], ",\t ")
		}
		fmt.Println()
	}
	return 0
}

func main() {
	var a = [][]int{
		{1, 2, 3},
		{4, -5, 6},
		{7, 8, -9},
	}

	var node = [3][3]Node{}

	fmt.Print(matrixProduct(a, &node))
}
