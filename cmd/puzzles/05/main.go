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
	rules, updates := readRulesAndUpdates("05/input")

	orderedUpdates := make([][]int, 0)
	for _, update := range updates {
		if isUpdateCorrectlyOrdered(update, rules) {
			orderedUpdates = append(orderedUpdates, update)
		}
	}

	sum := 0
	for _, update := range orderedUpdates {
		sum += getMiddlePageNumber(update)
	}
	return sum
}

func part2() int {
	rules, updates := readRulesAndUpdates("05/input")

	unorderedUpdates := make([][]int, 0)
	for _, update := range updates {
		if !isUpdateCorrectlyOrdered(update, rules) {
			unorderedUpdates = append(unorderedUpdates, update)
		}
	}

	for _, update := range unorderedUpdates {
		for fixSingleViolation(update, rules) {
			// until fixed
		}
	}

	sum := 0
	for _, update := range unorderedUpdates {
		sum += getMiddlePageNumber(update)
	}
	return sum
}

func isUpdateCorrectlyOrdered(update []int, rules map[int]map[int]struct{}) bool {
	for i, number := range update[1:] {
		successors, exist := rules[number]
		if !exist {
			continue
		}

		for _, pre := range update[0 : i+1] {
			_, violation := successors[pre]
			if violation {
				return false
			}
		}
	}
	return true
}

func fixSingleViolation(update []int, rules map[int]map[int]struct{}) bool {
	for i, number := range update[1:] {
		successors, exist := rules[number]
		if !exist {
			continue
		}

		for k, pre := range update[0 : i+1] {
			_, violation := successors[pre]
			if violation {
				update[i+1] = pre
				update[k] = number
				return true
			}
		}
	}
	return false
}

func getMiddlePageNumber(update []int) int {
	i := (len(update) - 1) / 2
	return update[i]
}

func readRulesAndUpdates(day string) (map[int]map[int]struct{}, [][]int) {
	lines := util.ReadFile(day)
	separator := getSeparator(lines)

	rules := make(map[int]map[int]struct{})
	for _, line := range lines[0:separator] {
		pre, suc := splitRule(line)
		successors, exist := rules[pre]
		if !exist {
			successors = make(map[int]struct{})
		}
		successors[suc] = struct{}{}
		rules[pre] = successors
	}

	updates := make([][]int, 0)
	for _, line := range lines[separator+1:] {
		update := splitUpdate(line)
		updates = append(updates, update)
	}

	return rules, updates
}

func getSeparator(lines []string) int {
	for i, line := range lines {
		if len(line) == 0 {
			return i
		}
	}
	return -1
}

func splitRule(line string) (int, int) {
	splits := strings.Split(line, "|")
	return util.MustParseInt(splits[0]), util.MustParseInt(splits[1])
}

func splitUpdate(line string) []int {
	splits := strings.Split(line, ",")
	update := make([]int, 0)
	for _, split := range splits {
		update = append(update, util.MustParseInt(split))
	}
	return update
}
