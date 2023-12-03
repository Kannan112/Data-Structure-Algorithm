package main

import (
	"fmt"
)

type Node struct {
	Value    string
	Previous *Node
	Next     *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

func (ddl *DoubleLinkedList) InsertionAtEnd(value string) {
	// node
	newNode := &Node{Value: value, Previous: nil, Next: nil}
	if ddl.Head == nil {
		ddl.Head = newNode
		ddl.Tail = newNode
		return
	}
	current := ddl.Tail
	fmt.Println("test1", current)
	current.Next = newNode
	fmt.Println("test2", current.Previous)
	current.Previous = current
	fmt.Println("test3", "currentNow", current, current.Previous)
	ddl.Tail = newNode
}

func (dll *DoubleLinkedList) InsertionAtBeg(value string) {
	// node
	newNode := &Node{Value: value, Previous: nil, Next: nil}

	if dll.Tail == nil {
		dll.Head = newNode
		dll.Tail = newNode
	}
	current := dll.Tail
	current.Previous = newNode
	current.Next = current

	//dll.Tail =

}

func (ddl *DoubleLinkedList) Display() {
	current := ddl.Head
	for current != nil {
		fmt.Print(current.Value, "<-")
		current = current.Next
	}
}

func main() {
	dll := &DoubleLinkedList{}
	// node1 := &Node{Value: "first", Previous: nil, Next: nil}
	// node2 := &Node{Value: "secod", Previous: nil, Next: nil}
	// node3 := &Node{Value: "third", Previous: nil, Next: nil}
	// node1.Next = node2
	// node2.Previous = node1
	// node2.Next = node3
	// node3.Previous = node2

	// dll.Head = node1
	// dll.Tail = node3

	dll.InsertionAtEnd("helo1")
	dll.InsertionAtEnd("helo2")
	dll.InsertionAtEnd("helo3")
	//dll.InsertionAtBeg("first1")
	dll.Display()

}
