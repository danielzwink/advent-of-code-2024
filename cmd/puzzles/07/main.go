package main

import (
	"advent-of-code-2024/pkg/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	equations := readEquations("07/input")
	availableOperators := []rune{'+', '*'}

	sum := 0
	for _, e := range equations {
		result := e[0]
		numbers := e[1:]
		requiredOperations := len(numbers) - 1
		permutations := createOperatorPermutations(requiredOperations, availableOperators)

		if isEquationPossible(result, numbers, permutations) {
			sum += result
		}
	}
	return sum
}

func part2() int {
	equations := readEquations("07/input")
	availableOperators := []rune{'+', '*', '|'}

	sum := 0
	for _, e := range equations {
		result := e[0]
		numbers := e[1:]
		requiredOperations := len(numbers) - 1
		permutations := createOperatorPermutations(requiredOperations, availableOperators)

		if isEquationPossible(result, numbers, permutations) {
			sum += result
		}
	}
	return sum
}

func createOperatorPermutations(requiredOperations int, availableOperators []rune) [][]rune {
	operatorsCount := len(availableOperators)
	rows := exp(operatorsCount, requiredOperations)

	permutations := make([][]rune, rows)
	for p := 0; p < rows; p++ {
		permutations[p] = make([]rune, requiredOperations)
	}

	for x := 0; x < requiredOperations; x++ {
		for y := 0; y < rows; y++ {
			o := (y / exp(operatorsCount, x)) % operatorsCount
			permutations[y][x] = availableOperators[o]
		}
	}
	return permutations
}

func isEquationPossible(expectedResult int, numbers []int, permutations [][]rune) bool {
	for _, operations := range permutations {
		result := numbers[0]

		for n, operator := range operations {
			if operator == '+' {
				result += numbers[n+1]
			} else if operator == '*' {
				result *= numbers[n+1]
			} else if operator == '|' {
				result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(numbers[n+1]))
			} else {
				panic(1)
			}
		}

		if result == expectedResult {
			return true
		}
	}
	return false
}

func exp(b, e int) int {
	result := 1
	for i := 0; i < e; i++ {
		result *= b
	}
	return result
}

func readEquations(day string) [][]int {
	lines := util.ReadFile(day)

	equations := make([][]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		equation := make([]int, 1)
		equation[0] = util.MustParseInt(parts[0])
		numbers := strings.Split(parts[1], " ")
		for _, number := range numbers {
			equation = append(equation, util.MustParseInt(number))
		}
		equations = append(equations, equation)
	}
	return equations
}
