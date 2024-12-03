package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines, err := utils.ReadLines("day-03/input.txt")
	if err != nil {
		panic(err)
	}

	totalValue := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, mul := range matches {
			fmt.Println("Full match:", mul[0])
			fmt.Println("Group 1:", mul[1])
			fmt.Println("Group 2:", mul[2])

			number1, err := strconv.Atoi(mul[1])
			if err != nil {
				panic(err)
			}

			number2, err := strconv.Atoi(mul[2])
			if err != nil {
				panic(err)
			}

			totalValue += number1 * number2
		}
	}

	fmt.Println("Total:", totalValue)
}

// Cool things and learnings:
// - Go regex usage is pretty straightforward. The capture groups and sub matching is pretty cool and useful.
