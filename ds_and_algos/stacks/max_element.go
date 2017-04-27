package main

import (
	"fmt"
)

func max(a1, a2 int) int {
	if(a1 > a2) {
		return a1
	} else {
		return a2
	}
}

type Node struct {
	Data int
	MaxData int
	Next *Node
}

type Stack struct {
	Top *Node
	Size int
}

func (s *Stack) IsEmpty() bool{
	return s.Top == nil
}

func (s *Stack) Push(data int) {
	node := Node{data, data, nil}
	s.Size += 1
	if(s.IsEmpty()) {
		s.Top = &node
		return
	}
	node.MaxData = max(data, s.Top.MaxData)
	node.Next = s.Top
	s.Top = &node
}

func (s *Stack) Pop() int {
	if(s.IsEmpty()){
		return -1
	}
	temp := s.Top
	s.Top = temp.Next
	temp.Next = nil
	s.Size -= 1
	return temp.Data
}

func main() {
	fmt.Println("Implementing Stack max element")
	newStack := Stack{}
	var count int
	_, _ = fmt.Scan(&count)

	for x := 0; x < count; x++ {
		var operation int
		_, _  = fmt.Scan(&operation)

		switch operation {
		case 1:
			var val int
			_, _ = fmt.Scan(&val)
			newStack.Push(val)
		case 2:
			newStack.Pop()
		case 3:
			fmt.Println(newStack.Top.MaxData)
		}

	}
}
