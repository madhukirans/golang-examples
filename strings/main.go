package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " madhu kiran seelam   "
	s := strings.Split(str, " ")
	fmt.Println(s, len(s))
}
