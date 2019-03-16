package main

import "fmt"

func fibonacci(n int) int{
	if n ==0 {
		return 0
	}
	if n ==1 || n==2 {
		return 1
	}
	f1 := 1
	f2 := 1
	for i:=3; i <n; i++{
		temp:= f1 + f2
		f1, f2 = f2, temp
	}

	fmt.Println(f2)
	return f2
}

func main(){
  fibonacci(10)
}