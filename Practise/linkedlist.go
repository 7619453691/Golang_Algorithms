package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (L *LinkedList) Length() int {
	if L.head == nil {
		return 0
	}
	count := 0
	current := L.head
	for current != nil {
		current = current.next
		count++
	}
	return count
}

func (L *LinkedList) display() {
	current := L.head
	for current != nil {
		fmt.Print(current.data, "->")
		current = current.next
	}
	fmt.Println("nil")
}

func (L *LinkedList) insertFirst(data int) {
	newNode := &node{data: data}

	newNode.next = L.head
	L.head = newNode
}

func (L *LinkedList) insertLast(data int) {
	newNode := &node{data: data}
	if L.head == nil {
		L.head = newNode
		return
	}
	current := L.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (L *LinkedList) insert(data int, position int) {
	newNode := &node{data: data}

	if position == 1 {
		newNode.next = L.head
		L.head = newNode
	} else {
		previous := L.head
		count := 1
		for count < position-1 {
			previous = previous.next
			count++
		}
		current := previous.next
		previous.next = newNode
		newNode.next = current
	}
}

func (L *LinkedList) deleteFirst() *node {
	if L.head == nil {
		return nil
	}
	temp := L.head
	L.head = L.head.next
	temp.next = nil

	return temp
}

// func (L *LinkedList) deleteLast() *node {
// 	if L.head == nil || L.head.next == nil {
// 		return L.head
// 	}
// 	current := L.head
// 	previous := int
// 	for nil != current.next {
// 		previous = current
// 		current = current.next
// 	}
// 	previous.next = nil
// 	return current
// }

func (L *LinkedList) delete(position int) {

	if position == 1 {
		L.head = L.head.next
	} else {
		previous := L.head
		count := 1

		for count < position-1 {
			previous = previous.next
			count++
		}
		current := previous.next
		previous.next = current.next
	}
}

func (L *LinkedList) find(head *node, searchKey int) bool {
	if L.head == nil {
		return false
	}
	current := L.head
	for current != nil {
		if current.data == searchKey {
			return true
		}
		current = current.next
	}
	return false
}

func main() {

	l := LinkedList{}
	// l := LinkedList{}

	// l.head = &node{data: 10}
	// second := &node{data: 1}
	// third := &node{data: 8}
	// fourth := &node{data: 11}

	// l.head.next = second
	// second.next = third
	// third.next = fourth

	// fmt.Println(l.Length())

	l.insertLast(5)
	l.insertLast(10)
	l.insertLast(15)
	l.insertLast(20)

	// l.insert(12, 3)

	// l.deleteFirst()

	// l.deleteLast()

	// l.delete(3)

	// head := &node{data: 10}
	// second := &node{data: 8}
	// third := &node{data: 1}
	// fourth := &node{data: 11}

	// head.next = second
	// second.next = third
	// third.next = fourth

	// l := LinkedList{}

	// l.display()

	// l.display()

	// if l.find(l.head, 12) {
	// 	fmt.Println("Key is found")
	// } else {
	// 	fmt.Println("Not found")
	// }

	l.display()
}
