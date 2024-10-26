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

	var bxNode *Node[int] = NewNode(-1)
	bxHead := bxNode

	var axNode *Node[int] = NewNode(-1)
	axHead := axNode

	var current = head

	for current != nil {
		if current.Data < x {
			if bxNode.Data == -1 {
				bxNode.Data = current.Data
			} else {
				bxNode.Append(current.Data)
				bxNode = bxNode.Next
			}
		} else {
			if axNode.Data == -1 {
				axNode.Data = current.Data
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
