package main

import (
	"fmt"
)

type TokenType uint16

//const (
//	A TokenType = iota
//	B
//	C
//	D
//)

type Base struct {
	a, b int
}

type Chield1 struct {
	Base
	a, p int
}

type Chield2 struct {
	Base
	a, q int
}

type Dimond struct {
	Chield1
	Chield2
	a, x int
}

func (this Base) func1() string {
	return fmt.Sprintf("in func1 base: %d ^d", this.a, this.b)
}

func (this Chield1) func1() string {
	return fmt.Sprintf("in Func1 Chiled1 a=&d b=%d p=&d", this.a, this.b, this.p)
}

func (this Chield2) func1() string {
	return fmt.Sprintf("in func1 Child2 a=%d b=%d q=%d", this.a, this.b, this.q)
}

func (this Dimond) func1() string {
	return fmt.Sprintf("in Func1 dimond a=%d b=%d p=%d q=%d x=%d", this.a, this.Chield1.b, this.p, this.q, this.x)
}

func main() {
	//	fmt.Print(A,B,C,D)
	base := Base{111, 222}
	base.func1()

	derived := Dimond{Chield1{Base{1, 2}, 3, 4}, Chield2{Base{10, 11}, 12, 13}, 20, 21}
	fmt.Println(derived.a)
	fmt.Println(derived.func1())
}
