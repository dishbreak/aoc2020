package main

import "strings"

type ingredientList struct {
	ingredients []string
	allegens    []string
}

func newIngredientListFromString(input string) ingredientList {
	input = strings.TrimSuffix(input, ")")
	parts := strings.Split(input, " (contains ")

	return ingredientList{
		ingredients: strings.Split(parts[0], " "),
		allegens:    strings.Split(parts[1], ", "),
	}
}
