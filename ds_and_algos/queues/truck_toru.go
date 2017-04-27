package main

import "fmt"

type Node struct {
	Petrol int
	Distance int
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

func (q *Queue) Push(petrol, dist int) {
	node := Node{petrol, dist, nil}
	q.Size += 1
	if(q.IsEmpty()) {
		q.Front = &node
		q.Rear = &node
		return
	}
	q.Rear.Next = &node
	q.Rear = &node
}

func (q *Queue) PopAndPush() int {
	if(q.IsEmpty()){
		return -1
	}
	temp := q.Front
	q.Front = temp.Next
	q.Rear.Next = temp
	q.Rear = temp
	temp.Next = &Node{}
	if(q.Front == nil) {
		q.Rear = nil
	}
	return temp.Petrol
}

func (q *Queue) CanReachEnd() bool {
	current := q.Front
	remainingPetrol := 0

	for current != nil {
		remainingPetrol = (remainingPetrol + current.Petrol - current.Distance)
		if (remainingPetrol < 0) {
			return false
		}
		current = current.Next
	}
	return true
}

func (q *Queue) TruckTour(size int) int {
	for x := 0; x < size; x++ {
		if(q.CanReachEnd()) {
			return x
		} else {
			q.PopAndPush()
		}
	}
	return -1
}

func main() {
	var count, petrol, dist int
	_, _ = fmt.Scan(&count)

	que := Queue{}
	for x := 0; x < count; x++ {
		_, _ = fmt.Scan(&petrol, &dist)
		que.Push(petrol, dist)
	}
	fmt.Println(que.TruckTour(count))
}
