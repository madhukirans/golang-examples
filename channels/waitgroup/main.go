package main

import (
	"sync"
<<<<<<< HEAD
	//"fmt"
//	"time"
=======
>>>>>>> Default Changelist
	"fmt"
)

func main(){
	var wg sync.WaitGroup
	ch := make (chan  int)
<<<<<<< HEAD
	wg.Add(2)
	go func (){
		defer wg.Done()
		for _, i := range []int{1,2,3,4,5} {
			fmt.Println("f1", i)
			ch <- i
			fmt.Println("f11", i)
			//time.Sleep(1*time.Second)
=======
	wg.Add(3)
	go func (){
		defer wg.Done()
		for i := range []int{1,2,3,4,5} {
			ch <- i
>>>>>>> Default Changelist
		}

	}()

	go func (){
		defer wg.Done()
<<<<<<< HEAD
		for _, i := range []int{1,2,3,4,5} {
			fmt.Println("f2", i)
			ch <- i
			fmt.Println("f22", i)
			//time.Sleep(1*time.Second)
		}

=======
		for i := range []int{1,2,3,4,5} {
			ch <- i
		}
>>>>>>> Default Changelist
	}()

	go func (){
		defer wg.Done()
<<<<<<< HEAD
		for _, i := range []int{1,2,3,4,5} {
			fmt.Println("f3", i)
			ch <- i
			fmt.Println("f33", i)
			//time.Sleep(1*time.Second)
		}

=======
		for i := range []int{1,2,3,4,5} {
			ch <- i
		}
>>>>>>> Default Changelist
	}()

	go func(){
		wg.Wait()
		close(ch)
	}()
<<<<<<< HEAD
	//fmt.Println(<-ch)
	//time.Sleep(3 * time.Second)
=======
>>>>>>> Default Changelist

	for i := range ch {
		fmt.Println(i)
	}



}
