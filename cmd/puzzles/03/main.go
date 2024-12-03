package main

import (
	"advent-of-code-2024/pkg/util"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

type multiplication struct {
	factor1 int
	factor2 int
}

func newMultiplication(factor1, factor2 string) multiplication {
	return multiplication{
		factor1: util.MustParseInt(factor1),
		factor2: util.MustParseInt(factor2),
	}
}

func (m multiplication) result() int {
	return m.factor1 * m.factor2
}

func part1() int {
	multiplications := readMultiplicationsP1("03/input")

	sum := 0
	for _, m := range multiplications {
		sum += m.result()
	}
	return sum
}

func part2() int {
	multiplications := readMultiplicationsP2("03/input")

	sum := 0
	for _, m := range multiplications {
		sum += m.result()
	}
	return sum
}

func readMultiplicationsP1(day string) []multiplication {
	line := util.ReadFile(day)[0]
	return readMultiplications(line)
}

func readMultiplicationsP2(day string) []multiplication {
	line := util.ReadFile(day)[0]

	validLines := make([]string, 0)
	doSplits := strings.Split(line, "do()")
	for _, do := range doSplits {
		dontSplits := strings.Split(do, "don't()")
		validLines = append(validLines, dontSplits[0])
	}
	return readMultiplications(validLines...)
}

func readMultiplications(lines ...string) []multiplication {
	matcher := regexp.MustCompile("mul\\(([0-9]+),([0-9]+)\\)")

	multiplications := make([]multiplication, 0)
	for _, line := range lines {
		matches := matcher.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			multiplications = append(multiplications, newMultiplication(match[1], match[2]))
		}
	}
	return multiplications
}
