package main

import "sort"

func intersection(one []string, other []string) []string {
	result := make([]string, 0)
	if len(one) == 0 || len(other) == 0 {
		return result
	}

	sort.Strings(one)
	sort.Strings(other)

	i := 0
	j := 0

	for i < len(one) && j < len(other) {
		switch {
		case one[i] < other[j]:
			i++
		case one[i] == other[j]:
			result = append(result, one[i])
			i++
			j++
		default:
			j++
		}

	}

	return result
}
