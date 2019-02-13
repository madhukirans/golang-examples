package main

import "fmt"

type animal interface {
	eat()
	walk()
}
type tail struct {
}

type trunk struct {
}

func (t *trunk) wave() {
	fmt.Println("I am trunking")
}
func (t *tail) wave() {
	fmt.Println("I am waving")
}

type elephant struct {
	// Nameless Structure gives directly its attributes and methods
	t1 tail
	trunk
}

func (e *elephant) eat() {
	fmt.Println("Elephant is eating")
}
func (e *elephant) walk() {
	fmt.Println("Elephant is walking")
}

func main() {
	e := elephant{}
	e.wave()
	e.t1.wave()
}
