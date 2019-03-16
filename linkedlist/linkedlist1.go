package main

import (
	"fmt"
	"strconv"
)

type Node1 struct{
	data int
	next *Node1
}

type LinkedList struct {
	head *Node1
	last *Node1 //without using this variable
	length int
}

func (n Node1) String() string{
	return strconv.Itoa(n.data) + " "

}

func (ll *LinkedList) append(data int){
	if ll.head == nil {
		ll.head = &Node1{data:data}
		ll.last = ll.head
		ll.length ++
		return
	} else {
		curr := ll.head
		for {
			if curr.next == nil {
				curr.next = &Node1{data:data}
				ll.last = ll.head.next
				ll.length ++
				return
			}  else {
				curr = curr.next
			}
		}
	}
}

func  (ll LinkedList)  String () string {
	if ll.head == nil {
		fmt.Println("Head:", ll.head )
		fmt.Println("Tail:", ll.last )
		return fmt.Sprintf("Head %s Tail %s\n", ll.head, ll.last)
	} else {
		current := ll.head
		str := fmt.Sprintf("Head: %s" , ll.head)
		for current.next != nil {
			str += fmt.Sprintf("Next:%s" , current.next)
			current = current.next
		}
		return str
	}
}

func main(){
	var ll = &LinkedList{}
	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(5)

	fmt.Println(ll)
}