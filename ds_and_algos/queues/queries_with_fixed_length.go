package main
import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Queue struct {
	Front *Node
	Rear *Node
	Max *Node
	Size int
}

func (q *Queue) IsEmpty() bool{
	return q.Rear == nil
}

func (q *Queue) Push(data int) {
	node := Node{data, nil}
	q.Size += 1
	if(q.IsEmpty()) {
		q.Front = &node
		q.Rear = &node
		q.Max = &node
		return
	}
	q.Rear.Next = &node
	if(q.Max.Data < node.Data) {
		q.Max = &node
	}
	q.Rear = &node
}

func (q *Queue) Pop() int {
	if(q.IsEmpty()){
		return -1
	}
	temp := q.Front
	q.Front = temp.Next
	if(q.Front == nil) {
		q.Rear = nil
	}
	if(q.Max == temp) {
		q.SetMax()
	}
	q.Size -= 1
	return temp.Data
}

func (q *Queue) SetMax() {
	q.Max = q.Front
	current := q.Front
	for current != nil {
		if(q.Max.Data < current.Data) {
			q.Max = current
		}
		current = current.Next
	}
}

func (que *Queue) getMin(min int) int {
	if(min == -1) {
		min = que.Max.Data
	} else if(min > que.Max.Data) {
		min = que.Max.Data
	}
	return min
}

func getMinOfMax(arr []int, length int) int {
	que := Queue{}
	min := -1
	size := len(arr)
	for x := 0; x < size; x++ {
		que.Push(arr[x])
		if(que.Size < length) {
			continue
		}
		min = que.getMin(min)
		que.Pop()
	}
	return min
}


func main() {
	var count, queries, val int
	_, _ = fmt.Scan(&count, &queries)

	arr := make([]int, count)
	for x := 0; x < count; x++ {
		_, _ = fmt.Scan(&val)
		arr[x] = val
	}

	for y := 0; y < queries; y++ {
		_, _ = fmt.Scan(&val)
		fmt.Println(getMinOfMax(arr, val))
	}
}
