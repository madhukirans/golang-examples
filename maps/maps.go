package main

import "fmt"

type MyStruct struct {
	str string
	error int
}
var map1 map[string] *MyStruct
func main(){
	map1 =  map[string] *MyStruct{}
	map1["a"] = &MyStruct{"a", 1}
	map1["b"] = &MyStruct{"b", 1}

	x()
	fmt.Println(map1["a"])
}

func x( ){
	a := map1["a"]
	a.error ++
}