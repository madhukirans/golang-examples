package main

import "fmt"

type Abc struct {
	x int
	y int
	map1 map[string]string
}

func (a Abc) String () string{
	return "madhu"
}
func main(){
	a := Abc{1,2, nil}
	fmt.Println(a)
}