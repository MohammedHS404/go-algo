package main

import "fmt"

func main() {

	head := NewNode(3)
	head.Append(5)
	head.Append(8)
	head.Append(5)
	head.Append(10)
	head.Append(2)
	head.Append(1)

	fmt.Println(head.String())

	x := 5

	var bxNode *Node[int] = nil
	var bxHead *Node[int] = nil

	var axNode *Node[int] = nil
	var axHead *Node[int] = nil

	var current = head

	for current != nil {
		fmt.Println(current.Data)

		next := current.Next
		current.Next = nil

		if current.Data < x {
			if bxNode == nil {
				bxNode = current
				bxHead = bxNode
			} else {
				bxNode.Next = current
				bxNode = bxNode.Next
			}
		} else {
			if axNode == nil {
				axNode = current
				axHead = axNode
			} else {
				axNode.Next = current
				axNode = axNode.Next
			}
		}

		current = next
	}

	bxNode.Next = axHead

	fmt.Println(bxHead.String())
}
