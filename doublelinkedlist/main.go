package main

import "fmt"

type Node struct {
	Value    string
	Previous *Node
	Next     *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

func main() {
	dll := &DoubleLinkedList{}
	node1 := &Node{Value: "first", Previous: nil, Next: nil}
	node2 := &Node{Value: "secod", Previous: nil, Next: nil}
	node3 := &Node{Value: "third", Previous: nil, Next: nil}

	node1.Next = node2
	node2.Previous = node1
	node2.Next = node3
	node3.Previous = node2

	dll.Head = node1
	dll.Tail = node3

	current := dll.Head

	for current != nil {
		fmt.Print(current.Value, "<-")
		current = current.Next
	}

}
