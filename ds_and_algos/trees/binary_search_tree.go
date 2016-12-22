package main

import (
        "fmt"
)

type Node struct {
        Data int
        Left *Node
        Right *Node
}

type BST struct {
        Root *Node
}

func (t *BST) Search(key int)  {

}

/*
        Implementation With recursion function
 */

// InsertNode: Insert given key in the BST
func (node *Node)InsertNode(key int) *Node{
        if(node == nil) {
                return &Node{key, nil, nil}
        }
        if(key < node.Data) {
                node.Left = node.Left.InsertNode(key)
        }else if(key > node.Data) {
                node.Right = node.Right.InsertNode(key)
        }
        return node
}

func (t *BST)Insert(key int) {
        t.Root = t.Root.InsertNode(key)
}

// DeleteNode: deletes the given key from BST
func (node *Node)DeleteNode() {
}

func (s *BST)Delete(key int) {

}

// InorderTraverseNode: Inorder Traversal of BST
func (node *Node)InorderTraverseNode() {
        if(node != nil) {
                node.Left.InorderTraverseNode()
                fmt.Print(node.Data, " --> ")
                node.Right.InorderTraverseNode()
        }
}

func (t *BST) InorderTraverse() {
	t.Root.InorderTraverseNode()
}

// PreorderTraversalNode: Preorder Traversal of BST
func (node *Node) PreorderTraversalNode() {
	if(node != nil) {
		fmt.Print(node.Data, " --> ")
		node.Left.PreorderTraversalNode()
		node.Right.PreorderTraversalNode()
	}
}

func (t *BST) PreorderTraverse() {
	t.Root.PreorderTraversalNode()
}

// PostorderTraversalNode: Postorder  Traversal of BST
func (node *Node) PostorderTraversalNode() {
	if(node != nil) {
		node.Left.PreorderTraversalNode()
		node.Right.PreorderTraversalNode()
		fmt.Print(node.Data, " --> ")
	}
}

func (t *BST) PostorderTraverse() {
	t.Root.PostorderTraversalNode()
}

// Minvalue: Find the minimum value of BST
func (s *BST) Minvalue() int{
        if(s.Root == nil) {
                return -1
        }
        current := s.Root
        for (current.Left != nil) {
                current = current.Left
        }
        return current.Data

}

// Maxvalue: Find the max value of BST
func (s *BST) Maxvalue() int {
        if(s.Root == nil){
                return -1
        }
        current := s.Root
        for (current.Right != nil) {
                current = current.Right
        }
        return current.Data
}

/*
        Implementation without using Recursion
 */

func (node *Node)InsertNodeWithoutRec(key int) *Node{
        new_node := Node{key, nil, nil}
        if(node == nil) {
                return &new_node
        }
        current := node
        for (current != nil) {
                if(key < current.Data) {
                        current = current.Left
                }else if (key > current.Data) {
                        current = current.Right
                }
        }
        current = &new_node
        return current
}

func main() {
        fmt.Println("BST Implementation.")
        mybst := BST{}
        mybst.Insert(5)
        mybst.Insert(10)
        mybst.Insert(20)
        mybst.Insert(30)
        mybst.Insert(0)
        mybst.InorderTraverse()
	mybst.PreorderTraverse()
}

