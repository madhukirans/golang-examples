package main

import "fmt"

func main(){
	//printNum(10)
	countUpAndDown(0, 10)

}

func countUpAndDown(start, end int) {
	fmt.Println(start)
	if  end == start  {
		return
	}
	countUpAndDown(start+1, end);
	fmt.Println(start);
}