package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

type BST struct {
	count int
	root  *node
}

func insert(root *node, data int) *node {
	if root == nil {
		return New(data)
	} else if data < root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}

	return root
}

func New(data int) *node {
	n := node{
		data: data,
	}
	return &n
}

func main() {
	var root *node

	root = insert(root, 10)
	root = insert(root, 8)
	root = insert(root, 12)
	root = insert(root, 24)
	root = insert(root, 7)

	fmt.Println(root.data == 10)
	fmt.Println(root.left.data == 8)
	fmt.Println(root.right.data == 12)
	fmt.Println(root.right.right.data == 24)
	fmt.Println(root.left.left.data == 7)
}
