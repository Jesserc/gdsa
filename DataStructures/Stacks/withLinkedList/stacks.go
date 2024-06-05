package main

import "fmt"

type node struct {
	next  *node
	value any
}

type stack struct {
	head   *node
	length int
}

func (s *stack) push(item any) {
	nwh := &node{value: item}
	h := s.head
	nwh.next = h
	s.head = nwh
	s.length++
}

func (s *stack) pop() {
	s.head = s.head.next
	s.length--
}

func (s *stack) getAllData() []any {
	if s.length == 0 {
		return nil
	}

	d := make([]any, 0)
	si := s.head
	for i := 0; i < s.length; i++ {
		d = append(d, si.value)
		si = si.next
	}
	return d
}

func main() {
	s := new(stack)
	s.push(3)
	s.push(2)
	s.push(1)

	fmt.Println(s.getAllData())
	s.pop()
	fmt.Println(s.getAllData())
}
