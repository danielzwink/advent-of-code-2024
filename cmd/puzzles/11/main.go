package main

import (
	"advent-of-code-2024/pkg/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	stones := readStones("11/input")
	return blink(25, stones)
}

func part2() int {
	stones := readStones("11/input")
	return blink(75, stones)
}

func blink(times int, stones map[int]int) int {
	for i := 0; i < times; i++ {
		temp := make(map[int]int)
		for stone, count := range stones {
			temp[stone] = count
		}

		for stone, count := range stones {
			if count == 0 {
				continue
			}
			temp[stone] -= count
			if temp[stone] == 0 {
				delete(temp, stone)
			}

			if stone == 0 {
				temp[1] += count
				continue
			}

			stoneAsText := strconv.Itoa(stone)
			length := len(stoneAsText)
			if length%2 == 0 {
				left := stoneAsText[0 : length/2]
				right := stoneAsText[length/2:]
				temp[util.MustParseInt(left)] += count
				temp[util.MustParseInt(right)] += count
				continue
			}

			temp[(stone * 2024)] += count
		}
		stones = temp
	}

	result := 0
	for _, count := range stones {
		result += count
	}
	return result
}

func readStones(day string) map[int]int {
	lines := util.ReadFile(day)
	stonesAsText := strings.Split(lines[0], " ")

	stones := make(map[int]int)
	for _, stoneAsText := range stonesAsText {
		stone := util.MustParseInt(stoneAsText)
		stones[stone] = 1
	}
	return stones
}
