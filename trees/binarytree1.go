package main

import "fmt"

type Node1 struct {
	data int
	left *Node1
	right *Node1
}


func main(){
	var root *Node1

	arr := []int{10,3,11,2,4,66,9,11}
	for _, i:= range arr{
		root.insert(i)
	}
	root.display()
	fmt.Println(arr1)
}

func (n *Node1) insert(data int){
	if n == nil{
		n = &Node1{data: data}
	} else if n.data <= data{
		if n.left == nil {
			n.right = &Node1{data: data, left: nil, right: nil}
		}
		n.left.insert(data)
	} else if n.data > data {
		if n.right == nil {
			n.left = &Node1{data: data, left: nil, right: nil}
		}
		n.right.insert(data)
	}

	fmt.Println(n.data, " ")
}

var arr1 []int
func (n *Node1) display() {
	if n == nil {
		return
	}

	n.left.display()
	arr1 = append(arr1, n.data)
	n.right.display()
}