package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	str string
}

type Cat struct {
	str string
}

func (d Dog) Speak() string{
	return "Bow!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	animal := []Animal{Dog{"Bow"}, Cat{"Meow"} }
	for _, a := range animal{
		fmt.Println(a.Speak())
	}
}
