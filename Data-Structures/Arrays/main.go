package main

import (
	"cmp"
	"fmt"
)

// Array is a generic data structure that holds elements of a specific ordered type.
type Array[T cmp.Ordered] struct {
	content []T
}

// pop removes the last element from the array.
func (arr *Array[T]) pop() T {
	pop, newArr := arr.content[0], arr.content[:len(arr.content)-1]
	arr.content = newArr
	return pop
}

// push appends a new element to the end of the array.
func (arr *Array[T]) push(value T) {
	arr.content = append(arr.content, value)
}

// func (arr *Array[T]) insert(value T) {
// 	arr.content = append(arr.content, value)
// }

// unshift inserts one or more elements at the beginning of the array.
func (arr *Array[T]) unshift(value ...T) {
	newLen := len(arr.content) + len(value)
	newArr := make([]T, newLen)
	copy(newArr[:], value)
	copy(newArr[len(value):], arr.content)
	arr.content = newArr
}

// shift removes an element at the beginning of the array.
func (arr *Array[T]) shift() T {
	shift, newArr := arr.content[0], arr.content[1:]
	arr.content = newArr
	return shift
}

// delete removes an element at the specified index from the array.
func (arr *Array[T]) delete(index int) {
	for i := index; i < len(arr.content)-1; i++ {
		arr.content[i] = arr.content[i+1]
	}
	arr.pop() // Remove duplicated last element
}

// New creates a new Array with the specified capacity.
func New[T cmp.Ordered](len int) *Array[T] {
	arr := Array[T]{content: make([]T, 0, len)}
	return &arr
}

func main() {
	// Create a new Array
	arr := New[string](4)

	// Push elements into the array
	arr.push("apple")
	arr.push("big")
	arr.push("cat")
	arr.push("dog")

	// Unshift elements to the beginning of the array
	arr.unshift("changed 1", "changed 2")

	// Delete elements ("changed 1" & "changed 2") from the array
	arr.delete(0)
	arr.delete(0)

	// Print the current content of the array
	fmt.Println(arr.content)

	// Delete the first element of the array and print the current content of the array
	first := arr.shift()
	fmt.Printf("%v was deleted\n", first)
	fmt.Println(arr.content)

	// Delete the last element of the array and print the current content of the array
	last := arr.pop()
	fmt.Printf("%v was deleted\n", last)
	fmt.Println(arr.content)

	// Reverse a string using two different methods
	word := "Hello World!"
	fmt.Println(reverseString(word))    // !dlroW olleH
	fmt.Println(reverseStringTwo(word)) // !dlroW olleH

	// Merge two sorted slices
	merged := mergeSortedSlices([]int{1, 3, 5, 7, 20, 30, 40}, []int{2, 4, 6, 50})
	fmt.Println("Merged slice:", merged)
}
