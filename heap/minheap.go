package main

import "fmt"

// 2 3 10 12 5
///

var arr [10]int
var size int

func main() {
	fmt.Println(arr)
	push(10)
	push(9)
	push(11)
	push(7)
	push(8)
	push(12)
	push(1)
	fmt.Println(arr)

	fmt.Println(pop(), " -  ", arr, size)
	fmt.Println(pop(), " -  ", arr, size)
	fmt.Println(pop(), " -  ", arr, size)
	fmt.Println(pop(), " -  ", arr, size)
	fmt.Println(pop(), " -  ", arr, size)
	fmt.Println(pop(), " -  ", arr, size)
	fmt.Println(pop(), " -  ", arr, size)

}

//	  12
//   /      \
//  9	       11
///  \      /   \
//7	8    10    1

func push(x int) error {
	//if len(arr) == size {
	//	return fmt.Errorf("Full")
	//}

	pos := size
	arr[pos] = x

	for pos > 0 {
		parent := (pos+1)/2 - 1
		if arr[parent] >= x {
			break
		}
		arr[parent], arr[pos] = arr[pos], arr[parent]
		pos = parent
	}
	size ++
	return nil
}

func pop() int {
	popValue := arr[0]
	arr[0] = arr[size]
	size--

	pos := 0
	for pos < size/2 {
		leftChild := pos*2 + 1
		rightChild := leftChild + 1

		if rightChild < size  && arr[leftChild] < arr[rightChild] {
			arr[pos], arr[rightChild] = arr[rightChild], arr[pos]
			pos = rightChild
		} else {
			if arr[pos] >= arr[leftChild] {
				break
			}
			arr[pos], arr[leftChild] = arr[leftChild], arr[pos]
			pos = leftChild

		}
	}

	return popValue
}
