package main

import (
	"advent-of-code-2024/pkg/util"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	reports := readReports("02/input")

	safeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}
	return safeReports
}

func part2() int {
	reports := readReports("02/input")

	safeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		} else {
			for i := 0; i < len(report); i++ {
				newReport := make([]int, 0)
				newReport = append(newReport, report[:i]...)
				newReport = append(newReport, report[i+1:]...)

				if isSafe(newReport) {
					safeReports++
					break
				}
			}
		}
	}
	return safeReports
}

func isSafe(report []int) bool {
	first := report[0]
	second := report[1]

	if first == second {
		return false
	}

	increasing := first < second
	decreasing := first > second

	for i := 1; i < len(report); i++ {
		previous := report[i-1]
		next := report[i]

		if increasing && (previous >= next || next-previous > 3) {
			return false
		}
		if decreasing && (previous <= next || previous-next > 3) {
			return false
		}
	}
	return true
}

func readReports(day string) [][]int {
	lines := util.ReadFile(day)

	reports := make([][]int, 0)
	for _, line := range lines {
		levels := strings.Split(line, " ")

		report := make([]int, 0)
		for _, level := range levels {
			report = append(report, util.MustParseInt(level))
		}
		reports = append(reports, report)
	}
	return reports
}
