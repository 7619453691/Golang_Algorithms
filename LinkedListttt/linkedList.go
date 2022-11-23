package linkedList

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	tail   *node
	length int
}

func (L *LinkedList) Length() int {
	return L.length
}

func (L *LinkedList) PrintList() {
	temp := L.head
	for temp != nil {
		fmt.Printf("%+v ->", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func (L *LinkedList) Insert(data int) {
	newNode := &node{data: data}
	if L.head == nil {
		L.head = newNode
	} else {
		L.tail.next = newNode

	}
	L.tail = newNode
	L.length++
}

func (L *LinkedList) Prepend(data int) {
	newNode := &node{data: data}
	if L.length == 0 {
		L.head = newNode
		L.tail = newNode
	} else {
		newNode.next = L.head
		L.head = newNode
	}
	L.length++
}

func (L *LinkedList) RemoveFirst() *node {
	if L.head == nil {
		return nil
	}
	temp := L.head
	if L.length == 1 {
		L.head = nil
		L.tail = nil
	} else {
		L.head = L.head.next
		temp.next = nil
	}
	L.length--
	return temp
}

func (L *LinkedList) Removelast() *node {
	if L.length == 0 {
		return nil
	}
	temp := L.head
	if L.length == 1 {
		L.head = nil
		L.tail = nil
	} else {
		pre := L.head
		for temp.next != nil {
			pre = temp
			temp = temp.next
		}
		L.tail = pre
		L.tail.next = nil
	}
	L.length--
	return temp
}

func (L *LinkedList) Reverse() {
	temp := L.head
	L.head = L.tail
	L.tail = temp
	var after *node
	var before *node
	for i := 0; i < L.length; i++ {
		after = temp.next
		temp.next = before
		before = temp
		temp = after
	}
}

func (L *LinkedList) GetNode(index int) *node {
	if index < 0 || index >= L.length {
		return nil
	}
	temp := L.head

	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp
}

func (L *LinkedList) InsertAt(index, value int) bool {
	if index < 0 || index > L.length {
		return false
	}
	if index == 0 {
		L.Prepend(value)
		return true
	}
	if index == L.length {
		L.Insert(value)
		return true
	}
	newNode := &node{data: value}
	temp := L.GetNode(index - 1)
	newNode.next = temp.next
	temp.next = newNode
	L.length++
	return true
}

func (L *LinkedList) RemoveAt(index int) *node {
	if index < 0 || index >= L.length {
		return nil
	}
	if index == 0 {
		L.RemoveFirst()
	}
	if index == L.length-1 {
		return L.Removelast()
	}
	prev := L.GetNode(index - 1)
	temp := prev.next
	prev.next = temp.next
	temp.next = nil
	L.length--
	return temp
}

func (L *LinkedList) SetNode(index, value int) bool {
	temp := L.GetNode(index)
	if temp != nil {
		temp.data = value
		return true
	}
	return false
}

func (L *LinkedList) GetHead() {
	fmt.Println("Head: ", L.head)
}
func (L *LinkedList) GetTail() {
	fmt.Println("Head: ", L.tail)
}
