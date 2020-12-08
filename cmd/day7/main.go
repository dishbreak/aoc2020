package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

var rulesMatcher *regexp.Regexp = regexp.MustCompile(`^([\w\s]+) bags contain (.*)$`)
var innerBagMatcher *regexp.Regexp = regexp.MustCompile(`(\d+) ([\w\s]+) bag`)

func part1(input []string) int {
	// We're going to construct a list of edges from the rule set and do a
	// depth-first-search (DFS) discovering all compliant outer bags.

	// Begin by reversing all the rules. For example,
	// "striped gold bags contain 1 dotted blue bag, 5 drab bronze bags, 2
	// mirrored orange bags, 2 shiny violet bags."
	// becomes
	// - dotted blue bag can be contained by striped gold bag
	// - drab bronze bag can be contained by striped gold bag
	// - mirrored orange bag can be contained by striped gold bag
	// - shiny violet bag can be contained by striped gold bag

	validOuterBags := make(map[string][]string)

	// REMEMBER for regexp, FindAllStringSubmatch will give you a slice of
	// slices.
	// For each slice, the first entry will be the full matched text, and the
	// rest of the entries will be capture groups from left to right.
	for _, rule := range input {
		// to deal with parsing the rule, we're going to take two passes at it.
		// One pass will separate the outer bag from the inner bag(s).
		results := rulesMatcher.FindAllStringSubmatch(rule, -1)
		if len(results) == 0 {
			continue
		}
		outerBag := results[0][1]
		// One pass will extract each inner bag.
		innerBagResults := innerBagMatcher.FindAllStringSubmatch(results[0][2], -1)
		for _, innerBagResult := range innerBagResults {
			if validOuterBags[innerBagResult[2]] == nil {
				validOuterBags[innerBagResult[2]] = make([]string, 0)
			}
			validOuterBags[innerBagResult[2]] = append(validOuterBags[innerBagResult[2]], outerBag)
		}
	}

	// next, create a map that we can use as a set. We'll use this as a way to
	// keep track of novel colors we've come across.
	possibleOuterBags := make(map[string]struct{})

	// we're going to do a non-recursive DFS, so we're going to create a stack
	// instead of using the call stack like we'd do in a recursive algo.
	var stack lib.StringStack
	// we're going to start our trip at the "shiny gold" node.
	stack.Push("shiny gold")

	// we'll use a stack to keep track of the nodes we need to visit next.
	// for every node we visit...
	for color, ok := stack.Pop(); ok; color, ok = stack.Pop() {
		// add the node to the set of possible outer bags
		var dummy struct{}
		possibleOuterBags[color] = dummy

		// check to see if there are any bags that can contain our bag
		nextColors, ok := validOuterBags[color]
		if !ok {
			// if no bag color can contain our bag, we've reached a leaf of the
			// tree.
			// We can get another color from the stack.
			continue
		}

		// for each bag that can contain our bag, push it onto the stack.
		// this will let us continue our DFS traversal.
		for _, nextColor := range nextColors {
			stack.Push(nextColor)
		}
	}

	// funny little thing. because we started our search with "shiny gold", it
	// ended up in our set. assuming the graph has no cycles, a shiny gold bag
	// can't contain a shiny gold bag. so we can remove it.
	delete(possibleOuterBags, "shiny gold")

	// the size of our "set" is now the answer to part 1!
	return len(possibleOuterBags)
}

type innerBag struct {
	Count int
	Color string
}

func part2(input []string) int {
	// with part 2, we're going to use recursion to solve the problem.
	// we're going to define a function and recursively call it.
	// unlike part 1, we're not going to reverse the rules.
	innerBags := make(map[string][]innerBag)

	for _, rule := range input {
		matchingRule := rulesMatcher.FindAllStringSubmatch(rule, -1)
		if len(matchingRule) == 0 {
			continue
		}

		outerBag := matchingRule[0][1]

		innerBagMatches := innerBagMatcher.FindAllStringSubmatch(matchingRule[0][2], -1)
		for _, innerBagMatch := range innerBagMatches {
			if innerBags[outerBag] == nil {
				innerBags[outerBag] = make([]innerBag, 0)
			}

			count, _ := strconv.Atoi(innerBagMatch[1])
			innerBags[outerBag] = append(innerBags[outerBag], innerBag{
				Count: count,
				Color: innerBagMatch[2],
			})
		}
	}

	// By definition, CountBagsWithin includes the containing bag in
	return CountBagsWithin("shiny gold", innerBags) - 1
}

// CountBagsWithin will recursively tabulate all the bags within a compliant
// bag, inclusive of the containing bag.
func CountBagsWithin(color string, innerBags map[string][]innerBag) int {
	// we always want to count the containing bag.
	result := 1

	// if we can't find any rules for this bag, there can be no other bags in
	// this bag. bail out.
	nextColors, ok := innerBags[color]
	if !ok {
		return result
	}

	// recursively call this function for each of the inner bags, multipliying
	// the result by its count.
	for _, nextColor := range nextColors {
		result += nextColor.Count * CountBagsWithin(nextColor.Color, innerBags)
	}

	return result
}
