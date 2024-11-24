package main

func main() {
	arr := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}

	println("Before heapify")
	printTreeBFS(arr)

	println("After heapify")

	HeapSort(arr, false)

	printTreeBFS(arr)
}

func HeapSort(arr []int, isMax bool) {
	buildHeap(arr, isMax)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		if isMax {
			maxHeapify(arr[:i], 0)
		} else {
			minHeapify(arr[:i], 0)
		}
	}
}

func printTreeBFS(arr []int) {

	println("Printing tree BFS")

	println("Tree")

	level := 0
	for i := 0; i < len(arr); {
		for j := 0; j < 1<<level && i < len(arr); j++ {
			print(arr[i], " ")
			i++
		}
		println()
		level++
	}

	println("Array")

	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}

	println()
}

func buildHeap(arr []int, isMax bool) {
	for i := len(arr) / 2; i >= 0; i-- {
		if isMax {
			maxHeapify(arr, i)
		} else {
			minHeapify(arr, i)
		}
	}
}

func maxHeapify(arr []int, i int) {
	left := left(i)
	right := right(i)

	largest := i

	if left < len(arr) && arr[left] > arr[i] {
		largest = left
	}

	if right < len(arr) && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest)
	}
}

func minHeapify(arr []int, i int) {
	left := left(i)
	right := right(i)

	smallest := i

	if left < len(arr) && arr[left] < arr[i] {
		smallest = left
	}

	if right < len(arr) && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		minHeapify(arr, smallest)
	}
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func parent(i int) int {
	return i / 2
}
