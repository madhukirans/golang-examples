package b

import (
	"fmt"
	//xxx "../a"
)



type B struct {
	x int
}

func (b *B) display(){
	fmt.Println(b.x)
}