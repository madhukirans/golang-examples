package main

import "fmt"

type Node struct {
	next   *Node
	data   int
	length int
	last *Node
}

func main() {
	var start = &Node{}
	var last = start
	last.Append(1)
	last.Append(2)
	last.Append(3)
	last.Append(4)
	start.Display()


}

func (last *Node) Append(data int) {
	if last.next == nil {
		last.data = data
	}
		last.next = &Node{}
		last = last.next
}

func (start *Node) Display(){
	var currNode = start
	for currNode.next != nil {
		fmt.Println(currNode.data)
		currNode = currNode.next
	}
}
