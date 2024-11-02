package main

func main() {
	// 4x4

	a := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	b := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	c := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	n := 4

	matrix_strassen(a, b, c, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, n)
	// matrix_multiply_recursive(a, b, c, 0, 0, 0, 0, 0, 0, n)

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

func matrix_strassen(
	A [][]int,
	B [][]int,
	C [][]int,
	Ar1 int,
	Ar2 int,
	Ac1 int,
	Ac2 int,
	Br1 int,
	Br2 int,
	Bc1 int,
	Bc2 int,
	Cr1 int,
	Cr2 int,
	Cc1 int,
	Cc2 int,
	n int) {
	// Assuming 2x2

	if n == 2 {
		S1 := B[Br1][Bc2] - B[Br2][Bc2]
		S2 := A[Ar1][Ac1] + A[Ar1][Ac2]
		S3 := A[Ar2][Ac1] + A[Ar2][Ac2]
		S4 := B[Br2][Bc1] - B[Br1][Bc1]
		S5 := A[Ar1][Ac1] + A[Ar2][Ac2]
		S6 := B[Br1][Bc1] + B[Br2][Bc2]
		S7 := A[Ar1][Ac2] - A[Ar2][Ac2]
		S8 := B[Br2][Bc1] + B[Br2][Bc2]
		S9 := A[Ar1][Ac1] - A[Ar2][Ac1]
		S10 := B[Br1][Bc1] + B[Br1][Bc2]

		P1 := A[Ar1][Ac1] * S1
		P2 := S2 * B[Br2][Bc2]
		P3 := S3 * B[Br1][Bc1]
		P4 := A[Ar2][Ac2] * S4
		P5 := S5 * S6
		P6 := S7 * S8
		P7 := S9 * S10

		C[Cr1][Cc1] += P5 + P4 - P2 + P6
		C[Cr1][Cc2] += P1 + P2
		C[Cr2][Cc1] += P3 + P4
		C[Cr2][Cc2] += P5 + P1 - P3 - P7
		return
	}

	n2 := n / 2

	// C11 = A11 * B11 + A12 * B21
	// A11 * B11
	matrix_strassen(A, B, C, Ar1, Ar2, Ac1, Ac2, Br1, Br2, Bc1, Bc2, Cr1, Cr2, Cc1, Cc2, n2)

	// A12 * B21
	matrix_strassen(A, B, C, Ar1, Ar2, Ac1+n2, Ac2+n2, Br1+n2, Br2+n2, Bc1, Bc2, Cr1, Cr2, Cc1, Cc2, n2)

	// C12 = A11 * B12 + A12 * B22

	// A11 * B12

	matrix_strassen(A, B, C, Ar1, Ar2, Ac1, Ac2, Br1, Br2, Bc1+n2, Bc2+n2, Cr1, Cr2, Cc1+n2, Cc2+n2, n2)

	// A12 * B22

	matrix_strassen(A, B, C, Ar1, Ar2, Ac1+n2, Ac2+n2, Br1+n2, Br2+n2, Bc1+n2, Bc2+n2, Cr1, Cr2, Cc1+n2, Cc2+n2, n2)

	// C21 = A21 * B11 + A22 * B21

	// A21 * B11

	matrix_strassen(A, B, C, Ar1+n2, Ar2+n2, Ac1, Ac2, Br1, Br2, Bc1, Bc2, Cr1+n2, Cr2+n2, Cc1, Cc2, n2)

	// A22 * B21

	matrix_strassen(A, B, C, Ar1+n2, Ar2+n2, Ac1+n2, Ac2+n2, Br1+n2, Br2+n2, Bc1, Bc2, Cr1+n2, Cr2+n2, Cc1, Cc2, n2)

	// C22 = A21 * B12 + A22 * B22

	// A21 * B12

	matrix_strassen(A, B, C, Ar1+n2, Ar2+n2, Ac1, Ac2, Br1, Br2, Bc1+n2, Bc2+n2, Cr1+n2, Cr2+n2, Cc1+n2, Cc2+n2, n2)

	// A22 * B22

	matrix_strassen(A, B, C, Ar1+n2, Ar2+n2, Ac1+n2, Ac2+n2, Br1+n2, Br2+n2, Bc1+n2, Bc2+n2, Cr1+n2, Cr2+n2, Cc1+n2, Cc2+n2, n2)
}
