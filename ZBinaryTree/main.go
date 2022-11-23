package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

type binaryTree struct {
	root *node
}

func (b binaryTree) buildTree(nodes []int) *node {
	idx := -1
	idx++
	if nodes[idx] == -1 {
		return nil
	}
	newNode := &node{data: nodes[idx]}
	newNode.left = b.buildTree(nodes)
	newNode.right = b.buildTree(nodes)

	return newNode
}

func (b binaryTree) preOrder(root *node) {
	if b.root == nil {
		return
	}
	fmt.Print(root.data, " ")
	b.preOrder(b.root.left)
	b.preOrder(b.root.right)
}

// func (b binaryTree) inOrder(root *node) {
// 	if b.root == nil {
// 		return
// 	}
// 	b.inOrder(b.root.left)
// 	fmt.Print(root.data, " ")
// 	b.inOrder(b.root.right)
// }

// func (b binaryTree) postOrder(root *node) {
// 	if b.root == nil {
// 		return
// 	}
// 	b.inOrder(b.root.left)
// 	b.inOrder(b.root.right)
// 	fmt.Print(root.data, " ")
// }

func main() {
	nodes := []int{1, 2, 4, -1, -1, 5, -1, -1, 3, -1, 6, -1, -1}
	bst := binaryTree{}
	bst.root = bst.buildTree(nodes)

	// fmt.Print(bst.root.data, " ")

	bst.preOrder(bst.root)
	// bst.postOrder(bst.root)
	// bst.inOrder(bst.root)
}
