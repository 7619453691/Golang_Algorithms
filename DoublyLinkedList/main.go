package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
	prev *node
}

type doublylinkedlist struct {
	head   *node
	tail   *node
	length int
}

func (d *doublylinkedlist) Length() int {
	return d.length
}

func (d *doublylinkedlist) printList() {
	temp := d.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func (d *doublylinkedlist) append(data int) {
	newNode := &node{data: data}
	if d.length == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	}
	d.length++
}

func (d *doublylinkedlist) removeLast() *node {
	if d.length == 0 {
		return nil
	}
	temp := d.tail

	if d.length == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.tail = d.tail.prev
		d.tail.next = nil
		temp.prev = nil
	}
	d.length--
	return temp
}

func (d *doublylinkedlist) prepend(data int) {
	newNode := &node{data: data}
	if d.length == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.length++
}

func (d *doublylinkedlist) removeFirst() *node {

	if d.length == 0 {
		return nil
	}
	temp := d.head

	if d.length == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
		d.head.prev = nil
		temp.next = nil
	}
	d.length--
	return temp
}

func (d *doublylinkedlist) get(index int) *node {
	if index < 0 || index >= d.length {
		return nil
	}
	temp := d.head
	if index < d.length/2 {
		for i := 0; i < index; i++ {
			temp = temp.next
		}
	} else {
		temp = d.tail
		for i := d.length - 1; i > index; i-- {
			temp = temp.prev
		}
	}
	return temp
}

func (d *doublylinkedlist) set(index int, data int) bool {
	temp := d.get(index)
	if temp != nil {
		temp.data = data
		return true
	}
	return false
}

func (d *doublylinkedlist) insert(index int, data int) bool {
	if index < 0 || index > d.length {
		return false
	}

	if index == 0 {
		d.prepend(data)
		return true
	}

	if index == d.length {
		d.append(data)
		return true
	}
	newNode := &node{data: data}
	before := d.get(index - 1)
	after := before.next
	newNode.prev = before
	newNode.next = after
	before.next = newNode
	after.prev = newNode
	d.length++
	return true
}

func (d *doublylinkedlist) remove(index int) *node {
	if index < 0 || index > d.length {
		return nil
	}

	if index == 0 {
		return d.removeFirst()
	}

	if index == d.length-1 {
		return d.removeLast()
	}

	temp := d.get(index)

	temp.next.prev = temp.prev
	temp.prev.next = temp.next
	temp.next = nil
	temp.next = nil

	d.length--
	return temp
}

func main() {

	d := doublylinkedlist{}

	d.append(1)
	d.append(2)
	d.append(3)

	// d.prepend(4)
	// d.prepend(5)

	// d.removeLast()

	// d.removeFirst()

	//	fmt.Println(d.get(2).data)

	// d.set(2, 5)

	// d.insert(2, 5)

	d.remove(1)

	d.printList()
}
