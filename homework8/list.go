package main

import (
	"errors"
	"fmt"
)

type listNode struct {
	value int
	next  *listNode
}

var head *listNode

//Add element to the list
func Add(head **listNode, value int) {
	if nil == *head {
		*head = new(listNode)
		(*head).value = value
		return
	}
	i := *head
	for ; i.next != nil; i = i.next {
		// todo   handle cycle

	}
	i.next = new(listNode)
	i.next.value = value
}

//Print the list's values
func Print() {

	i := head
	fmt.Println("The list elements are:")
	for ; i != nil; i = i.next {
		fmt.Println(i.value)

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

//Insert -Add the value at the given position
func Insert(value int, position uint) {
	length := Length()
	if position < 0 || position > length {
		fmt.Println("Incorrect value of position")
		return
	}
	node := new(listNode)
	node.value = value
	if position == 0 {
		node.next = head
		head = node
	} else {
		current := head
		var prev *listNode = nil
		var index uint = 0
		for index < position {
			prev = current
			current = current.next
			index++
		}
		prev.next = node
		node.next = current

	}
	fmt.Println(value, "was  added  at the given position")
	//length++

}

//Delete -delete the item at given position or return error if the position is outside the list boundaries
func Delete(position uint) error {
	length := Length()
	if position < 0 || position >= length {
		return errors.New("The index is outside the list boundaries")
	}
	current := head
	if position == 0 {
		head = current.next
	} else {
		var prev *listNode = nil
		var index uint = 0
		for index < position {
			prev = current
			current = current.next
			index++
		}
		prev.next = current.next

	}
	fmt.Println(current.value, "was deleted")
	return nil
}

//At -return the value at given position, or return error if the index is outside the list boundaries
func At(position uint) (int, error) {

	length := Length()
	if position < 0 || position >= length {
		return 0, errors.New("The index is outside the list boundaries")
	}

	current := head
	var index uint = 0
	for index < position {
		current = current.next
		index++
	}
	return current.value, nil
}

func main() {

	Add(&head, 1)
	Add(&head, 56)
	Add(&head, 13)
	Print()
	Insert(7, 0)
	Print()
	fmt.Println(Length())
	Delete(2)
	Print()
	fmt.Println(At(2))

}
