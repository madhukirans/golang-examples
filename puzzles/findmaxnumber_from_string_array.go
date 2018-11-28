package main

import "fmt"

var str = &[]int{78, 96, 98, 95, 772, 77, 779}

func main() {
	//for i :=0; i < len(*str); i++ {
	fmt.Println(findHighestNumberInIndex(str, 0))
	//}
}

func findHighestNumberInIndex(a *[]int, pos int) (int,int){
	highest := 0
	index := []int {}
	for _, val := range *str {
		if firstDigit(val) > highest {
			highest = firstDigit(val)
		}
	}

	for _, val := range *str {
		if firstDigit(val) == highest{
			index = append(index, val)
		}
	}





	return index, highest
}

func firstDigit(x int) int {
	for (x > 9) {
		x /= 10;
	}
	return x;
}
