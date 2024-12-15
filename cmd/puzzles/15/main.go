package main

import (
	"advent-of-code-2024/pkg/types"
	"advent-of-code-2024/pkg/util"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

const (
	Box   rune = 'O'
	Free  rune = '.'
	Robot rune = '@'
	Wall  rune = '#'
)

func part1() int {
	warehouse, movements := readWarehouseAndMovements("15/input")
	robot := findRobot(warehouse)

	for _, movement := range movements {
		robot = perform(movement, robot, warehouse)
	}
	return sumBoxGPSCoordinates(warehouse)
}

func part2() int {
	return 0
}

func perform(movement rune, robot *types.Coordinate, warehouse [][]rune) *types.Coordinate {
	move := types.DirectionSigns[movement]
	steps, boxes := evaluate(move, robot, warehouse)

	if steps == 0 {
		return robot
	}

	warehouse[robot.Y][robot.X] = Free

	if steps == 1 && boxes == 0 {
		robot = robot.Add(move)
		warehouse[robot.Y][robot.X] = Robot
		return robot
	}

	robot = robot.Add(move)
	warehouse[robot.Y][robot.X] = Robot

	step := robot
	for i := 1; i <= boxes; i++ {
		step = step.Add(move)
		warehouse[step.Y][step.X] = Box
	}
	return robot
}

func evaluate(move *types.Coordinate, robot *types.Coordinate, warehouse [][]rune) (int, int) {
	boxes := 0

	for {
		robot = robot.Add(move)
		c := warehouse[robot.Y][robot.X]

		switch c {
		case Box:
			boxes++
			continue
		case Free:
			return 1, boxes
		case Wall:
			return 0, boxes
		}
	}
}

func sumBoxGPSCoordinates(warehouse [][]rune) int {
	sum := 0
	for y, line := range warehouse {
		for x, c := range line {
			if c == Box {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func findRobot(warehouse [][]rune) *types.Coordinate {
	for y, line := range warehouse {
		for x, c := range line {
			if c == Robot {
				return types.NewCoordinate(x, y)
			}
		}
	}
	return nil
}

func readWarehouseAndMovements(day string) ([][]rune, string) {
	lines := util.ReadFile(day)
	separator := inputSeparator(lines)

	warehouse := make([][]rune, separator)
	for i, line := range lines[0:separator] {
		warehouse[i] = []rune(line)
	}

	var movements strings.Builder
	for _, line := range lines[separator+1:] {
		movements.WriteString(line)
	}

	return warehouse, movements.String()
}

func inputSeparator(lines []string) int {
	for i, line := range lines {
		if len(line) == 0 {
			return i
		}
	}
	return -1
}
