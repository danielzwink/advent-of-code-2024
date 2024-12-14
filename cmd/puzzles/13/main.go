package main

import (
	"advent-of-code-2024/pkg/util"
	"fmt"
	"regexp"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	less := readLinearEquationSystems("13/input")

	tokens := 0
	for _, les := range less {
		tokens += les.solve()
	}
	return tokens
}

func part2() int {
	less := readLinearEquationSystems("13/input")

	tokens := 0
	for _, les := range less {
		les.e1.result += 10000000000000
		les.e2.result += 10000000000000

		tokens += les.solve()
	}
	return tokens
}

type LinearEquationSystem struct {
	e1, e2 *Equation
}

func (l *LinearEquationSystem) solve() int {
	a := (l.e1.result*l.e2.b - l.e2.result*l.e1.b) / (l.e1.a*l.e2.b - l.e2.a*l.e1.b)
	b := (l.e1.result*l.e2.a - l.e2.result*l.e1.a) / (l.e1.b*l.e2.a - l.e2.b*l.e1.a)

	if l.e1.a*a+l.e1.b*b == l.e1.result && l.e2.a*a+l.e2.b*b == l.e2.result {
		return 3*a + b
	} else {
		return 0
	}
}

type Equation struct {
	a, b, result int
}

func readLinearEquationSystems(day string) []*LinearEquationSystem {
	lines := util.ReadFile(day)

	buttonARegExp, _ := regexp.Compile("Button A. X.([0-9]+), Y.([0-9]+)")
	buttonBRegExp, _ := regexp.Compile("Button B. X.([0-9]+), Y.([0-9]+)")
	priceRegExp, _ := regexp.Compile("Prize. X=([0-9]+), Y=([0-9]+)")

	less := make([]*LinearEquationSystem, 0)
	for i := 0; i < len(lines); i += 4 {
		buttonA := buttonARegExp.FindStringSubmatch(lines[i])
		buttonB := buttonBRegExp.FindStringSubmatch(lines[i+1])
		price := priceRegExp.FindStringSubmatch(lines[i+2])

		les := &LinearEquationSystem{
			e1: &Equation{
				a:      util.MustParseInt(buttonA[1]),
				b:      util.MustParseInt(buttonB[1]),
				result: util.MustParseInt(price[1])},
			e2: &Equation{
				a:      util.MustParseInt(buttonA[2]),
				b:      util.MustParseInt(buttonB[2]),
				result: util.MustParseInt(price[2])},
		}
		less = append(less, les)
	}
	return less
}
