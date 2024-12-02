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
}
