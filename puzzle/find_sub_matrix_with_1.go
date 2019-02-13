package main

import (
	"fmt"
)

func findSubMatrix(M [][]int) {

	maxValue := 0
	maxRow := 0
	maxCol := 0

	rows := len(M)
	cols := len(M[0])
	var S [][]int
	S = make([][]int, rows)
	for i := range S {
		S[i] = make([]int, cols)
		S[i][0] = M[i][0]
	}

	for i := range S[0] {
		S[0][i] = M[0][i]
	}

	fmt.Println(len(S), len(S[0]))

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if M[i][j] == 0 {
				S[i][j] = 0
			} else {
				S[i][j] = min(S[i][j-1], S[i-1][j], S[i-1][j-1]) + 1
				if maxValue < S[i][j] {
					maxValue = S[i][j]
					maxRow = j
					maxCol = i
				}
			}
		}
	}

	for i := range S {
		fmt.Println(S[i])
	}

	fmt.Println(maxValue, maxRow, maxCol)
	for i := maxRow; i > 0; i-- {
		for j := maxCol; j > 0; j-- {
			fmt.Print(M[i][j], " ")
		}
		fmt.Print("\n")
	}

}

func min(a ...int) int {
	min := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}

	return min
}

func main() {
	var M [][]int = [][]int{
		{0, 1, 1, 0, 1},
		{1, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
	}

	findSubMatrix(M)
}
