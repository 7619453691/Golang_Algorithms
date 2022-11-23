package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

type BinarySearchTree struct {
	root   *node
	length int
}

func (b BinarySearchTree) createNode(data int) *node {
	var n *node
	n.data = data
	n.left = nil
	n.right = nil

	return n
}

func (b BinarySearchTree) preOrderTraversal(root *node) {
	if b.root != nil {
		fmt.Printf("%d ", b.root.data)
		b.preOrderTraversal(b.root.left)
		b.preOrderTraversal(b.root.right)
	}
}

func (b BinarySearchTree) postOrderTraversal(root *node) {
	if b.root != nil {
		b.postOrderTraversal(b.root.left)
		b.postOrderTraversal(b.root.right)
		fmt.Printf("%d ", b.root.data)
	}
}

func (b BinarySearchTree) inOrderTraversal(root *node) {
	if b.root != nil {
		b.inOrderTraversal(b.root.left)
		fmt.Printf("%d ", b.root.data)
		b.inOrderTraversal(b.root.right)
	}
}

// func (b BinarySearchTree) isBST(root *node) int {
// 	prev := new(node)
// 	prev.data = nil
// 	if b.root != nil {

// 		if !b.isBST(b.root.left)) {
// 			return 0
// 		}
// 		if prev != nil && b.root.data <= prev.data {
// 			return 0
// 		}
// 		prev = b.root
// 		return b.isBST(b.root.right)
// 	} else {
// 		return 1
// 	}
// }

// func (b BinarySearchTree) search(root *node, key int) {
// 	if b.root != nil {
// 		return
// 	}
// 	if key == b.root.data {
// 		return 
// 	} else if key < b.root.data {
// 		return b.search(b.root.left, key)
// 	} else {
// 		return b.search(b.root.right, key)
// 	}
// } 

func (b BinarySearchTree) searchItr(root *node, key int)  {
	for b.root != nil {

		if key == b.root.data {
			return b.root
		}
		else if key < b.root.data {
			b.root = b.root.left 
		} else {
			b.root = b.root.right
		}
	}
	return nil
} 

func (b BinarySearchTree) insert(root *node, key int) {
	var prev *node
	prev = nil

	for b.root != nil {
		prev = b.root

		if key == b.root.data {
			fmt.Println("Cannot insert %d, alraedy in BST", key)
			return
		}
		else if key < b.root.data {
			b.root = b.root.left
		} else {
			b.root = b.root.right
		}
	} 
	new := b.createNode(key)
	if key < prev.data {
		prev.left = new
	} else {
		prev.right = new
	}
}



func main() {

	b := BinarySearchTree{}

	p := b.createNode(5)
	p1 := b.createNode(3)
	p2 := b.createNode(6)
	p3 := b.createNode(1)
	p4 := b.createNode(4)

	p.left = p1
	p.right = p2
	p1.left = p3
	p1.right = p4

	b.preOrderTraversal(p)
}
