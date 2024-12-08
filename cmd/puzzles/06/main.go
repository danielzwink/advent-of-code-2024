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
	area, bounds, guard := readArea("06/input")

	moves := [4]*direction{newDirection(0, -1, '^'), newDirection(1, 0, '>'), newDirection(0, 1, 'v'), newDirection(-1, 0, '<')}
	move, movesIdx := moves[0], 0

	sum := 1
	for {
		target := guard.Add(move.coordinate)

		if target.OutOf(bounds) {
			break
		} else if area[target.Y][target.X] == '#' {
			move, movesIdx = changeDirection(movesIdx, moves)
			area[guard.Y][guard.X] = move.symbol
			continue
		} else if area[target.Y][target.X] == '.' {
			sum++
		}

		guard = target
		area[guard.Y][guard.X] = move.symbol
	}
	return sum
}

func part2() int {
	area, bounds, guard := readArea2("06/input")

	moves := [4]*direction{newDirection(0, -1, '^'), newDirection(1, 0, '>'), newDirection(0, 1, 'v'), newDirection(-1, 0, '<')}
	move, movesIdx := moves[0], 0

	sum := 0
	for {
		target := guard.Add(move.coordinate)

		if target.OutOf(bounds) {
			break
		} else if area[target.Y][target.X][0] == '#' {
			move, movesIdx = changeDirection(movesIdx, moves)
			addSymbol(area, guard, move)
			continue
		}

		potentialMove, _ := changeDirection(movesIdx, moves)
		if isInfiniteLoop(guard, potentialMove, area, bounds) {
			sum++
		}

		guard = target
		addSymbol(area, guard, move)
	}
	return sum
}

func addSymbol(area [][][]rune, guard *types.Coordinate, move *direction) {
	symbols := area[guard.Y][guard.X]

	if len(symbols) == 1 && symbols[0] == '.' {
		symbols[0] = move.symbol
	} else {
		area[guard.Y][guard.X] = append(symbols, move.symbol)
	}
}

func isInfiniteLoop(guard *types.Coordinate, move *direction, area [][][]rune, bounds *types.Coordinate) bool {
	for {
		guard = guard.Add(move.coordinate)

		if guard.OutOf(bounds) || area[guard.Y][guard.X][0] == '#' {
			return false
		}
		for _, c := range area[guard.Y][guard.X] {
			if c == move.symbol {
				return true
			}
		}
	}
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

func printArea(area [][]rune) {
	for _, row := range area {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
	fmt.Println()
}

func printArea2(area [][][]rune) {
	for _, row := range area {
		for _, c := range row {
			fmt.Printf("%c", c[len(c)-1])
		}
		fmt.Println()
	}
	fmt.Println()
}

func readArea(day string) ([][]rune, *types.Coordinate, *types.Coordinate) {
	lines := util.ReadFile(day)

	area := make([][]rune, len(lines))
	var guard *types.Coordinate
	for y, line := range lines {
		area[y] = []rune(line)

		for x, c := range line {
			if c == '^' {
				guard = types.NewCoordinate(x, y)
			}
		}
	}

	bounds := types.NewCoordinate(len(area[0]), len(area))
	return area, bounds, guard
}

func readArea2(day string) ([][][]rune, *types.Coordinate, *types.Coordinate) {
	lines := util.ReadFile(day)

	area := make([][][]rune, len(lines))
	var guard *types.Coordinate
	for y, line := range lines {
		area[y] = make([][]rune, len(line))

		for x, c := range line {
			area[y][x] = make([]rune, 1)
			area[y][x][0] = c

			if c == '^' {
				guard = types.NewCoordinate(x, y)
			}
		}
	}

	bounds := types.NewCoordinate(len(area[0]), len(area))
	return area, bounds, guard
}
