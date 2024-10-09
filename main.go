package main

func main() {
	A := []int{5, 2, 4, 6, 1, 3}
	n := len(A)
	insertionSort(A, n)
	for i := 0; i < n; i++ {
		println(A[i])
	}
}

func insertionSort(A []int, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j] > key {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
}
