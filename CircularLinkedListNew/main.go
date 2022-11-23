package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head   *node
	tail   *node
	length int
}

func (l *linkedlist) addlast1(value int) {
	newest := node{data: value}
	if l.head == nil {
		newest.next = &newest
		l.head = &newest
		l.tail = &newest
		l.length++
	} else {
		newest.next = l.head
		l.tail.next = &newest
		l.tail = &newest
		l.length++
	}

}
func (l *linkedlist) addfirst(value int) {
	newest := node{data: value, next: nil}
	if l.head == nil {
		newest.next = &newest
		l.head = &newest
		l.tail = &newest
		l.length++
	}
	l.tail.next = &newest
	newest.next = l.head
	l.head = &newest
}
func (l *linkedlist) addany(value int, position int) {
	newest := node{data: value}
	if l.head == nil {
		newest.next = &newest
		l.head = &newest
		l.tail = &newest
		l.length++
	}
	i := 1
	p := l.head
	for i < position-1 {
		p = p.next
		i++
	}
	newest.next = p.next
	p.next = &newest
	l.length++
}
func (l *linkedlist) delfirst() int {
	if l.head == nil {
		fmt.Println("there are no nodes to delete")
	}
	e := l.head.data
	l.tail.next = l.head.next
	l.head = l.head.next
	l.length--
	return e

}
func (l *linkedlist) dellast() int {
	if l.head == nil {
		fmt.Println("there are no nodes to delete")
	}
	p := l.head
	for p.next != l.tail {
		p = p.next
	}
	l.tail = p
	p = p.next
	l.tail.next = l.head
	e := p.data
	l.length--
	return e

}
func (l *linkedlist) delany(postion int) int {
	if l.head == nil {
		fmt.Println("there are no nodes to delete ")

	}
	i := 1
	p := l.head
	for i < postion-1 {
		p = p.next
		i++
	}
	e := p.next.data
	p.next = p.next.next
	return e

}
func (l *linkedlist) display1() {
	if l.head == nil {
		fmt.Println("no data found")
		return
	}
	p := l.head
	for p != l.tail {
		fmt.Println(p.data)
		p = p.next
	}
	fmt.Println(l.tail.data)

}
func main() {
	c := linkedlist{}
	c.addlast1(20)
	c.addlast1(30)
	c.addlast1(40)
	c.addfirst(50)
	c.addfirst(60)
	c.addany(10, 3)
	c.delfirst()
	c.dellast()
	c.delany(3)
	c.display1()
}
