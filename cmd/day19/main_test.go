package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = [][]string{
	[]string{
		`0: 4 1 5`,
		`1: 2 3 | 3 2`,
		`2: 4 4 | 5 5`,
		`3: 4 5 | 5 4`,
		`4: "a"`,
		`5: "b"`,
	},
	[]string{
		"ababbb",
		"bababa",
		"abbbab",
		"aaabbb",
		"aaaabbb",
	},
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 2, part1(input))
}

func TestParseRule(t *testing.T) {
	assert.Equal(t, rule{number: 0, conditionals: [][]int{
		[]int{4, 1, 5},
	}}, newRule(input[0][0]))
	assert.Equal(t, rule{number: 1, conditionals: [][]int{
		[]int{2, 3},
		[]int{3, 2},
	}}, newRule(input[0][1]))

	assert.Equal(t, rule{number: 4, baseCondition: 'a'}, newRule(input[0][4]))
}

func TestPart2(t *testing.T) {
	part2Input := [][]string{
		[]string{
			"42: 9 14 | 10 1",
			"9: 14 27 | 1 26",
			"10: 23 14 | 28 1",
			`1: "a"`,
			"11: 42 31",
			"5: 1 14 | 15 1",
			"19: 14 1 | 14 14",
			"12: 24 14 | 19 1",
			"16: 15 1 | 14 14",
			"31: 14 17 | 1 13",
			"6: 14 14 | 1 14",
			"2: 1 24 | 14 4",
			"0: 8 11",
			"13: 14 3 | 1 12",
			"15: 1 | 14",
			"17: 14 2 | 1 7",
			"23: 25 1 | 22 14",
			"28: 16 1",
			"4: 1 1",
			"20: 14 14 | 1 15",
			"3: 5 14 | 16 1",
			"27: 1 6 | 14 18",
			`14: "b"`,
			"21: 14 1 | 1 14",
			"25: 1 1 | 1 14",
			"22: 14 14",
			"8: 42",
			"26: 14 22 | 1 20",
			"18: 15 15",
			"7: 14 5 | 1 21",
			"24: 14 1",
		},
		[]string{
			"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
			"bbabbbbaabaabba",
			"babbbbaabbbbbabbbbbbaabaaabaaa",
			"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
			"bbbbbbbaaaabbbbaaabbabaaa",
			"bbbababbbbaaaaaaaabbababaaababaabab",
			"ababaaaaaabaaab",
			"ababaaaaabbbaba",
			"baabbaaaabbaaaababbaababb",
			"abbbbabbbbaaaababbbbbbaaaababb",
			"aaaaabbaabaaaaababaa",
			"aaaabbaaaabbaaa",
			"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
			"babaaabbbaaabaababbaabababaaab",
			"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
		},
	}
	assert.Equal(t, 12, part2(part2Input))
}
