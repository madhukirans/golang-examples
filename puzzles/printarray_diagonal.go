package main

import "fmt"


var a = [][]int {
{1,2,3,4},
{5,6,7,8},
{9,10,11,12},
{13,14,15,16},
}


func main(){
	fmt.Println(a)
	len := 4
	array := [4*4] int{}
	start := 0
	end := 4*4 - 1
	fmt.Print("  ")
	for i := 0; i< len; i++ {
		x := 0
		for j:=i ; j >= 0 ; j-- {
			fmt.Print(a[j][x], " ")
			array[start] = a[j][x]
			start++
			array[end] = a[len-1-j][len-1-x]
			end --
			x++
		}
	}

	for i := 1; i < len; i++ {
		x := len - 1
		for j:=i ; j < len ; j++ {
			fmt.Print(a[x][j], " ")
			x--
		}
	}

	fmt.Println("\n", array)
}