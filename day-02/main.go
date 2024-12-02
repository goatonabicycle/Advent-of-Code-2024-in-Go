package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, err := utils.ReadLines("day-02/input.txt")
	if err != nil {
		panic(err)
	}

	totalSafeLines := 0

	for _, line := range lines {
		if isLineSafe(strings.Fields(line)) {
			totalSafeLines++
		}
	}

	fmt.Println("Safe lines:", totalSafeLines)
}

func isLineSafe(numbers []string) bool {
	nums := make([]int, len(numbers))
	for i, n := range numbers {
		num, err := strconv.Atoi(n)
		if err != nil {
			return false
		}
		nums[i] = num
	}

	increasing := true

	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]

		if i == 0 {
			increasing = diff > 0
		}

		if utils.Abso(diff) > 3 || utils.Abso(diff) == 0 {
			return false
		}

		if increasing && diff < 0 {
			return false
		}
		if !increasing && diff > 0 {
			return false
		}
	}

	return true
}

// Cool things and learnings:
// Go's strconv package is used for string conversions. From the docs:
//    - Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
// Slices can be created with a predefined size using make(). Then it's easier to iterate over them without having to check the length.
