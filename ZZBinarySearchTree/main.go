package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

type binarySearchTree struct {
	root *node
}

func (b binarySearchTree) inOrder(root *node) {
	if b.root == nil {
		return
	}
	b.inOrder(b.root.left)
	fmt.Print(b.root.data, " ")
	b.inOrder(b.root.right)
}

func (b binarySearchTree) insert(root *node, val int) *node {
	if b.root == nil {
		b.root = &node{data: val}
		return b.root
	}
	if b.root.data > val {
		b.root.left = b.insert(b.root.left, val)
	} else {
		b.root.right = b.insert(b.root.right, val)
	}
	return b.root
}

func (b binarySearchTree) search(root *node, key int) bool {
	if b.root == nil {
		return false
	}
	if b.root.data > key {
		return b.search(b.root.left, key)
	} else if b.root.data == key {
		return true
	} else {
		return b.search(b.root.right, key)
	}
}

func (b binarySearchTree) delete(root *node, val int) *node {
	if b.root.data > val {
		b.root.left = b.delete(b.root.left, val)
	} else if b.root.data < val {
		b.root.right = b.delete(b.root.right, val)
	} else {
		if b.root.left == nil && b.root.right == nil {
			return nil
		}
		if b.root.left == nil {
			return b.root.right
		} else if b.root.right == nil {
			return b.root.left
		}

		is := b.inOrderSuccessor(b.root.right)
		b.root.data = is.data
		b.root.right = b.delete(b.root.right, is.data)
	}
	return b.root
}

func (b binarySearchTree) inOrderSuccessor(root *node) *node {
	for b.root.left != nil {
		b.root = b.root.left
	}
	return b.root
}

func (b binarySearchTree) printInRange(root *node, x int, y int) {
	if b.root == nil {
		return
	}
	if b.root.data >= x && b.root.data <= y {
		b.printInRange(b.root.left, x, y)
		fmt.Print(b.root.data, " ")
		b.printInRange(b.root.right, x, y)
	} else if b.root.data >= y {
		b.printInRange(b.root.left, x, y)
	} else {
		b.printInRange(b.root.right, x, y)
	}
}

func main() {
	b := binarySearchTree{}

	values := [9]int{8, 5, 3, 1, 4, 6, 10, 11, 14}

	b.root = nil

	for i := 0; i < len(values); i++ {
		b.root = b.insert(b.root, values[i])
	}

	b.inOrder(b.root)
	fmt.Println()

	if b.search(b.root, 10) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}

	// b.delete(b.root, 10)
	// b.inOrder(b.root)

}
