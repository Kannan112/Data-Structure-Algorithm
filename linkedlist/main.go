package main

import (
	"fmt"
)

type Node struct {
	value string
	next  *Node
}

type LikedList struct {
	head *Node
}

func main() {

	list := LikedList{}
	list.add("joseph")
	list.add("alan")
	list.add("royal")
	list.add("abhinadn")
	list.add("successfull")
	list.Delete("alan")
	list.Printlist()

}

func (l *LikedList) add(str string) {
	newnode := &Node{value: str, next: nil}
	if l.head == nil {
		l.head = newnode
		return
	}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newnode
}

func (l *LikedList) Printlist() {
	current := l.head
	for current != nil {
		fmt.Println(current, " ")
		current = current.next
	}
	fmt.Println()
}

func (l *LikedList) Delete(name string) {
	if l.head == nil {
		fmt.Println("add some values")
		return
	}
	list := l.head
	if list.next.value == name {
		list.next = list.next.next
	}

	for list.next != nil {
		if list.next.value == name {
			list.next = list.next.next
			return
		}
		if list.value != name {
			fmt.Println("empty")
			return
		}
		list = list.next
	}

}
