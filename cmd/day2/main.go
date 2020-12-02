package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := GetInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []*Password) int {
	count := 0
	for _, password := range input {
		if password.IsValid() {
			count++
		}
	}

	return count
}

func part2(input []*Password) int {
	count := 0
	for _, password := range input {
		if password.IsTobogganCorpValid() {
			count++
		}
	}

	return count
}

type Password struct {
	MinChars   int
	MaxChars   int
	TargetChar string
	Password   string
	RawString  string
}

func (p *Password) IsValid() bool {
	// to find the number of occurrences of the character, we're going to
	// generate a copy WITHOUT any occurrences of the character, then compare
	// string lengths.
	nonTargetChars := strings.Replace(p.Password, p.TargetChar, "", -1)
	count := len(p.Password) - len(nonTargetChars)

	valid := p.MinChars <= count && count <= p.MaxChars

	return valid
}

func (p *Password) IsTobogganCorpValid() bool {
	inFirstPostion := p.TargetChar[0] == p.Password[p.MinChars-1]
	inSecondPosition := p.TargetChar[0] == p.Password[p.MaxChars-1]

	return (inFirstPostion || inSecondPosition) && !(inFirstPostion && inSecondPosition)
}

func NewPassword(input string) (*Password, error) {
	parts := strings.Fields(input)
	minMax := parts[0]
	target := parts[1]
	passwd := parts[2]

	minMaxParts := strings.Split(minMax, "-")
	minVal, err := strconv.Atoi(minMaxParts[0])
	if err != nil {
		return nil, err
	}
	maxVal, err := strconv.Atoi(minMaxParts[1])
	if err != nil {
		return nil, err
	}

	return &Password{
		MinChars:   minVal,
		MaxChars:   maxVal,
		TargetChar: strings.Split(target, ":")[0],
		Password:   passwd,
		RawString:  input,
	}, nil
}

func GetInput() ([]*Password, error) {
	f, err := os.Open("inputs/day2.txt")
	if err != nil {
		return nil, err
	}

	result := make([]*Password, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password, err := NewPassword(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, password)
	}

	return result, nil
}
