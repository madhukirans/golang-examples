package main

import (
	"sync"
	//"fmt"
//	"time"
	"fmt"
)

func main(){
	var wg sync.WaitGroup
	ch := make (chan  int)
	wg.Add(2)
	go func (){
		defer wg.Done()
		for _, i := range []int{1,2,3,4,5} {
			fmt.Println("f1", i)
			ch <- i
			fmt.Println("f11", i)
			//time.Sleep(1*time.Second)
		}

	}()

	go func (){
		defer wg.Done()
		for _, i := range []int{1,2,3,4,5} {
			fmt.Println("f2", i)
			ch <- i
			fmt.Println("f22", i)
			//time.Sleep(1*time.Second)
		}

	}()

	go func (){
		defer wg.Done()
		for _, i := range []int{1,2,3,4,5} {
			fmt.Println("f3", i)
			ch <- i
			fmt.Println("f33", i)
			//time.Sleep(1*time.Second)
		}

	}()

	go func(){
		wg.Wait()
		close(ch)
	}()
	//fmt.Println(<-ch)
	//time.Sleep(3 * time.Second)

	for i := range ch {
		fmt.Println(i)
	}



}
