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
