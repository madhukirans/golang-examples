package main

import (
	"fmt"
	"github.com/madhukirans/golang-examples/circular_imports/Temp"
)

func main(){
	x := new(Temp.X)
	fmt.Println(x)
}
