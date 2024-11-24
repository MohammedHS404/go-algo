package main

import (
	"math/rand/v2"
)

func main() {
	// print item, position, and frequency
	a := []string{"cat", "dog", "Rabbit", "Rat", "Turtle", "Fish", "Bird", "Lion", "Tiger", "Elephant", "Monkey", "Giraffe", "Zebra", "Horse", "Cow", "Pig", "Sheep", "Goat", "Chicken", "Duck", "Turkey", "Penguin", "Parrot", "Ostrich", "Eagle", "Hawk", "Falcon", "Owl", "Sparrow", "Robin", "Bluejay", "Cardinal", "Woodpecker", "Hummingbird", "Seagull", "Pelican", "Swan", "Goose", "Flamingo", "Stork", "Crane", "Heron", "Pigeon", "Crow", "Raven", "Vulture", "Albatross", "Kiwi", "Emu", "Cassowary", "Rhea", "Roadrunner", "Quail", "Loon", "Grebe", "Cormorant", "Anhinga", "Frigatebird", "Booby", "Gannet", "Tern", "Skua", "Jaeger", "Phalarope", "Avocet", "Stilt", "Sandpiper", "Plover", "Killdeer", "Turnstone", "Dunlin", "Sanderling", "Curlew", "Godwit", "Snipe", "Woodcock", "Ruff", "Redshank", "Greenshank", "Spotted Sandpiper", "Solitary Sandpiper", "Willet", "Upland Sandpiper", "Whimbrel", "Curlew Sandpiper", "Bar-tailed Godwit", "Black-tailed Godwit", "Hudsonian Godwit", "Marbled Godwit", "Ruddy Turnstone", "Black Turnstone", "Great Knot", "Red Knot", "Surfbird", "Ruddy Turnstone", "Sanderling", "Dunlin", "Purple Sandpiper", "Rock Sandpiper", "Baird's Sandpiper", "Least Sandpiper", "White-rumped Sandpiper", "Buff-breasted Sandpiper", "Pectoral Sandpiper", "Semipalmated Sandpiper", "Western Sandpiper", "Short-billed Dowitcher", "Long-billed Dowitch"}

	mp := make(map[string]int)

	for i := 0; i < len(a); i++ {
		mp[a[i]] = i
	}

	hashmp := make(map[int]int)

	n := int64(len(a))
	print("n: ", n, "\n")
	n = factorial(int64(n))
	print("n!: ", n, "\n")

	for i := int64(0); i < n; i++ {
		sample := randomSample[string](a, 10)
		hash := 0
		for j := 0; j < len(sample); j++ {
			hash += mp[sample[j]]
		}

		hashmp[hash]++
	}

	for k, v := range hashmp {
		println(k, " ", v)
	}

}
func factorial(number int64) int64 {

	// if the number has reached 1 then we have to
	// return 1 as 1 is the minimum value we have to multiply with
	if number == 1 {
		return 1
	}

	// multiplying with the current number and calling the function
	// for 1 lesser number
	factorialOfNumber := number * factorial(number-1)

	// return the factorial of the current number
	return factorialOfNumber
}

func testPermutation() {
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

func randomSample[T any](a []T, m int) []T {
	s := make([]T, m)
	n := len(a)
	for i := 0; i < m; i++ {
		randomIndex := rand.IntN(n-i) + i
		s[i] = a[randomIndex]
	}

	return s
}
