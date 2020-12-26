package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInputAsSections("inputs/day19.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type rule struct {
	number        int
	baseCondition byte
	conditionals  [][]int
}

type ruleset []rule

func newRule(statement string) rule {
	parts := strings.Fields(statement)

	parts[0] = strings.TrimSuffix(parts[0], ":")
	number, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	if strings.HasPrefix(parts[1], `"`) {
		return rule{
			number:        number,
			baseCondition: parts[1][1],
		}
	}

	conditionals := make([][]int, 0)
	condition := make([]int, 0)
	for _, item := range parts[1:] {
		switch item {
		case "|":
			conditionals = append(conditionals, condition)
			condition = make([]int, 0)
		default:
			number, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			condition = append(condition, number)
		}
	}
	conditionals = append(conditionals, condition)

	return rule{
		number:        number,
		conditionals:  conditionals,
		baseCondition: '\x00',
	}
}

func newRuleset(rules []string) ruleset {
	result := make([]rule, len(rules))
	for _, ruleLine := range rules {
		parsedRule := newRule(ruleLine)
		result[parsedRule.number] = parsedRule
	}
	return result
}

func (r ruleset) evaluateRule(message string, ruleIdx, position int) int {
	myRule := r[ruleIdx]
	if myRule.baseCondition != '\x00' {
		if position > len(message)-1 || message[position] != myRule.baseCondition {
			return -1
		}
		return position + 1
	}

	matched := -1
	for _, conditional := range myRule.conditionals {
		scanned := position
		for _, provision := range conditional {
			scanned = r.evaluateRule(message, provision, scanned)
			if scanned == -1 {
				break
			}
		}
		if matched < scanned {
			matched = scanned
		}
	}
	return matched
}

func (r ruleset) validate(message string) bool {
	return r.evaluateRule(message, 0, 0) == len(message)
}

func part1(input [][]string) int {
	r := newRuleset(input[0])

	result := 0
	for _, message := range input[1] {
		if r.validate(message) {
			result++
		}
	}

	return result
}

func part2(input [][]string) int {
	return 0
}
