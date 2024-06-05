package main

import (
	"errors"
	"fmt"
)

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
	nwh.next = s.head
	s.head = nwh
	s.length++
}

func (s *stack) pop() (any, error) {
	if s.length == 0 {
		return nil, errors.New("stack underflow")
	}
	item := s.head.value
	s.head = s.head.next
	s.length--
	return item, nil
}

func (s *stack) peek() (any, error) {
	if s.length == 0 {
		return nil, errors.New("stack underflow")
	}
	return s.head.value, nil
}

func (s *stack) getAllData() ([]any, error) {
	if s.length == 0 {
		return nil, errors.New("stack underflow")
	}

	d := make([]any, 0)
	si := s.head
	for i := 0; i < s.length; i++ {
		d = append(d, si.value)
		si = si.next
	}
	return d, nil
}

func main() {
	s := new(stack)
	s.push(3)
	s.push(2)
	s.push(1)

	fmt.Println(s.peek()) // Access top element without modifying stack // 1, nil

	data, err := s.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)           // 1
	fmt.Println(s.getAllData()) // [2 3]
}
