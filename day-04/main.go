package main

import (
	"advent-of-code-2024/utils"
	"fmt"
)

func main() {
	lines, err := utils.ReadLines("day-04/input.txt")
	if err != nil {
		panic(err)
	}

	totalXMASItems := 0
	for y, line := range lines {
		for x := range line {
			if isXMASForward(x, line) {
				totalXMASItems++
			}

			if isXMASBackward(x, line) {
				totalXMASItems++
			}

			if isXMASDown(x, y, lines) {
				totalXMASItems++
			}

			if isXMASUp(x, y, lines) {
				totalXMASItems++
			}

			if isXMASDiagonalLeftUp(x, y, lines) {
				totalXMASItems++
			}

			if isXMASDiagonalRightUp(x, y, lines) {
				totalXMASItems++
			}

			if isXMASDiagonalLeftDown(x, y, lines) {
				totalXMASItems++
			}

			if isXMASDiagonalRightDown(x, y, lines) {
				totalXMASItems++
			}
		}
	}

	fmt.Println("Total:", totalXMASItems)

	// Part 2
	totalXMAS := 0
	for y := 0; y < len(lines)-2; y++ {
		for x := 0; x < len(lines[0])-2; x++ {
			if isValidXMAS(x, y, lines) {
				totalXMAS++
			}
		}
	}
	fmt.Println("Total X-MAS patterns:", totalXMAS)

}

func isXMASForward(x int, line string) bool {
	if x+3 >= len(line) {
		return false
	}
	return line[x] == 'X' &&
		line[x+1] == 'M' &&
		line[x+2] == 'A' &&
		line[x+3] == 'S'
}

func isXMASBackward(x int, line string) bool {
	if x-3 < 0 {
		return false
	}
	return line[x] == 'X' &&
		line[x-1] == 'M' &&
		line[x-2] == 'A' &&
		line[x-3] == 'S'
}

func isXMASDown(x, y int, lines []string) bool {
	if y+3 >= len(lines) {
		return false
	}
	return lines[y][x] == 'X' &&
		lines[y+1][x] == 'M' &&
		lines[y+2][x] == 'A' &&
		lines[y+3][x] == 'S'
}

func isXMASUp(x, y int, lines []string) bool {
	if y-3 < 0 {
		return false
	}
	return lines[y][x] == 'X' &&
		lines[y-1][x] == 'M' &&
		lines[y-2][x] == 'A' &&
		lines[y-3][x] == 'S'
}

func isXMASDiagonalLeftUp(x, y int, lines []string) bool {
	if y-3 < 0 || x-3 < 0 {
		return false
	}
	return lines[y][x] == 'X' &&
		lines[y-1][x-1] == 'M' &&
		lines[y-2][x-2] == 'A' &&
		lines[y-3][x-3] == 'S'
}

func isXMASDiagonalRightUp(x, y int, lines []string) bool {
	if y-3 < 0 || x+3 >= len(lines[0]) {
		return false
	}
	return lines[y][x] == 'X' &&
		lines[y-1][x+1] == 'M' &&
		lines[y-2][x+2] == 'A' &&
		lines[y-3][x+3] == 'S'
}

func isXMASDiagonalLeftDown(x, y int, lines []string) bool {
	if y+3 >= len(lines) || x-3 < 0 {
		return false
	}
	return lines[y][x] == 'X' &&
		lines[y+1][x-1] == 'M' &&
		lines[y+2][x-2] == 'A' &&
		lines[y+3][x-3] == 'S'
}

func isXMASDiagonalRightDown(x, y int, lines []string) bool {
	if y+3 >= len(lines) || x+3 >= len(lines[0]) {
		return false
	}
	return lines[y][x] == 'X' &&
		lines[y+1][x+1] == 'M' &&
		lines[y+2][x+2] == 'A' &&
		lines[y+3][x+3] == 'S'
}

// Part 2
func isValidXMAS(x, y int, lines []string) bool {
	if (isMASDiagonalRightDown(x, y, lines) && isMASDiagonalLeftDown(x, y, lines)) ||
		(isMASDiagonalRightDownReverse(x, y, lines) && isMASDiagonalLeftDown(x, y, lines)) ||
		(isMASDiagonalRightDown(x, y, lines) && isMASDiagonalLeftDownReverse(x, y, lines)) ||
		(isMASDiagonalRightDownReverse(x, y, lines) && isMASDiagonalLeftDownReverse(x, y, lines)) {
		return true
	}

	// Bottom-to-top patterns
	if (isMASDiagonalRightUp(x, y, lines) && isMASDiagonalLeftUp(x, y, lines)) ||
		(isMASDiagonalRightUpReverse(x, y, lines) && isMASDiagonalLeftUp(x, y, lines)) ||
		(isMASDiagonalRightUp(x, y, lines) && isMASDiagonalLeftUpReverse(x, y, lines)) ||
		(isMASDiagonalRightUpReverse(x, y, lines) && isMASDiagonalLeftUpReverse(x, y, lines)) {
		return true
	}

	return false
}

func isMASDiagonalRightDown(x, y int, lines []string) bool {
	return lines[y][x] == 'M' &&
		lines[y+1][x+1] == 'A' &&
		lines[y+2][x+2] == 'S'
}

func isMASDiagonalRightDownReverse(x, y int, lines []string) bool {
	return lines[y+2][x+2] == 'M' &&
		lines[y+1][x+1] == 'A' &&
		lines[y][x] == 'S'
}

func isMASDiagonalLeftDown(x, y int, lines []string) bool {
	return lines[y][x+2] == 'M' &&
		lines[y+1][x+1] == 'A' &&
		lines[y+2][x] == 'S'
}

func isMASDiagonalLeftDownReverse(x, y int, lines []string) bool {
	return lines[y+2][x] == 'M' &&
		lines[y+1][x+1] == 'A' &&
		lines[y][x+2] == 'S'
}

func isMASDiagonalRightUp(x, y int, lines []string) bool {
	return lines[y+2][x] == 'M' &&
		lines[y+1][x+1] == 'A' &&
		lines[y][x+2] == 'S'
}

func isMASDiagonalRightUpReverse(x, y int, lines []string) bool {
	return lines[y][x+2] == 'M' &&
		lines[y+1][x+1] == 'A' &&
		lines[y+2][x] == 'S'
}

func isMASDiagonalLeftUp(x, y int, lines []string) bool {
	return lines[y+2][x+2] == 'M' &&
		lines[y+1][x+1] == 'A' &&
		lines[y][x] == 'S'
}

func isMASDiagonalLeftUpReverse(x, y int, lines []string) bool {
	return lines[y][x] == 'M' &&
		lines[y+1][x+1] == 'A' &&
		lines[y+2][x+2] == 'S'
}
