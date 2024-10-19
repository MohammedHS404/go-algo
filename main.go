package main

func main() {
	A := []int{5, 2, 4, 6, 1, 3}
	n := len(A)
	insertionSort(A, 0, n)
	for i := 0; i < n; i++ {
		println(A[i])
	}
}

func insertionSortSort(A []int, i int) {
	key := A[i]
	j := i - 1
	for j >= 0 && A[j] > key {
		A[j+1] = A[j]
		j--
	}
	A[j+1] = key
}

func insertionSort(A []int, i int, n int) {
	if i == n {
		return
	}

	insertionSortSort(A, i)

	insertionSort(A, i+1, n)
}
