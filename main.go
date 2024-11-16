package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	a := []string{"cat", "dog", "Rabbit", "Rat", "Turtle"}
	randomPermute(a)

	for _, v := range a {
		fmt.Println(v)
	}
}

func randomPermute[T any](a []T) {
	n := len(a)
	for i := 0; i < n; i++ {
		temp := rand.IntN(n)
		a[i], a[temp] = a[temp], a[i]
	}
}
