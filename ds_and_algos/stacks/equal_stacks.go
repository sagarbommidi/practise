package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	Top *Node
	Bottom *Node
	Sum int
	Size int
}

func (s *Stack) IsEmpty() bool{
	return s.Top == nil
}

func (s *Stack) Print() {
	current := s.Top
	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}
	fmt.Println(" NILL, SUM: ", s.Sum)
}

func (s *Stack) Push(data int) {
	node := Node{data, nil}
	s.Size += 1
	s.Sum += data
	if(s.IsEmpty()) {
		s.Top = &node
		s.Bottom = &node
		return
	}
	s.Bottom.Next = &node
	s.Bottom = &node
}

func (s *Stack) Pop() int {
	if(s.IsEmpty()){
		return -1
	}
	temp := s.Top
	s.Top = temp.Next
	temp.Next = nil
	s.Size -= 1
	s.Sum -= temp.Data
	return temp.Data
}


func main() {
	fmt.Println("Implementing equal stacks")
	var l1, l2, l3 int
	_, _ = fmt.Scan(&l1, &l2, &l3)

	s1 := Stack{}
	for x := 0; x < l1; x++ {
		var val int
		_, _ = fmt.Scan(&val)
		s1.Push(val)
	}

	s2 := Stack{}
	for x := 0; x < l2; x++ {
		var val int
		_, _ = fmt.Scan(&val)
		s2.Push(val)
	}

	s3 := Stack{}
	for x := 0; x < l3; x++ {
		var val int
		_, _ = fmt.Scan(&val)
		s3.Push(val)
	}

	s1.Print()
	s2.Print()
	s3.Print()

	for true {
		if(s1.Sum == s2.Sum && s1.Sum == s3.Sum) {
			fmt.Println(s1.Sum)
			break
		}
		if(s1.Sum > s2.Sum) {
			s1.Pop()
		} else if (s1.Sum < s2.Sum) {
			s2.Pop()
		}
		if(s1.Sum > s3.Sum) {
			s1.Pop()
		} else if (s1.Sum < s3.Sum) {
			s3.Pop()
		}
		fmt.Println("Sum of stacks: ", s1.Sum, ", ", s2.Sum, ", ", s3.Sum)
	}
}