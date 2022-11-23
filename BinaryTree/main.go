package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

type BinaryTree struct {
	root *node
}

func (bst *BinaryTree) insert(value int) bool {
	newNode := &node{value: value}

	if bst.root == nil {
		bst.root = newNode
		return true
	}
	temp := bst.root
	for true {
		if newNode.value == temp.value {
			return false
		}
		if newNode.value < temp.value {
			if temp.left == nil {
				temp.left = newNode
				return true
			}
			temp = temp.left
		} else {
			if temp.right == nil {
				temp.right = newNode
				return true
			}
			temp = temp.right
		}
	}
	return false
}

func (bst *BinaryTree) contains(value int) bool {
	if bst.root == nil {
		return false
	}
	temp := bst.root
	for temp != nil {
		if value < temp.value {
			temp = temp.left
		} else if value > temp.value {
			temp = temp.right
		} else {
			return true
		}
	}
	return false
}

func main() {
	bst := BinaryTree{}

	bst.insert(47)
	bst.insert(21)
	bst.insert(76)
	bst.insert(18)
	bst.insert(52)
	bst.insert(82)

	bst.insert(27)

	fmt.Println(bst.root.left.right.value)

	fmt.Println(bst.contains(76))

}
