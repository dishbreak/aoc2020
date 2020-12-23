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

func noop(result, operand int) int {
	return result
}

type stackFrame struct {
	operation   func(int, int) int
	accumulator int
}

func evaluateStatement(statement string) int {
	result := 0
	operation := add

	stack := list.New()

	for _, field := range strings.Fields(statement) {
		switch {
		case field == "+":
			operation = add
		case field == "*":
			operation = mult
		case strings.HasPrefix(field, "("):
			i := 0
			for ; field[i] == '('; i++ {
				stack.PushBack(stackFrame{
					operation:   operation,
					accumulator: result,
				})
				result = 0
				operation = add
			}
			field = string(field[i:])
			operand, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			result = operand
		case strings.HasSuffix(field, ")"):
			i := len(field) - 1
			for ; field[i] == ')'; i-- {
			}
			parens := len(field) - i - 1
			field = string(field[:i+1])

			operand, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}

			result = operation(result, operand)

			for j := 0; j < parens; j++ {
				if stack.Len() == 0 {
					panic(fmt.Errorf("unbalanced statement: %s", statement))
				}

				frame, ok := stack.Back().Value.(stackFrame)
				if !ok {
					panic(fmt.Errorf("unexpected stack value %v", stack.Back()))
				}

				result = frame.operation(frame.accumulator, result)
				stack.Remove(stack.Back())
			}
		default:
			operand, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			result = operation(result, operand)
		}
	}

	return result
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

func tokenize(statement string) []Token {
	result := make([]Token, 0)
	for _, part := range strings.Fields(statement) {
		switch {
		case part == "*":
			result = append(result, Operator{Precedence: 1, Evaluate: mult})
		case part == "+":
			result = append(result, Operator{Precedence: 2, Evaluate: add})
		default:
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

func evaluateStatementV2(statement string) int {
	if statement == "" {
		return 0
	}

	tokens := tokenize(statement)
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

	for _, statement := range input {
		result += evaluateStatement(statement)
	}

	return result
}

func part2(input []string) int {
	result := 0

	for _, statement := range input {
		result += evaluateStatementV2(statement)
	}

	return result
}
