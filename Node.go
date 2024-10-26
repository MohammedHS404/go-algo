package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Node represents a singly linked list node
type Node[T any] struct {
	Next *Node[T]
	Data T
}

// NewNode creates a new node with given data
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{Data: data}
}

// Append adds a new node with the specified data at the end of the list
func (n *Node[T]) Append(data T) {
	end := NewNode(data)
	next := n

	for next.Next != nil {
		next = next.Next
	}

	next.Next = end
}

// Delete removes the first occurrence of a node with the specified data
func (n *Node[T]) Delete(data T) {
	current := n

	for current.Next != nil {
		if reflect.DeepEqual(current.Next.Data, data) {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// String provides a string representation of the linked list
func (n *Node[T]) String() string {
	var sb strings.Builder
	current := n

	for current != nil {
		sb.WriteString(fmt.Sprintf("%v -> ", current.Data))
		current = current.Next
	}

	// Remove the last " -> "
	str := sb.String()
	if len(str) > 4 {
		str = str[:len(str)-4]
	}

	return str
}

// FromArray creates a linked list from an array
func FromArray[T any](array []T) *Node[T] {
	if len(array) == 0 {
		return nil
	}

	head := NewNode(array[0])
	current := head

	for _, data := range array[1:] {
		current.Next = NewNode(data)
		current = current.Next
	}

	return head
}
