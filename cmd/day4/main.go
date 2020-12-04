package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
	missingFields  int
}

func (p *Passport) RegisterData(input string) error {
	parts := strings.Fields(input)

	for _, part := range parts {
		parsedField := strings.Split(part, ":")
		field, value := parsedField[0], parsedField[1]
		var err error
		p.missingFields--
		switch field {
		case "byr":
			p.BirthYear, err = strconv.Atoi(value)
		case "iyr":
			p.IssueYear, err = strconv.Atoi(value)
		case "eyr":
			p.ExpirationYear, err = strconv.Atoi(value)
		case "hgt":
			p.Height = value
		case "hcl":
			p.HairColor = value
		case "ecl":
			p.EyeColor = value
		case "pid":
			p.PassportID = value
		default:
			p.missingFields++
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Passport) NoMissingFields() bool {
	return p.missingFields == 0
}

func (p *Passport) IsValid() bool {
	if !p.NoMissingFields() {
		return false
	}

	if p.BirthYear < 1920 || p.BirthYear > 2002 {
		return false
	}

	if p.IssueYear < 2010 || p.IssueYear > 2020 {
		return false
	}

	if p.ExpirationYear < 2020 || p.ExpirationYear > 2030 {
		return false
	}

	switch {
	case strings.HasSuffix(p.Height, "in"):
		if inches, _ := strconv.Atoi(strings.TrimSuffix(p.Height, "in")); inches < 59 || inches > 76 {
			return false
		}
	case strings.HasSuffix(p.Height, "cm"):
		if cms, _ := strconv.Atoi(strings.TrimSuffix(p.Height, "cm")); cms < 150 || cms > 193 {
			return false
		}
	default:
		return false
	}

	validColor := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	if !validColor.MatchString(p.EyeColor) {
		return false
	}

	validHex := regexp.MustCompile("^#[a-f|0-9]{6}$")
	if !validHex.MatchString(p.HairColor) {
		return false
	}

	validPid := regexp.MustCompile("^[0-9]{9}$")
	if !validPid.MatchString(p.PassportID) {
		return false
	}

	return true
}

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []*Passport) int {
	count := 0
	for _, passport := range input {
		if passport.NoMissingFields() {
			count++
		}
	}
	return count
}

func part2(input []*Passport) int {
	count := 0
	for _, passport := range input {
		if passport.IsValid() {
			count++
		}
	}
	return count
}

func getInput() ([]*Passport, error) {
	f, err := os.Open("inputs/day4.txt")
	if err != nil {
		return nil, err
	}

	result := make([]*Passport, 0)

	passport := &Passport{missingFields: 7}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "":
			result = append(result, passport)
			passport = &Passport{missingFields: 7}
		default:
			if err := passport.RegisterData(line); err != nil {
				return result, err
			}
		}
	}
	result = append(result, passport)
	return result, nil
}
