package main

import "fmt"

// Stack is the Abstract Data type following FIFO (First In First Out) principle which has certain functionalities like-
// (a) push- Insert new data at one end
// (b) pop- Remove (and return) data from the end insertion has happened
// (c) peek- Return the data from the end insertion has happened (without removal)

// There are many real-life examples of a stack. Consider an example of plates stacked over one
// another in the canteen. The plate which is at the top is the first one to be removed, i.e. the plate
// which has been placed at the bottommost position remains in the stack for the longest period of time. So, it
// can be simply seen to follow LIFO(Last In First Out)/FILO(First In Last Out) order.

// Since Interface declares the functionalities of the type that defines it,
// declaring the above mentioned methods in the interfaces essentially decalres
// a Stack ADT.
// Any data type which implements the Stack interface becomes the Stack ADT. it can
// be an array, a Linked List, etc. In this case we are implementing Stack on a Linked List.
// So Stack is the Abstract Data type and Linked List is the data structure here.

type Stack interface {
	push(item int)
	pop() int
	peek()
	isEmpty() bool
}

// Check if stack is empty
func (s *stack) isEmpty() bool {
	if s.head == nil {
		return true
	}
	return false
}

// Node of the linked list
type node struct {
	data int
	next *node
}

// Stack always stores the address of the head node
type stack struct {
	head *node
}

// push new element at one end
func (s *stack) push(item int) {
	newNode := node{
		data: item,
		next: nil,
	}
	if s.head == nil {
		s.head = &newNode
	} else {
		temp := s.head
		newNode.next = temp
		s.head = &newNode
	}
	fmt.Printf("PUSHED: %d\n", newNode.data)
	printStack(s)
	return
}

// remove and return the element from the same end as insertion
func (s *stack) pop() int {
	if s.isEmpty() {
		fmt.Println("STACK UNDERFLOW!")
		return 0
	}
	temp := s.head
	s.head = s.head.next
	return temp.data
}

// peek the top of the stack
func (s *stack) peek() {
	temp := s.head
	for temp.next != nil {
		temp = temp.next
	}
	fmt.Println(temp.data)
}

func printStack(s *stack) {
	if s.head == nil {
		return
	}
	temp := s.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

// constructor for new node
func new() Stack {
	return &stack{
		head: nil,
	}
}
func main() {
	st := new()
	st.push(1)
	st.push(2)
	st.push(3)
	st.push(4)
	st.push(5)
	fmt.Println("POP:", st.pop())
	fmt.Println("POP:", st.pop())
	fmt.Println("POP:", st.pop())
	fmt.Println("POP:", st.pop())
	fmt.Println("POP:", st.pop())
	fmt.Println("POP:", st.pop())
}
