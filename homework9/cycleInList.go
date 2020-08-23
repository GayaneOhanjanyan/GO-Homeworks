package main

import (
	"fmt"
)

// Node of list
type Node struct {
	value int
	next  *Node
}

var head *Node

//Add element to the list
func Add(value int) {
	node := new(Node)
	node.value = value
	if 0 == Length() {
		head = node
	} else {

		i := head
		for ; i.next != nil; i = i.next {
			// todo   handle cycle

		}
		i.next = new(Node)
		i.next.value = value
	}
}

//Length -returns the list item count
func Length() uint {
	var count uint
	i := head
	for ; i != nil; i = i.next {
		count++
	}

	return count
}

//CreatCycle ....
func CreatCycle(head **Node) {
	i := *head
	for ; i.next != nil; i = i.next {
	}
	i.next = (*head).next

}

//DetectCycle -detect if there is any cycle
func (head *Node) DetectCycle() bool {
	point1 := head
	point2 := head
	for point2 != nil && point2.next != nil {
		point1 = point1.next
		point2 = point2.next.next
		if point1 == point2 {
			fmt.Println("Cycle found")
			return true

		}

	}
	fmt.Println("No cycle found")
	return false
}
func main() {
	Add(7)
	Add(11)
	Add(13)
	Add(8)
	CreatCycle(&head)
	head.DetectCycle()
}
