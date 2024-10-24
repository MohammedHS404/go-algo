package main

import "fmt"

func main() {
	// 4x4 matrix
	a := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	b := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	c := [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	n := 4

	matrix_multiply_recursive(a, b, c, 0, 0, 0, 0, 0, 0, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
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

func matrix_multiply_recursive(
	A [][]int, B [][]int, C [][]int,
	aRow int, aCol int, bRow int, bCol int, cRow int, cCol int, n int) {
	if n == 1 {
		fmt.Printf("A[%d][%d] * B[%d][%d]\n", aRow, aCol, bRow, bCol)
		C[cRow][cCol] += A[aRow][aCol] * B[bRow][bCol]
		return
	}

	n2 := n / 2

	// C11 = A11 * B11 + A12 * B21
	// A11 * B11
	matrix_multiply_recursive(A, B, C, aRow, aCol, bRow, bCol, cRow, cCol, n2)
	// A12 * B21
	matrix_multiply_recursive(A, B, C, aRow, aCol+n2, bRow+n2, bCol, cRow, cCol, n2)
	// C12 = A11 * B12 + A12 * B22
	// A11 * B12
	matrix_multiply_recursive(A, B, C, aRow, aCol, bRow, bCol+n2, cRow, cCol+n2, n2)
	// A12 * B22
	matrix_multiply_recursive(A, B, C, aRow, aCol+n2, bRow+n2, bCol+n2, cRow, cCol+n2, n2)
	// C21 = A21 * B11 + A22 * B21
	// A21 * B11
	matrix_multiply_recursive(A, B, C, aRow+n2, aCol, bRow, bCol, cRow+n2, cCol, n2)
	// A22 * B21
	matrix_multiply_recursive(A, B, C, aRow+n2, aCol+n2, bRow+n2, bCol, cRow+n2, cCol, n2)
	// C22 = A21 * B12 + A22 * B22
	// A21 * B12
	matrix_multiply_recursive(A, B, C, aRow+n2, aCol, bRow, bCol+n2, cRow+n2, cCol+n2, n2)
	// A22 * B22
	matrix_multiply_recursive(A, B, C, aRow+n2, aCol+n2, bRow+n2, bCol+n2, cRow+n2, cCol+n2, n2)
}
