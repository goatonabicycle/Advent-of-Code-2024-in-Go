package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	lines, err := utils.ReadLines("day-01/input.txt")
	if err != nil {
		panic(err)
	}

	var left, right []int
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			var l, r int
			fmt.Sscanf(parts[0], "%d", &l)
			fmt.Sscanf(parts[1], "%d", &r)
			left = append(left, l)
			right = append(right, r)
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i, l := range left {
		totalDistance += int(math.Abs(float64(l - right[i])))
	}

	fmt.Println("Sum:", totalDistance)

	// Part 2
	rightCounts := make(map[int]int)
	for _, r := range right {
		rightCounts[r]++
	}

	totalSimilarity := 0
	for _, l := range left {
		totalSimilarity += l * rightCounts[l]
	}

	fmt.Println("Similarity:", totalSimilarity)
}

// Cool things and learnings:
// - strings.Fields is JS's split(" ")
// - Sscanf parses strings into numbers using the format %d (decimal)
// - The pointer idea is interesting. Think of & as "put the result in this variable"
// - Go is a bit weird about math.Abs with floats. Having to typecast to int feels a bit weird. But types!
// - Go doesn't have a filter function like JS. I could have either made a helper function or just iterate.
// - Even better, I could use a map to count the occurrences of each number.
