package main

import "fmt"

type node struct {
	data int
	next *node
}

type stack struct {
	top    *node
	height int
}

func (s *stack) printStack() {
	temp := s.top
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (s *stack) push(data int) {
	newNode := &node{data: data}
	if s.height == 0 {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
	s.height++
}

func (s *stack) pop() *node {
	if s.height == 0 {
		return nil
	}

	temp := s.top
	s.top = s.top.next
	temp.next = nil
	s.height--
	return temp
}

func main() {
	s := stack{}

	s.push(1)
	s.push(2)
	s.push(3)

	s.pop()

	s.printStack()
}
