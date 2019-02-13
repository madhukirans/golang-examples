package main

import "fmt"

func main() {
	str := []byte("liril")
	len := len(str)
	largest := 0

	for i := 1; i < len; i++ {

		//For odd palandrome
		low := i - 1
		high := i + 1
		for low >= 0 && high < len && str[low] == str[high] {
			if largest < (high - low + 1) {
				largest = high - low + 1
			}
			low--
			high++
		}
	}

	fmt.Println(largest)

}
