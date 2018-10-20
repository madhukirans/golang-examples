package main

import (
	"sync"
	"fmt"
)

func main(){
	var wg sync.WaitGroup
	ch := make (chan  int)
	wg.Add(3)
	go func (){
		defer wg.Done()
		for i := range []int{1,2,3,4,5} {
			ch <- i
		}

	}()

	go func (){
		defer wg.Done()
		for i := range []int{1,2,3,4,5} {
			ch <- i
		}
	}()

	go func (){
		defer wg.Done()
		for i := range []int{1,2,3,4,5} {
			ch <- i
		}
	}()

	go func(){
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}



}
