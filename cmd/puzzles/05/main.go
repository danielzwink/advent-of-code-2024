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

	correctUpdates := make([][]int, 0)
	for _, update := range updates {
		if isUpdateInRightOrder(update, rules) {
			correctUpdates = append(correctUpdates, update)
		}
	}

	sum := 0
	for _, update := range correctUpdates {
		i := (len(update) - 1) / 2
		sum += update[i]
	}

	return sum
}

func part2() int {
	rules, updates := readRulesAndUpdates("05/input")

	incorrectUpdates := make([][]int, 0)
	for _, update := range updates {
		if !isUpdateInRightOrder(update, rules) {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	for _, update := range incorrectUpdates {
		for fixIncorrectUpdate(update, rules) {
		}
	}

	sum := 0
	for _, update := range incorrectUpdates {
		i := (len(update) - 1) / 2
		sum += update[i]
	}

	return sum
}

func isUpdateInRightOrder(update []int, rules map[int]map[int]struct{}) bool {
	for i, n := range update[1:] {
		rule, exists := rules[n]
		if exists {
			for _, pre := range update[0 : i+1] {
				_, violation := rule[pre]
				if violation {
					return false
				}
			}
		}
	}
	return true
}

func fixIncorrectUpdate(update []int, rules map[int]map[int]struct{}) bool {
	for i, n := range update[1:] {
		rule, exists := rules[n]
		if exists {
			for k, pre := range update[0 : i+1] {
				_, violation := rule[pre]
				if violation {
					update[i+1] = pre
					update[k] = n
					return true
				}
			}
		}
	}
	return false
}

func readRulesAndUpdates(day string) (map[int]map[int]struct{}, [][]int) {
	lines := util.ReadFile(day)

	rules := make(map[int]map[int]struct{})
	updates := make([][]int, 0)
	readRules := true

	for _, line := range lines {
		if len(line) == 0 {
			readRules = false
		} else if readRules {
			first, second := splitRule(line)
			value, exists := rules[first]
			if !exists {
				value = make(map[int]struct{})
			}
			value[second] = struct{}{}
			rules[first] = value
		} else {
			update := splitUpdate(line)
			updates = append(updates, update)
		}
	}
	return rules, updates
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
