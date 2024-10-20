package main

import (
	"math"
	"sort"
)

func main() {
	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	fleets := carFleet(target, speed, position)
	println(fleets)
}

func carFleet(target int, speed []int, position []int) int {
	n := len(position)

	position, speed = sortPS(position, speed)

	fleets := 0

	currentTime := 0

	for i := n - 1; i >= 0; i-- {
		timeLeft := int(math.Ceil(float64(target-position[i]) / float64(speed[i])))

		if timeLeft != currentTime && timeLeft >= currentTime {
			fleets++
			currentTime = timeLeft
		}
	}

	return fleets
}

func sortPS(position, speed []int) ([]int, []int) {
	n := len(position)
	cars := make([][2]int, n)

	for i := 0; i < n; i++ {
		cars[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] < cars[j][0]
	})

	for i := 0; i < n; i++ {
		position[i] = cars[i][0]
		speed[i] = cars[i][1]
	}

	return position, speed
}
