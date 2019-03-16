package main

import "fmt"

type Parent struct {
	x int
}

type Child1 struct{
	Parent
}


func func1(obj Child1){
	fmt.Println(obj.x)
}

func main(){
	p := Parent{x:1}
	c := Child1{p}

	fmt.Println(c.x)
	func1(p)
}