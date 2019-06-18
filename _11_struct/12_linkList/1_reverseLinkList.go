/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */


//反转链表
//Reverse A Linked List Using Go
package main

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

// Get head / first
func (list *List) First() *Node {
	return list.head

}

// Add to tail
func (list *List) Add(value int) {
	newNode := &Node{data: value}

	if list.head == nil {
		list.head = newNode
	} else {
		list.tail.next = newNode
	}
	list.tail = newNode
}


func reverse(list *List) {

	// Set current node to head.
	var currentNode = list.head

	// Set nexNode and prevNode to nil.
	var nextNode *Node = nil
	var prevNode *Node = nil

	// Loop from currentNode to nil.
	for currentNode != nil {
		// Set nextNode to currentNode.next.
		nextNode = currentNode.next

		// Set currentNode.next to prevNode
		currentNode.next = prevNode

		// Set prevNode to currentNode.
		prevNode = currentNode

		// Set currentNode to nextNode.
		currentNode = nextNode
	}
	// set head to prevNode
	list.head = prevNode
}