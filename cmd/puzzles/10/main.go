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

var moves = []*types.Coordinate{
	types.NewCoordinate(0, -1), // up
	types.NewCoordinate(1, 0),  // right
	types.NewCoordinate(0, 1),  // down
	types.NewCoordinate(-1, 0)} // left

func part1() int {
	trailMap, trailBounds, trailHeads := readTrailMap("10/input")

	score := 0
	for _, trailHead := range trailHeads {
		result := findTrailEnds(trailHead, 0, trailMap, trailBounds)

		unique := make(map[string]struct{})
		for _, c := range result {
			unique[c.Key()] = struct{}{}
		}
		score += len(unique)
	}
	return score
}

func part2() int {
	trailMap, trailBounds, trailHeads := readTrailMap("10/input")

	score := 0
	for _, trailHead := range trailHeads {
		result := findTrailEnds(trailHead, 0, trailMap, trailBounds)
		score += len(result)
	}
	return score
}

func findTrailEnds(currentPosition *types.Coordinate, currentHeight int, trailMap [][]int, trailBounds *types.Coordinate) []*types.Coordinate {
	trailEnds := make([]*types.Coordinate, 0)

	for _, move := range moves {
		next := currentPosition.Add(move)

		if next.Within(trailBounds) {
			nextHeight := trailMap[next.Y][next.X]

			if nextHeight != currentHeight+1 {
				continue
			}

			if nextHeight == 9 {
				trailEnds = append(trailEnds, next)
			} else {
				furtherTrailEnds := findTrailEnds(next, nextHeight, trailMap, trailBounds)
				trailEnds = append(trailEnds, furtherTrailEnds...)
			}

		}
	}
	return trailEnds
}

func readTrailMap(day string) ([][]int, *types.Coordinate, []*types.Coordinate) {
	lines := util.ReadFile(day)

	trailMap := make([][]int, len(lines))
	trailBounds := types.NewCoordinate(len(lines), len(lines[0]))
	trailHeads := make([]*types.Coordinate, 0)
	for y, line := range lines {
		trailMap[y] = make([]int, len(line))
		for x, r := range line {
			trailMap[y][x] = util.MustParseInt(string(r))

			if trailMap[y][x] == 0 {
				trailHeads = append(trailHeads, types.NewCoordinate(x, y))
			}
		}
	}
	return trailMap, trailBounds, trailHeads
}
