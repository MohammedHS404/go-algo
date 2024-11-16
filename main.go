package main

import "math/rand"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	randomPermute(a)
	for _, v := range a {
		println(v)
	}
}

func randomPermute[T any](a []T) {
	n := len(a)
	for i := 0; i < n; i++ {
		temp := rand.Intn(n)
		a[i], a[temp] = a[temp], a[i]
	}
}
