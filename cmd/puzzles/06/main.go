package main

import (
	"advent-of-code-2024/pkg/types"
	"advent-of-code-2024/pkg/util"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	area := readArea("06/input")
	bound := types.NewCoordinate(len(area[0]), len(area))
	guard := findGuardPosition(area)

	moves := [4]*direction{newDirection(0, -1, '^'), newDirection(1, 0, '>'), newDirection(0, 1, 'v'), newDirection(-1, 0, '<')}
	move, movesIdx := moves[0], 0

	sum := 1
	for {
		target := guard.Add(move.coordinate)
		if target.OutOf(bound) {
			break
		}

		if area[target.Y][target.X] == '#' {
			move, movesIdx = changeDirection(movesIdx, moves)
		} else {
			guard = guard.Add(move.coordinate)
			if area[guard.Y][guard.X] == '.' {
				sum++
			}
			area[guard.Y][guard.X] = move.symbol
		}
	}
	return sum
}

func part2() int {
	area := readArea("06/input")
	bound := types.NewCoordinate(len(area[0]), len(area))
	guard := findGuardPosition(area)

	moves := [4]*direction{newDirection(0, -1, '^'), newDirection(1, 0, '>'), newDirection(0, 1, 'v'), newDirection(-1, 0, '<')}
	move, movesIdx := moves[0], 0

	sum := 0
	for {
		target := guard.Add(move.coordinate)
		if target.OutOf(bound) {
			break
		}

		potentialMove, _ := changeDirection(movesIdx, moves)
		potentialTarget := guard
		for {
			potentialTarget = potentialTarget.Add(potentialMove.coordinate)

			if potentialTarget.OutOf(bound) || area[potentialTarget.Y][potentialTarget.X] == '#' {
				break
			}
			if area[potentialTarget.Y][potentialTarget.X] == potentialMove.symbol {
				sum++
				break
			}
		}

		if area[target.Y][target.X] == '#' {
			move, movesIdx = changeDirection(movesIdx, moves)
		} else {
			guard = guard.Add(move.coordinate)
			area[guard.Y][guard.X] = move.symbol
		}
	}
	return sum
}

type direction struct {
	coordinate *types.Coordinate
	symbol     rune
}

func newDirection(x, y int, symbol rune) *direction {
	return &direction{coordinate: &types.Coordinate{X: x, Y: y}, symbol: symbol}
}

func changeDirection(idx int, moves [4]*direction) (*direction, int) {
	idx++
	idx = idx % len(moves)
	return moves[idx], idx
}

func findGuardPosition(area [][]rune) *types.Coordinate {
	for y, row := range area {
		for x, c := range row {
			if c == '^' {
				return types.NewCoordinate(x, y)
			}
		}
	}
	return nil
}

func printArea(area [][]rune) {
	for _, row := range area {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
	fmt.Println()
}

func readArea(day string) [][]rune {
	lines := util.ReadFile(day)

	area := make([][]rune, len(lines))
	for i, line := range lines {
		area[i] = []rune(line)
	}

	return area
}
