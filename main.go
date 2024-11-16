package main

import (
	"math/rand/v2"
)

func main() {
	a := []string{"cat", "dog", "Rabbit", "Rat", "Turtle"}

	mp := make(map[string]map[int]int)

	for i := 0; i < 100000; i++ {
		randomPermuteWithoutIdentity(a)
		for j := 0; j < len(a); j++ {
			mp[a[j]][j]++
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			print(mp[a[i]][j], " ")
		}
		println()
	}

}

func randomPermuteWithoutIdentity[T any](a []T) {
	n := len(a)
	for i := 0; i < n; i++ {
		temp := i + rand.IntN(n)
		for temp == i {
			temp = rand.IntN(n)
		}
		a[i], a[temp] = a[temp], a[i]
	}
}
