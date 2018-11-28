package main

import (
	"fmt"
	"bytes"
	"errors"
)

type Node struct {
	Left, Right *Node
	Data        int
}

func main() {
	data := []int{17, 14, 7, 2, 3, 19, 1,99}
	fmt.Println("hello world")
	var root = &Node{Data: data[0]}
	fmt.Println("root", root)

	for _, v := range data[1:] {
		root.Insert(v)
	}
	root.Display()
	fmt.Println("\nHeight Left:", root.Left.Data)
	fmt.Println("Height Right:", root.Right.Data)

	fmt.Println("\nHeight Left:", root.Left.Height())
	fmt.Println("Height Right:", root.Right.Height())
	//fmt.Println(root)


	W := new(bytes.Buffer)
	root.printTree(W)

	var (
		err error
		n   int
	)
	data1 := make([]byte, 1024)
	n, err = W.Read(data1)
	fmt.Println("The error：", err)
	fmt.Printf("%s\n", data1[:n])

	//for n, err = R.Read(data1); err == nil; n, err = R.Read(data1) {
	//
	//}
	//fmt.Printf("The data：%#v\n", data1[:n])

}

func (node *Node) Insert(data int) (error) {
	if node == nil {
		node = &Node{Data: data}
		return errors.New("Root Node")
		//return errors.New("Root is empty")
	}

	switch {
	case node.Data < data:
		if node.Left == nil {
			node.Left = &Node{Data: data}
			fmt.Println("",node.Data , "->Left" , data )
			return nil
		}
		node.Left.Insert(data)
	case node.Data > data:
		if node.Right == nil {
			node.Right = &Node{Data: data}
			fmt.Println("",node.Data , "->Right" , data )
			return nil
		}
		node.Right.Insert(data)
	}
	return nil
}

//func (node *Node) String() string {
//	if node != nil {
//
//	}
//	return "madhu"
//}

func (node *Node) Display() {
	if node == nil {
		return
	}

	node.Left.Display()
	fmt.Print(node.Data,",")
	node.Right.Display()
}

func (node *Node) printTree(out *bytes.Buffer) {
	if (node.Left != nil) {
		node.Left.printTree1(out, false, "");
	}
	node.printNodeValue(out);
	if (node.Right != nil) {
		node.Right.printTree1(out, true, "");
	}
}

func (node *Node) printNodeValue(out *bytes.Buffer) {
	out.Write([]byte(fmt.Sprintf("%d", node.Data)));
	out.Write([]byte("\n"));
}

// use string and not stringbuffer on purpose as we need to change the indent at each recursion
func (node *Node) printTree1(out *bytes.Buffer, isRight bool, indent string) {

	if (node.Left != nil) {
		str := "      "
		if isRight {
			str = " |    "
		}
		str = indent + str
		node.Left.printTree1(out, false, str)
	}

	out.Write([]byte(indent))
	if (isRight) {
		out.Write([]byte("\\"))
	} else {
		out.Write([]byte ("/"))
	}
	out.Write([]byte("--"))

	node.printNodeValue(out)

	if (node.Right != nil) {
		str := " |    "
		if isRight {
			str = "      "
		}
		str = indent + str
		node.Right.printTree1(out, true, str)
	}

}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}

	return 1 + max(node.Left.Height(), node.Right.Height())
}

func max(l, r int) int {
	if l < r {
		return r
	}

	return l
}
