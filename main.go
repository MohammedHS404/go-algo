package main

import "fmt"

func main() {
	node := NewNode(1)
	node.Append(2)
	node.Append(3)
	node.Append(4)
	node.Append(5)
	node.Append(6)
	node.Append(7)

	k := 3

	kth := getKthToLastRecursive(node, k, 0)

	fmt.Println(kth.String())
}

func getKthToLastRecursive(node *Node[int], k int, counter int) *Node[int] {
	if node == nil {
		return nil
	}

	current := getKthToLastRecursive(node.Next, k, counter+1)

	if counter == k {
		return node
	}

	return current
}

func getKthToLastIterative(node *Node[int], k int) *Node[int] {
	counter := 0
	current := node

	pointer := node

	for current != nil {
		if counter == k {
			pointer = current
			break
		}

		counter++
		current = current.Next
	}
	return pointer
}
