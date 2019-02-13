package main

import "fmt"

type Abc struct {
	x   int
	str string
}

func main() {
	var x Abc
	a := Abc{x: 1, str: "Madhu"}
	fmt.Println(a)
	func1(&a)
	fmt.Println(a)

	b := new(Abc)
	func1(b)
	//println(b)

	func2(x)
	fmt.Println(b)
}

func func1(a *Abc) {
	a.x = 2
	a.str = "kiran"
	fmt.Println(a)
}

func func2(a Abc) {
	a.x = 3
	a.str = "kiran1"
	fmt.Println(a)
}
