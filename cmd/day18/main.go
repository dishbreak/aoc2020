package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day18.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func add(result, operand int) int {
	return result + operand
}

func mult(result, operand int) int {
	return result * operand
}

type TokenType int

const (
	OPERATOR TokenType = iota
	VALUE
	LPAREN
	RPAREN
)

type Token interface {
	GetType() TokenType
}

type Operator struct {
	Precedence int
	Evaluate   func(int, int) int
}

func (o Operator) GetType() TokenType {
	return OPERATOR
}

type Value struct {
	V int
}

func (v Value) GetType() TokenType {
	return VALUE
}

type Paren struct {
	parenType TokenType
}

func (p Paren) GetType() TokenType {
	return p.parenType
}

func tokenize(statement string, operators map[string]Operator) []Token {
	result := make([]Token, 0)
	for _, part := range strings.Fields(statement) {

		if operator, ok := operators[part]; ok {
			result = append(result, operator)
		} else {
			if strings.HasPrefix(part, "(") {
				i := 0
				for ; part[i] == '('; i++ {
					result = append(result, Paren{parenType: LPAREN})
				}
				part = string(part[i:])
			}

			rparens := 0
			if strings.HasSuffix(part, ")") {
				j := len(part) - 1
				for ; part[j] == ')'; j-- {
					rparens++
				}
				part = string(part[:j+1])
			}

			parsed, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}

			result = append(result, Value{V: parsed})
			for i := 0; i < rparens; i++ {
				result = append(result, Paren{parenType: RPAREN})
			}
		}
	}
	return result
}

func infixToRpn(tokens []Token) []Token {
	result := make([]Token, 0)
	opsStack := list.New()
	for _, token := range tokens {
		switch token.GetType() {
		case VALUE:
			result = append(result, token)
		case OPERATOR:
			currentOperator := token.(Operator)
			halt := false

			for opsStack.Len() != 0 && !halt {
				lastToken := opsStack.Back().Value.(Token)
				switch lastToken.GetType() {
				case OPERATOR:
					lastOperator := lastToken.(Operator)
					if lastOperator.Precedence >= currentOperator.Precedence {
						result = append(result, lastOperator)
						opsStack.Remove(opsStack.Back())
					} else {
						halt = true
					}
				default:
					halt = true
				}
			}
			fallthrough
		case LPAREN:
			opsStack.PushBack(token)
		case RPAREN:
			halt := false
			for opsStack.Len() != 0 && !halt {
				lastToken := opsStack.Back().Value.(Token)
				switch lastToken.GetType() {
				case OPERATOR:
					result = append(result, lastToken)
					opsStack.Remove(opsStack.Back())
				case LPAREN:
					opsStack.Remove(opsStack.Back())
					halt = true
				}
			}
		}
	}

	for opsStack.Len() != 0 {
		lastToken := opsStack.Back().Value.(Token)
		result = append(result, lastToken)
		opsStack.Remove(opsStack.Back())

	}

	return result
}

func evaluateStatement(statement string, operators map[string]Operator) int {
	if statement == "" {
		return 0
	}

	tokens := tokenize(statement, operators)
	rpnStream := infixToRpn(tokens)

	outputStack := list.New()
	for _, token := range rpnStream {
		switch token.GetType() {
		case VALUE:
			value := token.(Value)
			outputStack.PushBack(value.V)
		case OPERATOR:
			op := token.(Operator)
			rhs := outputStack.Back().Value.(int)
			outputStack.Remove(outputStack.Back())
			lhs := outputStack.Back().Value.(int)
			outputStack.Remove(outputStack.Back())
			outputStack.PushBack(op.Evaluate(lhs, rhs))
		}
	}

	return outputStack.Back().Value.(int)
}

func part1(input []string) int {
	result := 0

	operators := map[string]Operator{
		"*": Operator{Precedence: 1, Evaluate: mult},
		"+": Operator{Precedence: 1, Evaluate: add},
	}

	for _, statement := range input {
		result += evaluateStatement(statement, operators)
	}

	return result
}

func part2(input []string) int {
	result := 0

	operators := map[string]Operator{
		"*": Operator{Precedence: 1, Evaluate: mult},
		"+": Operator{Precedence: 2, Evaluate: add},
	}

	for _, statement := range input {
		result += evaluateStatement(statement, operators)
	}

	return result
}
