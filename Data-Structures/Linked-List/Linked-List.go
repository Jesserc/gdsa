package main

import (
	"fmt"
	"log"
)

type node struct {
	value any
	next  *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) error {
	if n.next != nil {
		return fmt.Errorf("n.next should be 'nil' but got [%v]", n.next)
	}

	prevHead := l.head
	l.head = n             // update linked list head to the new node
	l.head.next = prevHead // set next node as the previous head
	l.length++             // increment linked list length

	return nil
}

func (l *LinkedList) deleteItem(value any) {
	if l.length == 0 {
		return
	}

	if l.head.value == value { // handle case where data belongs to the first node (head)
		l.head = l.head.next
		l.length--
	}

	prevNode := l.head
	for i := 0; i < l.length; i++ {
		if prevNode.next != nil { // we do this to handle the last node (tail)

			/*
				// this is bad because we will be deleting the next node and not the current one
				if prevNode.value == value {
					prevNode.next = prevNode.next.next
				}
			*/

			// The issue from above is handled here by checking the next value
			// The only downside is that the 'head' node won't be checked, but we're handling this above already
			if prevNode.next.value == value {
				prevNode.next = prevNode.next.next // skip current.next, effectively deleting it
				l.length--                         // reduce length to indicate an item was deleted
			}
		}

		prevNode = prevNode.next
	}
}
func (l *LinkedList) getAllData() []any {
	if l.length == 0 {
		return nil
	}

	d := make([]any, 0)
	currentL := l.head
	for i := 0; i < l.length; i++ {
		d = append(d, currentL.value)
		currentL = currentL.next
	}
	return d
}

func New(value any) *LinkedList {
	return &LinkedList{head: &node{value: value}, length: 1}
}

func main() {
	l := New("Item 1")

	err := l.prepend(&node{value: "Item 2"})
	if err != nil {
		log.Fatal(err)
	}

	err = l.prepend(&node{value: "Item 3"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(l.getAllData())

	l.deleteItem("Item 3")
	fmt.Println(l.getAllData())

	l.deleteItem("Item 2")
	fmt.Println(l.getAllData())
}
