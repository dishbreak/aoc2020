package main

import (
	"fmt"
)

type BoardingPass struct {
	Row int
	Col int
}

func NewBoardingPass(input string) BoardingPass {
	rowSymbols := string(input[:7])
	colSymbols := string(input[7:len(input)])
	return BoardingPass{
		Row: binarySearch(rowSymbols, 'F', 'B', 0, 127),
		Col: binarySearch(colSymbols, 'L', 'R', 0, 7),
	}
}

func (b BoardingPass) GetID() int {
	return (b.Row * 8) + b.Col
}

func binarySearch(input string, low, high rune, min, max int) int {
	for _, x := range input {
		switch x {
		case high:
			min = (min+max)/2 + 1
		case low:
			max = (min + max) / 2
		}
	}
	return min
}

func main() {
	pass := "BFFFFFB"
	min, max := 0, 127
	for _, x := range pass {
		switch x {
		case 'B':
			min = (min+max)/2 + 1
		case 'F':
			max = (min + max) / 2
		}
	}
	fmt.Println(min * 8)
	fmt.Println(binarySearch(pass, 'F', 'B', 0, 127) * 8)
	fmt.Println(NewBoardingPass("BFFFFFBRRR").GetID())
}
