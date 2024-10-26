package main

import "fmt"

func main() {
	head := NewNode(1)
	head.Append(2)
	head.Append(3)
	head.Append(4)
	head.Append(5)
	head.Append(6)
	head.Append(7)

	node := head

	k := 6

	for i := 0; i < k; i++ {
		node = node.Next
	}

	deleteMiddleNode(node)

	fmt.Println(head.String())
}

func deleteMiddleNode(node *Node[int]) {
	if node == nil || node.Next == nil {
		return
	}

	node.Data = node.Next.Data
	node.Next = node.Next.Next
}
