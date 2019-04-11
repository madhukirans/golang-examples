package main

import (
	"fmt"

	"time"
<<<<<<< HEAD
	"reflect"
)

func main(){

	i := [10]int{1,2,3}

	j := make ([]int , 10)

	fmt.Println(reflect.TypeOf(j))
	//i = append(i, 1)
	j = append(j, i[0:]...)

	fmt.Println(j)
	i[4] = 12
	fmt.Println(j, i)

=======
)

func main(){
>>>>>>> Default Changelist
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
<<<<<<< HEAD
}

Hi Malaravan,
Here are the stock options, which I received from Oracle. Please find the attachment.
According to today's valuation I have $15,049  * 70Rs = 1053430/- Indian rupees.
=======
}
>>>>>>> Default Changelist
