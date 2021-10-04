package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day21.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	ingredients := make(map[string]int)
	allergens := make(map[string][]string)

	for _, line := range input {
		if line == "" {
			continue
		}
		ilist := newIngredientListFromString(line)

		for _, ingredient := range ilist.ingredients {
			ingredients[ingredient]++
		}

		for _, allergen := range ilist.allegens {
			knownIngredients, ok := allergens[allergen]
			if !ok {
				knownIngredients = make([]string, len(ingredients))
				copy(knownIngredients, ilist.ingredients)
			} else {
				knownIngredients = intersection(knownIngredients, ilist.ingredients)
			}
			allergens[allergen] = knownIngredients
		}
	}

	for _, knownIngredients := range allergens {
		for _, knowinIngredient := range knownIngredients {
			delete(ingredients, knowinIngredient)
		}
	}

	acc := 0
	for _, count := range ingredients {
		acc += count
	}
	return acc
}

func part2(input []string) int {
	return 0
}
