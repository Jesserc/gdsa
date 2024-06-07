package main

import "fmt"

type node struct {
	value    any
	previous *node
	next     *node
}

type DoublyLinkedList struct {
	head   *node
	length int
}

func (l *DoublyLinkedList) append(n *node) error {
	if n.next != nil {
		return fmt.Errorf("n.next should be 'nil' but got [%v]", n.next)
	}

	prevNode := l.head
	for i := 0; i < l.length; i++ {
		if prevNode.next == nil {
			n.previous = prevNode
			prevNode.next = n
			l.length++
			return nil
		}
		prevNode = prevNode.next
	}

	return nil
}

func (l *DoublyLinkedList) prepend(n *node) error {
	if n.next != nil {
		return fmt.Errorf("n.next should be 'nil' but got [%v]", n.next)
	}

	prevHead := l.head // cache the current head node
	l.head.previous = n
	l.head = n             // Set `n` as the new head node
	l.head.next = prevHead // make `n` (the new head) point to the previous head node
	l.length++

	return nil
}

func (l *DoublyLinkedList) insert(index int, n *node) {

	prevNode := l.head
	for i := 0; i < l.length; i++ {
		if i == index-1 {
			// target is the node currently at the position where we want to insert the new node `n`
			target := prevNode.next

			// Set the target node's previous pointer to `n`, since `n` will be inserted right before the target
			target.previous = n

			// Set the new node's next pointer to the target node, making `n` point to the node that follows it
			n.next = target

			// Set the previous node's next pointer to `n`, so the node before the insertion point now points to `n` instead of `target`
			prevNode.next = n

			// Set the new node's previous pointer to the previous node, establishing a backward link
			n.previous = prevNode
		}
		// Move to the next node in the list on each iteration
		prevNode = prevNode.next
	}
}
func (l *DoublyLinkedList) getAllData() []any {
	d := make([]any, 0)

	prevNode := l.head
	for i := 0; i < l.length; i++ {
		d = append(d, prevNode.value)
		prevNode = prevNode.next
	}
	return d
}

func main() {
	l := DoublyLinkedList{
		head: &node{
			value:    "Hello world",
			previous: nil,
			next:     nil,
		},
		length: 1,
	}

	l.prepend(&node{value: "(Prepended: Hi world)"})
	l.append(&node{value: "(Appended Hey world)"})

	fmt.Println(l.getAllData())
}
