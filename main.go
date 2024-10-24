package main

func main() {
	a := [][]int{{1, 2}, {3, 4}}
	b := [][]int{{1, 2}, {3, 4}}
	c := [][]int{{0, 0}, {0, 0}}

	matrix_multiply(a, b, c, 2)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			print(c[i][j], " ")
		}
		println()
	}
}

func matrix_multiply(A [][]int, B [][]int, C [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
}

func matrix_multiply_recursive(A [][]int, B [][]int, C [][]int, n int) {
	if n == 1 {
		C[0][0] = A[0][0] * B[0][0]
		return
	}

	n2 := n / 2

	A11 := make([][]int, n2)
	A12 := make([][]int, n2)
	A21 := make([][]int, n2)
	A22 := make([][]int, n2)

	for i := 0; i < n2; i++ {
		A11[i] = make([]int, n2)
		A12[i] = make([]int, n2)
		A21[i] = make([]int, n2)
		A22[i] = make([]int, n2)
	}

	B11 := make([][]int, n2)
	B12 := make([][]int, n2)
	B21 := make([][]int, n2)
	B22 := make([][]int, n2)

	for i := 0; i < n2; i++ {
		B11[i] = make([]int, n2)
		B12[i] = make([]int, n2)
		B21[i] = make([]int, n2)
		B22[i] = make([]int, n2)
	}

	C11 := make([][]int, n2)
	C12 := make([][]int, n2)
	C21 := make([][]int, n2)
	C22 := make([][]int, n2)

	for i := 0; i < n2; i++ {
		C11[i] = make([]int, n2)
		C12[i] = make([]int, n2)
		C21[i] = make([]int, n2)
		C22[i] = make([]int, n2)
	}

	for i := 0; i < n2; i++ {
		for j := 0; j < n2; j++ {
			A11[i][j] = A[i][j]
			A12[i][j] = A[i][j+n2]
			A21[i][j] = A[i+n2][j]
			A22[i][j] = A[i+n2][j+n2]

			B11[i][j] = B[i][j]
			B12[i][j] = B[i][j+n2]
			B21[i][j] = B[i+n2][j]
			B22[i][j] = B[i+n2][j+n2]
		}
	}

	matrix_multiply_recursive(A11, B11, C11, n2/2)
	matrix_multiply_recursive(A11, B12, C12, n2/2)

	matrix_multiply_recursive(A21, B11, C21, n2/2)
	matrix_multiply_recursive(A21, B12, C22, n2/2)

	matrix_multiply_recursive(A12, B21, C11, n2/2)
	matrix_multiply_recursive(A12, B22, C12, n2/2)

	matrix_multiply_recursive(A22, B21, C21, n2/2)
	matrix_multiply_recursive(A22, B22, C22, n2/2)
}
