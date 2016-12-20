// Implementation of default stack methods in GoLang
package main

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

func (n *Node) SetNext(next *Node) {
	n.Next = next
}

func (n *Node) setData(data interface{}) {
	n.Data = data
}

type Stack struct {
	Top *Node
	Size int
}

func (s *Stack) IsEmpty() bool{
	return s.Top == nil
}

func (s *Stack) Push(data interface{}) {
	node := Node{data, nil}
	s.Size += 1
	if(s.Top == nil) {
		s.Top = &node
	}
	node.Next = s.Top
	s.Top = &node
}

func (s *Stack) Pop() interface{} {
	if(s.Top == nil){
		return -1
	}
	temp := s.Top
	s.Top = temp.Next
	temp.Next = nil
	s.Size -= 1
	return temp.Data
}

func (s *Stack) PushAndPrint(data interface{}) {
	s.Push(data)
	fmt.Println("Size of the Stack is: ", s.Size)
	fmt.Println("Top element of the stack is: ", s.Top.Data)
}

func (s *Stack) PopAndPrint() {
	s.Pop();
	fmt.Println("Size of the Stack is: ", s.Size)
	fmt.Println("Top element of the stack is: ", s.Top.Data)
}

func main() {
	fmt.Println("Welcome to Stack Implementation")
	stac := Stack{}
	stac.PushAndPrint(5)
	stac.PushAndPrint(1)
	stac.PushAndPrint(6)
	stac.PopAndPrint()
	stac.PopAndPrint()
}
