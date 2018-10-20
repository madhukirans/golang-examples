package main

import (
	"fmt"

	"time"
)

func main(){
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)

	time.Sleep(time.Second * 10)
}

func producer(ch chan int){
	for i := range []int{1,2,3,4,5,6,7} {
		fmt.Println("Before produce:", i)
		ch <- i
		//time.Sleep(time.Millisecond * 500)
		fmt.Println("After produce:", i)

	}
}

func consumer(ch chan int) {
	for {
		fmt.Println("Before Consume:")
		fmt.Println(<-ch)
		//time.Sleep(time.Millisecond * 500)

		fmt.Println("After Consume:")

	}
}