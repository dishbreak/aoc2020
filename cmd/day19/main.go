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

// rule is a structure we can use to hold the parsed rule from the input.
type rule struct {
	number        int
	baseCondition byte
	conditionals  [][]int
}

type ruleset map[int]rule

// newRule will parse the string from the input file and create a rule struct.
func newRule(statement string) rule {
	// break the string into fields delimited by whitespace.
	parts := strings.Fields(statement)

	// remove the trailing colon from the first field, then parse the rule ID as
	// a number.
	parts[0] = strings.TrimSuffix(parts[0], ":")
	number, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	// if the second field starts with a quote mark, this rule matches a single
	// character.
	// return a rule wih a baseCondition and number set.
	if strings.HasPrefix(parts[1], `"`) {
		return rule{
			number:        number,
			baseCondition: parts[1][1],
		}
	}

	// otherwise, let's start making a composite rule.
	// create a slice to hold all our slices, and a slice to hold the first subrule.
	conditionals := make([][]int, 0)
	condition := make([]int, 0)
	for _, item := range parts[1:] {
		switch item {
		// if we encounter a pipe, add the existing conditional to the slice
		// of slices and create a new slice.
		case "|":
			conditionals = append(conditionals, condition)
			condition = make([]int, 0)
		// Otherwise, parse the field as a number and add it to the existing conditional.
		default:
			number, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			condition = append(condition, number)
		}
	}
	// add the last working conditional to the conditionals slice.
	conditionals = append(conditionals, condition)

	// return a composite rule
	return rule{
		number:        number,
		conditionals:  conditionals,
		baseCondition: '\x00',
	}
}

func newRuleset(rules []string) ruleset {
	result := make(map[int]rule)
	for _, ruleLine := range rules {
		parsedRule := newRule(ruleLine)
		result[parsedRule.number] = parsedRule
	}
	return result
}

// evaluateRule will return the number of characters matched or -1 if the string
// does not match the rule.
// The function gets called recursively, invoking subrules as needed.
// When multiple rules match, we'll pick the greatest number of matched characters.
func (r ruleset) evaluateRule(message string, ruleIdx, position int) int {
	myRule := r[ruleIdx]

	// if the selected rule has a base condition, evaluate it and return immediately.
	if myRule.baseCondition != '\x00' {
		// if we've run out of characters OR the character at the given position
		// doesn't match, return -1
		if position > len(message)-1 || message[position] != myRule.baseCondition {
			return -1
		}
		// otherwise, return the number of characters now matched.
		return position + 1
	}

	// otherwise, we're going to evaluate our subrules.
	// set aside matched to keep track of the subrule with the most matches.
	matched := -1
	// we model a composite rule as an array of conditional, where each
	// conditional names a sequence of provisions that must match.
	for _, conditional := range myRule.conditionals {
		// we're going to assume that all the previous characters before
		// position have been scanned.
		scanned := position
		for _, provision := range conditional {
			// evaluate provisions one by one, using the scanned count from the
			// previous invocation to determine where the next provision will
			// start checking.
			scanned = r.evaluateRule(message, provision, scanned)
			// if a provision doesn't match, there's no point in continuing to
			// evaluate the conditional.
			if scanned == -1 {
				break
			}
		}
		// update matched to be the total scanned if this conditional scanned
		// more characters.
		if matched < scanned {
			matched = scanned
		}
	}
	// return the number of characters matched.
	return matched
}

// countMatches counts the number of times a given rule matches the input and
// then returns the match count and the next index in the string.
func (r ruleset) countMatches(message string, ruleNum, position int) (matches, matchedChars int) {
	scanned := position
	matchedChars = 0
	matches = 0

	for scanned != -1 {
		scanned = r.evaluateRule(message, ruleNum, scanned)
		// if the rule doesn't match, we've found all the matches we will find
		// for this rule.
		if scanned == -1 {
			break
		}
		matchedChars = scanned
		matches++
	}
	return
}

// validate ensures that rule 0 matches all characters in the message.
// this will handle the corner case where a message matches all rules but has
// additional characters on the end.
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
	r := newRuleset(input[0])
	result := 0

	// by doing substitution, we can understand the following with our recursive
	// rules:
	// * rule 8 will become a chain of invocations of rule 42:
	//   8: 42 | 42 42 | 42 42 42 | ...
	// * rule 11 will become a chain of invocations of rule 42, followed by the
	// **same number of invocations of rule 31**
	//   11: 42 31 | 42 42 31 31 | 42 42 42 31 31 31 | ...
	// because rule 0 is rule 8 followed by rule 11, we can sidestep rule 0 and
	// count the number of times that rule 42 matches, then the number of times
	// rule 31 matches.
	// the message is valid if the following conditions hold:
	// * rule 42 matches m times, where m >= 2
	// * the remaining string matches rule 31 n times, where n >=1
	// * all characters match
	// * m - n  >= 1
	for _, message := range input[1] {
		rule42matches, position := r.countMatches(message, 42, 0)
		if rule42matches < 2 {
			continue
		}
		rule31matches, position := r.countMatches(message, 31, position)
		if position != len(message) || rule31matches == 0 || rule42matches-rule31matches < 1 {
			continue
		}
		result++
	}
	return result
}
