package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var item T
	if len(s.items) == 0 {
		return item, errors.New("stack underflow")
	}
	l := len(s.items) - 1
	item = s.items[l]
	s.items = s.items[:l]
	return item, nil
}

func main() {
	s := &Stack[int]{}

	s.Push(3)
	s.Push(2)
	s.Push(1)
	item, err := s.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*s)   // [2 3]
	fmt.Println(item) // 1
}
