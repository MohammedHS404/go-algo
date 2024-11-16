package main

import (
	"math/rand/v2"
)

func main() {
	a := []string{"cat", "dog", "Rabbit", "Rat", "Turtle"}

	mp := make(map[string]map[int]int)

	for i := 0; i < len(a); i++ {
		mp[a[i]] = make(map[int]int)
	}

	for i := 0; i < 100000; i++ {
		a = []string{"cat", "dog", "Rabbit", "Rat", "Turtle"}
		permuteWithoutIdentity(a)
		for j := 0; j < len(a); j++ {
			mp[a[j]][j]++
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(a); k++ {
				// print item, position, and frequency
				println(a[j], " ", k, " ", mp[a[j]][k], " ")
			}
		}
		println()
	}

}

func permuteRandom[T any](a []T) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		temp := rand.IntN(n-i) + i
		a[i], a[temp] = a[temp], a[i]
	}
}

func permuteRandomWithAll[T any](a []T) {
	n := len(a)
	for i := 0; i < n; i++ {
		temp := rand.IntN(n)
		a[i], a[temp] = a[temp], a[i]
	}
}

func permuteWithoutIdentity[T any](a []T) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		temp := rand.IntN(n-1-i) + i + 1
		a[i], a[temp] = a[temp], a[i]
	}
}
