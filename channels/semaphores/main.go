package main

import "fmt"

func main() {
	ch:= make(chan int)
	sem := make(chan bool)

	go func() {
		for _, v := range []int {1,2,3,4,5} {
			ch <- v
		}
		sem <- true

	}()

	go func() {
		for _,v := range []int {6,7,8,9,10} {
			ch <- v
		}

		sem <- true
	}()

	go func(){
		<- sem
		<- sem
		close(ch)
	}()


	for i := range ch {
		fmt.Println(i)
	}

}
