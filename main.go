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
		if current.Data < x {
			if bxNode == nil {
				bxNode = NewNode(current.Data)
				bxHead = bxNode
			} else {
				bxNode.Append(current.Data)
				bxNode = bxNode.Next
			}
		} else {
			if axNode == nil {
				axNode = NewNode(current.Data)
				axHead = axNode
			} else {
				axNode.Append(current.Data)
				axNode = axNode.Next
			}
		}

		current = current.Next
	}

	bxNode.Next = axHead

	fmt.Println(bxHead.String())
}
