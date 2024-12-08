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
	antennas, bounds := readAntennas("08/input")

	uniqueAntinodes := make(map[string]struct{})
	for _, coordinates := range antennas {
		for n, c1 := range coordinates {
			for _, c2 := range coordinates[n+1:] {
				if c1 != c2 {
					a1 := c1.Add(c1.Diff(c2))
					if a1.Within(bounds) {
						uniqueAntinodes[a1.Key()] = struct{}{}
					}

					a2 := c2.Add(c2.Diff(c1))
					if a2.Within(bounds) {
						uniqueAntinodes[a2.Key()] = struct{}{}
					}
				}
			}
		}
	}
	return len(uniqueAntinodes)
}

func part2() int {
	antennas, bounds := readAntennas("08/input")

	uniqueAntinodes := make(map[string]struct{})
	for _, coordinates := range antennas {
		for n, c1 := range coordinates {
			for _, c2 := range coordinates[n+1:] {
				if c1 != c2 {
					a1, diff1 := c1, c1.Diff(c2)
					for {
						a1 = a1.Add(diff1)
						if a1.Within(bounds) {
							uniqueAntinodes[a1.Key()] = struct{}{}
						} else {
							break
						}
					}

					a2, diff2 := c2, c2.Diff(c1)
					for {
						a2 = a2.Add(diff2)
						if a2.Within(bounds) {
							uniqueAntinodes[a2.Key()] = struct{}{}
						} else {
							break
						}
					}
				}
			}
		}
	}
	for _, coordinates := range antennas {
		for _, c := range coordinates {
			uniqueAntinodes[c.Key()] = struct{}{}
		}
	}
	return len(uniqueAntinodes)
}

func readAntennas(day string) (map[rune][]*types.Coordinate, *types.Coordinate) {
	lines := util.ReadFile(day)

	antennas := make(map[rune][]*types.Coordinate)
	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				antenna, exists := antennas[c]
				if !exists {
					antenna = make([]*types.Coordinate, 0)
				}
				antenna = append(antenna, types.NewCoordinate(x, y))
				antennas[c] = antenna
			}
		}
	}

	bounds := types.NewCoordinate(len(lines[0]), len(lines))
	return antennas, bounds
}
