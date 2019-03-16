package main

import (
	"fmt"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		for _, v := range []int{1, 2, 3, 4} {
			c1 <- v
			//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c1)
	}()

	go func() {
		for _, v := range []int{6, 7, 8, 9} {
			c2 <- v
			//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c2)
	}()

	c3 := merge(c1, c2)
	for i := range c3 {
		fmt.Println(i)
	}
}

func merge(c1, c2 <-chan int) <-chan int {
	outChan := make(chan int)
	go func() {
		defer close (outChan)
		for c1 != nil || c2 != nil {
			select {
			case x, ok := <-c1:
				if !ok {
					c1 = nil
					continue
				}
				outChan <- x
			case x, ok := <-c2:
				if !ok {
					c2 = nil
					continue
				}
				outChan <- x
			}
		}
	}()
	return outChan
}
