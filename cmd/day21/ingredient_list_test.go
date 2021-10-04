package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIngredientList(t *testing.T) {
	type testCase struct {
		input    string
		expected ingredientList
	}

	testCases := []testCase{
		{"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
			ingredientList{
				[]string{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
				[]string{"dairy", "fish"},
			},
		},

		{"trh fvjkl sbzzf mxmxvkd (contains dairy)",
			ingredientList{
				[]string{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
				[]string{"dairy"},
			},
		},

		{"sqjhc fvjkl (contains soy)",
			ingredientList{
				[]string{"sqjhc", "fvjkl"},
				[]string{"soy"},
			},
		},

		{"sqjhc mxmxvkd sbzzf (contains fish)",
			ingredientList{
				[]string{"sqjhc", "mxmxvkd", "sbzzf"},
				[]string{"fish"},
			},
		},
	}

	for idx, tc := range testCases {
		name := fmt.Sprintf("test case %d", idx)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, newIngredientListFromString(tc.input))
		})
	}
}
