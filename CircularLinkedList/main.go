package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type CircularLinkedList struct {
	last   *node
	length int
}

func (c *CircularLinkedList) Length() int {
	return c.length
}

func (c *CircularLinkedList) IsEmpty() bool {
	return c.length == 0
}

func (c *CircularLinkedList) Display() {

	var first *node
	if c.last == nil {
		return
	}
	first = c.last.next

	for first != c.last {
		fmt.Println(first.data)
		first = first.next
	}
	fmt.Println(first.data)
}

func (c *CircularLinkedList) Prepend(data int) {
	temp := &node{data: data}

	if c.last == nil {
		c.last = temp
	} else {
		temp.next = c.last.next
	}
	c.last.next = temp
	c.length++
}

func (c *CircularLinkedList) insertLast(data int) {
	temp := &node{data: data}

	if c.last == nil {
		c.last = temp
		c.last.next = c.last
	} else {
		temp.next = c.last.next
		c.last.next = temp
		c.last = temp
	}
	c.length++
}

func (c *CircularLinkedList) RemoveFirst() int {
	var temp *node
	var result int
	if c.IsEmpty() {
		fmt.Println("Circular singly linked list is already empty")
	}

	temp = c.last.next
	result = temp.data

	if c.last.next == c.last {
		c.last = nil
	} else {
		c.last.next = temp.next
	}
	c.length--
	return result
}

func main() {
	c := CircularLinkedList{}

	c.Prepend(1)
	c.Prepend(2)
	c.Prepend(3)

	c.insertLast(4)
	c.insertLast(5)
	c.insertLast(6)

	c.RemoveFirst()

	c.Display()
}
