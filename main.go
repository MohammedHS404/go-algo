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

func permuteAllWithoutIdentity[T any](a []T) {
	n := len(a)
	for i := 0; i < n; i++ {
		tempBefore := -1

		if i > 0 {
			if i == 1 {
				tempBefore = 0
			} else {
				tempBefore = rand.IntN(i)
			}
		}

		tempAfter := -1

		if i < n-1 {
			if i == n-2 {
				tempAfter = n - 1
			} else {
				tempAfter = rand.IntN(n-1-i) + i + 1
			}
		}

		temp := -1

		if tempBefore != -1 && tempAfter != -1 {
			temp = rand.IntN(2)

			if temp == 0 {
				a[i], a[tempBefore] = a[tempBefore], a[i]
			} else {
				a[i], a[tempAfter] = a[tempAfter], a[i]
			}
		} else if tempBefore != -1 {
			a[i], a[tempBefore] = a[tempBefore], a[i]
		} else if tempAfter != -1 {
			a[i], a[tempAfter] = a[tempAfter], a[i]
		} else {
			panic("Dont know what to do")
		}

		a[i], a[temp] = a[temp], a[i]
	}
}
