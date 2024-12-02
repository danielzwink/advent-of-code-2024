package main

import (
	"advent-of-code-2024/pkg/util"
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	locationIDs1, locationIDs2 := readLocationIDs("01/input")

	sort.Ints(locationIDs1)
	sort.Ints(locationIDs2)

	totalDistance := 0
	for i := 0; i < len(locationIDs1); i++ {
		distance := util.Abs(locationIDs1[i] - locationIDs2[i])
		totalDistance += distance
	}

	return totalDistance
}

func part2() int {
	locationIDs1, locationIDs2 := readLocationIDs("01/input")

	occurrences := make(map[int]int)
	for _, locationID := range locationIDs2 {
		value, exists := occurrences[locationID]

		if exists {
			value++
		} else {
			value = 1
		}

		occurrences[locationID] = value
	}

	similarityScore := 0
	for _, locationID := range locationIDs1 {
		value, exists := occurrences[locationID]

		if exists {
			similarityScore += value * locationID
		}
	}
	return similarityScore
}

func readLocationIDs(day string) ([]int, []int) {
	lines := util.ReadFile(day)
	locationIDs1 := make([]int, 0)
	locationIDs2 := make([]int, 0)

	for _, line := range lines {
		row := strings.Split(line, "   ")
		locationID1 := util.MustParseInt(row[0])
		locationID2 := util.MustParseInt(row[1])

		locationIDs1 = append(locationIDs1, locationID1)
		locationIDs2 = append(locationIDs2, locationID2)
	}
	return locationIDs1, locationIDs2
}
