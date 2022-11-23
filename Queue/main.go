package main

import "fmt"

type node struct {
	data int
	next *node
}

type queue struct {
	first  *node
	last   *node
	length int
}

func (q *queue) printQueue() {
	temp := q.first

	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (q *queue) enQueue(data int) {
	newNode := &node{data: data}

	if q.length == 0 {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length++
}

func (q *queue) deQueue() *node {
	if q.length == 0 {
		return nil
	}

	temp := q.first

	if q.length == 1 {
		q.first = nil
		q.last = nil
	} else {
		q.first = q.first.next
		temp.next = nil
	}
	q.length--
	return temp
}

func main() {
	q := queue{}

	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)

	q.deQueue()

	q.printQueue()
}
