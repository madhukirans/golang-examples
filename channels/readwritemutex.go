package main

import (
	"sync"
	"fmt"
	"strings"
	"time"
)

func main(){
	rchan := make(chan int)
	wchan := make(chan int)
	mutex := sync.RWMutex{}

	go func(){
		r := 0
		w := 0
		for {
			select {
			case x := <-rchan:
				r += x
			case x := <-wchan:
				w += x
			}

			fmt.Printf("%s %s\n", strings.Repeat("R", r), strings.Repeat("W", w))
		}
	}()

	wg := sync.WaitGroup{}

	for i:=0; i <10 ; i++{
		wg.Add(1)
		go func (){
			//If you don't give the RLock then the following snippet will execute irrespective of write lock.
			// That means below code will execute when line 50 (wchan <- 1) is running
			mutex.RLock()
			rchan <- 1
			time.Sleep(time.Duration(1000) * time.Millisecond)
			rchan <- -1
			mutex.RUnlock()
			wg.Done()
		}()
	}

	for i:=0; i <5 ; i++{
		wg.Add(1)
		go func (){
			mutex.Lock()
			wchan <- 1
			time.Sleep(time.Duration(1000) * time.Millisecond)
			wchan <- -1
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}