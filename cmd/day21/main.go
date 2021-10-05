package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day21.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %s\n", part2(input))
}

func analyze(input []string) (ingredients map[string]int, allergenToIngredient map[string]string) {
	ingredients = make(map[string]int)
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
				knownIngredients = make([]string, len(ilist.ingredients))
				copy(knownIngredients, ilist.ingredients)
			} else {
				knownIngredients = intersection(knownIngredients, ilist.ingredients)
			}
			allergens[allergen] = knownIngredients
		}
	}

	unsafeIngredients := make(map[string]int)
	allergenToIngredient = make(map[string]string)

	for len(allergenToIngredient) < len(allergens) {
		for allergen, possibleIngredients := range allergens {
			if _, ok := allergenToIngredient[allergen]; ok {
				continue
			}
			if len(possibleIngredients) == 1 {
				unsafeIngredients[possibleIngredients[0]]++
				allergenToIngredient[allergen] = possibleIngredients[0]
				continue
			}
			replacement := make([]string, 0)
			for _, possibleIngredient := range possibleIngredients {
				if _, ok := unsafeIngredients[possibleIngredient]; !ok {
					replacement = append(replacement, possibleIngredient)
				}
			}
			allergens[allergen] = replacement
		}
	}
	return
}

func part1(input []string) int {
	ingredients, allergens := analyze(input)
	for _, knownIngredient := range allergens {
		delete(ingredients, knownIngredient)
	}

	acc := 0
	for _, count := range ingredients {
		acc += count
	}
	return acc
}

func part2(input []string) string {
	_, allergens := analyze(input)

	keys := make([]string, len(allergens))
	values := make([]string, len(allergens))

	i := 0
	for key, _ := range allergens {
		keys[i] = key
		i++
	}

	sort.Strings(keys)

	for idx, key := range keys {
		values[idx] = allergens[key]
	}

	return strings.Join(values, ",")
}
