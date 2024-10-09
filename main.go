package main

func main() {
	A := []int{5, 2, 4, 6, 1, 3}
	n := len(A)
	quickSort(A, n)
	for i := 0; i < n; i++ {
		println(A[i])
	}
}

func quickSort(A []int, n int) {
	for i := 0; i < n; i++ {
		j := i + 1
		for j < n {
			if A[j] < A[i] {
				temp := A[i]
				A[i] = A[j]
				A[j] = temp
			}
			j++
		}
	}
}
