package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type TravelingParty struct {
	responses [26]int
	members   int
}

func (t *TravelingParty) RegisterResponse(response string) {
	t.members++
	for _, statement := range response {
		t.responses[statement-'a']++
	}
}

func (t *TravelingParty) UniqueResponses() int {
	result := 0
	for _, statement := range t.responses {
		if statement > 0 {
			result++
		}
	}
	return result
}

func (t *TravelingParty) UnanimousResponses() int {
	result := 0
	for _, statement := range t.responses {
		if statement == t.members {
			result++
		}
	}
	return result
}

func part1(input []string) int {
	party := &TravelingParty{}
	result := 0
	for _, response := range input {
		switch response {
		case "":
			result += party.UniqueResponses()
			party = &TravelingParty{}
		default:
			party.RegisterResponse(response)
		}
	}
	return result
}

func part2(input []string) int {
	party := &TravelingParty{}
	result := 0
	for _, response := range input {
		switch response {
		case "":
			result += party.UnanimousResponses()
			party = &TravelingParty{}
		default:
			party.RegisterResponse(response)
		}
	}
	return result
}

func getInput() ([]string, error) {
	f, err := os.Open("inputs/day6.txt")
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		result = append(result, s.Text())
	}

	result = append(result, "")
	return result, nil
}
