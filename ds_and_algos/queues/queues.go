package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type Queue struct {
	Front *Node
	Rear *Node
	Size int
}

func (q *Queue) IsEmpty() bool{
	return q.Rear == nil
}

func (q *Queue) Push(data int) {
	node := Node{data, nil, nil}
	q.Size += 1
	if(q.IsEmpty()) {
		q.Front = &node
		q.Rear = &node
		return
	}
	q.Rear.Next = node
	q.Rear = &node
}

func (q *Queue) Pop() int {
	if(q.IsEmpty()){
		return -1
	}
	q.Front = q.Front.Next
	if(q.Front == nil) {
		q.Rear = nil
	}
	q.Size -= 1
	return q.Front.Data
}

func (q *Queue) Print() {
	current := q.Front
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func main() {
	fmt.Print("Sagar Bommidi")
}

