package main

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (B BinarySearchTree) insert(value int) {
	B.root = B.insert(B.root, value)
}

func (B BinarySearchTree) insert(root *TreeNode, value int) *TreeNode {
	if B.root == nil {
		B.root := &TreeNode{value: value}
		return B.root
	}
	if value < B.root.data {
		B.root.left = B.insert(B.root.left, value)
	} else {
		B.root.right = B.insert(B.root.right, value)
	}
}

func (B BinarySearchTree) inOrder() {
	B.inOrder(B.root)
}

func (B BinarySearchTree) inOrder(root *TreeNode) {
	if B.root == nil {
		return
	}
	B.inOrder(B.root.left)
	fmt.Print(B.root.data, " ")
	B.inOrder(B.root.right)
}

func (B BinarySearchTree) search(key int) {
	return B.search(B.root, key)
}
func (B BinarySearchTree) search(root *TreeNode, key int) *TreeNode {
	if B.root == nil || B.root.data == key {
		return B.root
	}
	if key < B.root.data {
		return B.search(B.root.left, key)
	} else {
		return B.search(B.root.right, key)
	}
}
func main() {
	bst := BinarySearchTree{}

	bst.insert(5)
	bst.insert(3)
	bst.insert(7)
	bst.insert(1)

	bst.inOrder()

	fmt.Println()

	if nil != bst.search(3) {
		fmt.Println("Key is found")
	} else {
		fmt.Println("Not found")
	}
}
