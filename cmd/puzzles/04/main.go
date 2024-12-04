package main

import (
	"advent-of-code-2024/pkg/util"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	search := readWordSearch("04/input")

	yBound := len(search)
	xBound := len(search[0])

	sum := 0
	for y, row := range search {
		for x, char := range row {
			if char == 'X' {
				// x+3 / y --> to the right
				if x+3 < xBound {
					if search[y][x+1] == 'M' && search[y][x+2] == 'A' && search[y][x+3] == 'S' {
						sum++
					}
				}
				// x+3 / y+3 --> to bottom right
				if x+3 < xBound && y+3 < yBound {
					if search[y+1][x+1] == 'M' && search[y+2][x+2] == 'A' && search[y+3][x+3] == 'S' {
						sum++
					}
				}
				// x / y+3 --> to the bottom
				if y+3 < yBound {
					if search[y+1][x] == 'M' && search[y+2][x] == 'A' && search[y+3][x] == 'S' {
						sum++
					}
				}
				// x-3 / y+3 --> to bottom left
				if x-3 >= 0 && y+3 < yBound {
					if search[y+1][x-1] == 'M' && search[y+2][x-2] == 'A' && search[y+3][x-3] == 'S' {
						sum++
					}
				}
				// x-3 / y --> to the left
				if x-3 >= 0 {
					if search[y][x-1] == 'M' && search[y][x-2] == 'A' && search[y][x-3] == 'S' {
						sum++
					}
				}
				// x-3 / y-3 --> to left top
				if x-3 >= 0 && y-3 >= 0 {
					if search[y-1][x-1] == 'M' && search[y-2][x-2] == 'A' && search[y-3][x-3] == 'S' {
						sum++
					}
				}
				// x / y-3 --> to the top
				if y-3 >= 0 {
					if search[y-1][x] == 'M' && search[y-2][x] == 'A' && search[y-3][x] == 'S' {
						sum++
					}
				}
				// x+3 / y-3 --> to top right
				if x+3 < xBound && y-3 >= 0 {
					if search[y-1][x+1] == 'M' && search[y-2][x+2] == 'A' && search[y-3][x+3] == 'S' {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func part2() int {
	search := readWordSearch("04/input")

	yBound := len(search)
	xBound := len(search[0])

	sum := 0
	for y, row := range search {
		for x, char := range row {
			if char == 'A' {
				if x-1 >= 0 && y-1 >= 0 && x+1 < xBound && y+1 < yBound {
					topLeft := search[y-1][x-1]
					topRight := search[y-1][x+1]
					bottomLeft := search[y+1][x-1]
					bottomRight := search[y+1][x+1]

					if topLeft == 'M' && bottomRight == 'S' && topRight == 'M' && bottomLeft == 'S' {
						sum++
					}
					if bottomLeft == 'M' && topRight == 'S' && bottomRight == 'M' && topLeft == 'S' {
						sum++
					}
					if topLeft == 'M' && bottomRight == 'S' && bottomLeft == 'M' && topRight == 'S' {
						sum++
					}
					if topRight == 'M' && bottomLeft == 'S' && bottomRight == 'M' && topLeft == 'S' {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func readWordSearch(day string) [][]rune {
	lines := util.ReadFile(day)

	matrix := make([][]rune, len(lines))
	for y, line := range lines {
		row := make([]rune, len(line))
		for x, char := range line {
			row[x] = char
		}
		matrix[y] = row
	}
	return matrix
}
