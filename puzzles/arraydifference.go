package main

import "fmt"

//Find the elements which are not exists in second array (A-B)
func findMissingFormA(a, b []int, n, m int) {
	for i := 0; i < n; i++ {
		var j int
		for j = 0; j < m; j++ {
			if (a[i] == b[j]) {
				break;
			}
		}

		if (j == m) {
			//cout << a[i] << " ";
			fmt.Println(a[i])
		}
	}
}



// Driver code
func main() {
	a := []int{1, 2, 6, 3, 4, 5,10}
	b := []int{2, 4, 3, 1, 0,11}
	n := len(a)
	m := len(b) /// len(b[1])
	findMissingFormA(a, b, n, m)
	findMissingFormA(b, a, m, n)
}
