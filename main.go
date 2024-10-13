package main

func main() {
	A := []int{2, 4, 7, 6, 1, 3, 2, 5}
	n := len(A)
	mergeSort(A, 0, n-1)
	for i := 0; i < n; i++ {
		println(A[i])
	}
}

func mergeSort(A []int, start int, end int) {
	if start >= end {
		return
	}

	midpoint := (start + end) / 2

	mergeSort(A, start, midpoint)
	mergeSort(A, midpoint+1, end)

	merge(A, start, midpoint, end)
}

func merge(A []int, start int, midpoint int, end int) {
	if start >= end {
		return
	}

	leftArrayLength := midpoint - start + 1

	rightArrayLength := end - midpoint

	leftArray := make([]int, leftArrayLength)

	for i := 0; i < leftArrayLength; i++ {
		leftArray[i] = A[i+start]
	}

	rightArray := make([]int, rightArrayLength)

	for i := 0; i < rightArrayLength; i++ {
		rightArray[i] = A[midpoint+i+1]
	}

	i_left := 0
	i_right := 0
	i := start

	for i_left < leftArrayLength && i_right < rightArrayLength {
		if leftArray[i_left] < rightArray[i_right] {
			A[i] = leftArray[i_left]
			i_left++
		} else {
			A[i] = rightArray[i_right]
			i_right++
		}
		i++
	}

	for i_left < leftArrayLength {
		A[i] = leftArray[i_left]
		i_left++
		i++
	}

	for i_right < rightArrayLength {
		A[i] = rightArray[i_right]
		i_right++
		i++
	}
}
