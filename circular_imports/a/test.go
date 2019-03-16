package a

import (
	"fmt"
)

type A struct {
	x int
}

func (a *A) display(){
	fmt.Println(a.x)
}