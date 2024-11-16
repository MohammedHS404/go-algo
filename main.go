package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

func main() {
	a := []string{"cat", "dog", "Rabbit", "Rat", "Turtle"}
	weights := []int{7, 10, 5, 1, 3}

	mp := make(map[string]int)

	for i := 0; i < 1000000; i++ {
		s, _ := weightedRandom(a, weights)
		mp[s]++
	}

	type kv struct {
		Key   string
		Value int
	}

	var sortedMp []kv

	for k, v := range mp {
		sortedMp = append(sortedMp, kv{k, v})
	}

	sort.Slice(sortedMp, func(i, j int) bool {
		return sortedMp[i].Value > sortedMp[j].Value
	})

	for _, kv := range sortedMp {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}

}

func weightedRandom(a []string, weights []int) (string, error) {
	n := len(a)
	if n == 0 || len(weights) != n {
		return "", fmt.Errorf("weightedRandom: invalid input")
	}

	accumWeights := make([]int, n)

	accumWeights[0] = weights[0]

	for i := 1; i < n; i++ {
		if weights[i] < 0 {
			return "", fmt.Errorf("weightedRandom: negative weights are not allowed")
		}
		accumWeights[i] = accumWeights[i-1] + weights[i]
	}

	if accumWeights[n-1] == 0 {
		return "", fmt.Errorf("weightedRandom: sum of weights must be greater than 0")
	}

	wi := rand.IntN(accumWeights[n-1]) // random number in the range [0, total weight)

	// Use binary search to find the correct index
	index := sort.Search(n, func(i int) bool { return accumWeights[i] > wi })

	return a[index], nil
}
