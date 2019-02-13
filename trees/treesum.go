package main

import "fmt"

type NodeT struct {
	left  *NodeT
	right *NodeT
	data  int
	count int
}

func main() {
	data := []int{4, 2, 5, 7, 1, 9, 8, 2, 8,8,2,8,4}
	var root = &NodeT{data: data[0]}
	fmt.Println("root", root)

	for i := 1; i < len(data); i++ {
		root.insert(data[i])
	}

	fmt.Println(root)

	root.Display()

	fmt.Println(root.findGivenSumInSubTree(32))
}

func (n *NodeT) findGivenSumInSubTree(sum int) bool{
	if n == nil {
		return false
	}

	val  :=  n.data
	left := 0
	right := 0
	if n.left != nil{
		left = left + n.left.getSum()
	}
	fmt.Println("left:", val + left)
	if n.right != nil {
		right = right + n.right.getSum()
	}
	fmt.Println("right:", val + right)

	if sum == (val + left) || sum == (val+right){
		return true,
	}

	n.left.findGivenSumInSubTree(sum)
	n.right.findGivenSumInSubTree(sum)

	return false
}

func (n *NodeT) getSum() int {
	if n == nil {
		return 0
	}
	val  :=  n.data
	if n.left != nil{
		val = val + n.left.getSum()
	}
	if n.right != nil {
		val = val + n.right.getSum()
	}

	return val
}

func (n *NodeT) insert(data int) {

	if n == nil {
		n = &NodeT{data: data}
		return
	}

	if n.data == data {
		n.count ++
		fmt.Println("Dimdnot insert", n.data, n.count)
		return
	}

	if n.data > data {
		if n.left == nil {
			n.left = &NodeT{data: data}
			fmt.Println("Inserted at left", n.left.data)
		} else {
			//n = n.left
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &NodeT{data: data}
			fmt.Println("Inserted at right", n.right.data)
		} else {
			n.right.insert(data)
		}
	}
}



func (n *NodeT) Display() {
	if n == nil {
		return
	}

	fmt.Println(" ", n.data, n.count)
	n.left.Display()
	n.right.Display()
}
