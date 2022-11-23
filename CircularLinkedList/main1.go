package main

import "fmt"

type node struct {
	data int
	next *node
}
type circular struct {
	head   *node
	tail   *node
	length int
}

func (c *circular) addlast1(value int) {
	newest := node{data: value}
	if c.head == nil {
		newest.next = &newest
		c.head = &newest
		c.tail = &newest
		c.length++
	} else {
		newest.next = c.head
		c.tail.next = &newest
		c.tail = &newest
		c.length++
	}

}
func (c *circular) addfirst(value int) {
	newest := node{data: value, next: nil}
	if c.head == nil {
		newest.next = &newest
		c.head = &newest
		c.tail = &newest
		c.length++
	}
	c.tail.next = &newest
	newest.next = c.head
	c.head = &newest
}
func (c *circular) addany(value int, position int) {
	newest := node{data: value}
	if c.head == nil {
		newest.next = &newest
		c.head = &newest
		c.tail = &newest
		c.length++
	}
	i := 1
	p := c.head
	for i < position-1 {
		p = p.next
		i++
	}
	newest.next = p.next
	p.next = &newest
	c.length++
}
func (c *circular) delfirst() int {
	if c.head == nil {
		fmt.Println("there are no nodes to delete")
	}
	e := c.head.data
	c.tail.next = c.head.next
	c.head = c.head.next
	c.length--
	return e

}
func (c *circular) dellast() int {
	if c.head == nil {
		fmt.Println("there are no nodes to delete")
	}
	p := c.head
	for p.next != c.tail {
		p = p.next
	}
	c.tail = p
	p = p.next
	c.tail.next = c.head
	e := p.data
	c.length--
	return e

}
func (c *circular) delany(postion int) int {
	if c.head == nil {
		fmt.Println("there are no nodes to delete ")

	}
	i := 1
	p := c.head
	for i < postion-1 {
		p = p.next
		i++
	}
	e := p.next.data
	p.next = p.next.next
	return e

}
func (c *circular) display1() {
	if c.head == nil {
		fmt.Println("no data found")
		return
	}
	p := c.head
	for p != c.tail {
		fmt.Println(p.data)
		p = p.next
	}
	fmt.Println(c.tail.data)

}
func main() {
	c := circular{}
	c.addlast1(20)
	c.addlast1(30)
	c.addlast1(40)
	/*c.addfirst(50)
	c.addfirst(60)
	c.addany(10, 3)
	c.delfirst()
	c.dellast()
	c.delany(3)*/
	c.display1()
}
