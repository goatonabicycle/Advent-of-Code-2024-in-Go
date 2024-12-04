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
	enabled := true

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			fullMatch := match[0]

			switch {
			case fullMatch == "don't()":
				enabled = false
			case fullMatch == "do()":
				enabled = true
			case len(match[1]) > 0:
				if enabled {
					num1, _ := strconv.Atoi(match[1])
					num2, _ := strconv.Atoi(match[2])
					totalValue += num1 * num2
				}
			}
		}
	}

	fmt.Println("Total:", totalValue)
}

// Cool things and learnings:
// - Go regex usage is pretty straightforward. The capture groups and sub matching is pretty cool and useful.
