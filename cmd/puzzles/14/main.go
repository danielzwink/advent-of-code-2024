package main

import (
	"advent-of-code-2024/pkg/types"
	"advent-of-code-2024/pkg/util"
	"fmt"
	"regexp"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	robots := readRobotPositions("14/input")

	bounds := types.NewCoordinate(101, 103)
	positions := moveRobots(robots, bounds, 100)
	topLeft, topRight, bottomLeft, bottomRight := distributeIntoQuadrants(positions, bounds)
	return topLeft * topRight * bottomLeft * bottomRight
}

func part2() int {
	robots := readRobotPositions("14/input")

	bounds := types.NewCoordinate(101, 103)
	seconds := 0
	for {
		seconds++
		positions := moveRobots(robots, bounds, seconds)

		distribution := make(map[string]int)
		for _, position := range positions {
			distribution[position.Key()]++
		}

		// not my solution
		if len(positions) == len(distribution) {
			//printDistribution(distribution, bounds)
			return seconds
		}
	}
}

func moveRobots(robots []robot, bounds *types.Coordinate, seconds int) []*types.Coordinate {
	positions := make([]*types.Coordinate, len(robots))

	for i, r := range robots {
		x := (r.position.X + seconds*r.velocity.X) % bounds.X
		y := (r.position.Y + seconds*r.velocity.Y) % bounds.Y

		if x < 0 {
			x = bounds.X + x
		}
		if y < 0 {
			y = bounds.Y + y
		}
		positions[i] = types.NewCoordinate(x, y)
	}
	return positions
}

func distributeIntoQuadrants(positions []*types.Coordinate, bounds *types.Coordinate) (int, int, int, int) {
	middles := types.NewCoordinate((bounds.X-1)/2, (bounds.Y-1)/2)

	topLeftLower := types.NewCoordinate(0, 0)
	topLeftUpper := types.NewCoordinate(middles.X, middles.Y)
	topRightLower := types.NewCoordinate(middles.X+1, 0)
	topRightUpper := types.NewCoordinate(bounds.X, middles.Y)
	bottomLeftLower := types.NewCoordinate(0, middles.Y+1)
	bottomLeftUpper := types.NewCoordinate(middles.X, bounds.Y)
	bottomRightLower := types.NewCoordinate(middles.X+1, middles.Y+1)
	bottomRightUpper := types.NewCoordinate(bounds.X, bounds.Y)

	topLeft, topRight, bottomLeft, bottomRight := 0, 0, 0, 0
	for _, position := range positions {
		if position.WithinBounds(topLeftLower, topLeftUpper) {
			topLeft++
		} else if position.WithinBounds(topRightLower, topRightUpper) {
			topRight++
		} else if position.WithinBounds(bottomLeftLower, bottomLeftUpper) {
			bottomLeft++
		} else if position.WithinBounds(bottomRightLower, bottomRightUpper) {
			bottomRight++
		}
	}
	return topLeft, topRight, bottomLeft, bottomRight
}

func printDistribution(distribution map[string]int, bounds *types.Coordinate) {
	for y := 0; y < bounds.Y; y++ {
		for x := 0; x < bounds.X; x++ {
			position := types.NewCoordinate(x, y)
			count := distribution[position.Key()]

			if count == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", count)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type robot struct {
	position *types.Coordinate
	velocity *types.Coordinate
}

func readRobotPositions(day string) []robot {
	lines := util.ReadFile(day)

	robots := make([]robot, len(lines))
	for i, line := range lines {
		r, _ := regexp.Compile("p=([0-9]+),([0-9]+).v=(-?[0-9]+),(-?[0-9]+)")
		m := r.FindStringSubmatch(line)

		robots[i] = robot{
			position: types.NewCoordinate(util.MustParseInt(m[1]), util.MustParseInt(m[2])),
			velocity: types.NewCoordinate(util.MustParseInt(m[3]), util.MustParseInt(m[4])),
		}
	}
	return robots
}
