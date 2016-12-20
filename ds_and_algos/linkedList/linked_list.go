// Implementation of default LinkedList methods in GoLang
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

type MyList struct {
	Head *Node
}

func (m *MyList) Add(data interface{}, index int) {
	newNode := Node{data, nil}
	if m.Head == nil {
		m.Head = &newNode
		return
	}
	if index == 0 {
		temp := m.Head
		m.Head = &newNode
		newNode.SetNext(temp)
		return
	}
	count := 1
	temp := m.Head
	for temp.Next != nil {
		if index == count {
			break
		}
		temp = temp.Next
		count += 1
	}
	tempNext := temp.Next
	temp.SetNext(&newNode)
	newNode.SetNext(tempNext)
}

func (m *MyList) Count() int {
	if m.Head == nil {
		return 0
	}
	count := 1
	temp := m.Head
	for temp.Next != nil {
		count += 1
		temp = temp.Next
	}
	return count
}

func (m *MyList) Get(pos int) *Node {
	// ToDo: Implement me
	return &Node{}
}

func (m *MyList) Reverse() {
	if m.Head == nil {
		return
	}
	prev := Node{}.Next
	current := m.Head
	for current != nil {
		next := current.Next
		current.SetNext(prev)
		prev = current
		current = next
	}
	m.Head = prev
}

func (m *MyList) Print() {
	current := m.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (m *MyList) Delete(pos int) {
	current := m.Head
	if current == nil {
		return
	}
	if(pos == 0) {
		m.Head = current.Next
		return
	}
	count := 0
	prev := current
	for (current.Next != nil) {
		prev = current
		current = current.Next
		count += 1
		if(count == pos) {
			break
		}
	}
	prev.Next = current.Next
	current.Next = nil

}

func main() {
	fmt.Println("Welcome to MyList Implementation")
	myl := MyList{}
	myl.Add(6, -1)
	myl.Add(10, -1)
	myl.Add(12, -1)
	myl.Add(14, -1)
	myl.Add(8, 1)
	myl.Add(4, 0)
	myl.Print()
	myl.Reverse()
	myl.Print()
	myl.Delete(2)
	myl.Print()
	myl.Delete(2)
	myl.Print()
}
