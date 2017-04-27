package main

import (
	"fmt"
)

func IsPrime(num int) bool {
	if(num < 2) {
		return false
	}
	if(num == 2 ) {
		return true
	}
	for x := 2; x < num/2+1; x++ {
		if(num % x == 0) {
			return false
		}
	}
	return true
}

func GetPrimes(count int) []int {
	primes := make([]int, 0)
	found := false
	for count > 0 {
		for x := 2; ; x++ {
			if(IsPrime(x) == true) {
				primes = append(primes, x)
			}
			if(len(primes) == count) {
				found = true
				break
			}
		}
		if found {
			break
		}

	}
	return primes
}

type Node struct {
	Data int
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
	node := Node{data, nil}
	s.Size += 1
	if(s.IsEmpty()) {
		s.Top = &node
		return
	}
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

func (s *Stack) Print() {
	for s.Top != nil {
		fmt.Println(s.Pop())
	}
}

func main() {
	fmt.Println("Implementing Waiter problem from stocks")
	var aCount, pCount, data int
	_, _ = fmt.Scan(&aCount, &pCount)

	primes := GetPrimes(pCount)
	fmt.Println(primes)

	initialStack := Stack{}
	finalStack := Stack{}
	resultStack := Stack{}

	for x := 0; x < aCount; x++ {
		var val int
		_, _ = fmt.Scan(&val)
		initialStack.Push(val)
	}

	for i := 0; i < pCount; i++ {
		for !initialStack.IsEmpty() {
			data = initialStack.Pop()
			if(data % primes[i] == 0) {
				resultStack.Push(data)
			} else {
				finalStack.Push(data)
			}
		}
		resultStack.Print()
		initialStack = finalStack
		finalStack = Stack{}
	}
	initialStack.Print()
}
