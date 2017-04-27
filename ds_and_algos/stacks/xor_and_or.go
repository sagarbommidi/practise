package main

import "fmt"

func max(a1, a2 int) int {
	if a1 > a2 {
		return a1
	} else {
		return a2
	}
}

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	Top  *Node
	Size int
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

func (s *Stack) Push(data int) {
	node := Node{data, nil}
	s.Size += 1
	if s.IsEmpty() {
		s.Top = &node
		return
	}
	node.Next = s.Top
	s.Top = &node
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	temp := s.Top
	s.Top = temp.Next
	temp.Next = nil
	s.Size -= 1
	return temp.Data
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.Top.Data
}

func main() {
	var count, peek int
	_, _ = fmt.Scan(&count)

	maxVal := -1
	minStack := Stack{}
	for x := 0; x < count; x++ {
		var val int
		_, _ = fmt.Scan(&val)
		if minStack.Size == 0 {
			minStack.Push(val)
			continue
		}
		for !minStack.IsEmpty() {
			peek = minStack.Peek()
			maxVal = max(maxVal, (peek ^ val))
			if peek > val {
				minStack.Pop()
			} else {
				break
			}
		}
		minStack.Push(val)
	}

	top := minStack.Pop()
	for !minStack.IsEmpty() {
		temp := minStack.Pop()
		maxVal = max(maxVal, (top ^ temp))
		top = temp
	}
	fmt.Println(maxVal)
}
