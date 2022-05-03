package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []instruction) int {
	// we'll execute the code from the example! Our "memory" will be a slice of
	// instructions, and we'll modify the counter and accumulator based on the instruction.
	accumulator := 0
	visited := make([]bool, len(input))
	for i := 0; i <= len(input) && !visited[i]; {
		// We'll mark this instruction as visited, so that our for loop exits if
		// we hit this instruction again.
		visited[i] = true

		// now we'll parse the command
		switch input[i].Command {
		case "nop":
			// nop just moves the program counter
			i++
		case "acc":
			// acc will add to the accumulator and move the program counter
			accumulator += input[i].Argument
			i++
		case "jmp":
			// jmp will modify the program counter
			i += input[i].Argument
		}
	}
	return accumulator
}

type executionResult struct {
	Exited      bool
	Accumulator int
}

type executionSnapshot struct {
	ProgramCounter int
	Accumulator    int
}

func part2(input []instruction) int {
	// we're going to brute-force this problem.

	// first, we'll execute the code once, capturing all the instructions that
	// got hit. if we need only change one instruction to get the program to
	// pass, it will be one of the instructions executed in the loop.
	candidateInstructions := make([]executionSnapshot, 0)
	visited := make([]bool, len(input))
	accumulator := 0
	for i := 0; i <= len(input) && !visited[i]; {
		// We'll mark this instruction as visited, so that our for loop exits if
		// we hit this instruction again.
		visited[i] = true

		// now we'll parse the command
		switch input[i].Command {
		case "nop":
			// nop just moves the program counter
			candidateInstructions = append(candidateInstructions, executionSnapshot{ProgramCounter: i, Accumulator: accumulator})
			i++
		case "acc":
			// acc will add to the accumulator and move the program counter
			accumulator += input[i].Argument
			i++
		case "jmp":
			// jmp will modify the program counter
			candidateInstructions = append(candidateInstructions, executionSnapshot{ProgramCounter: i, Accumulator: accumulator})
			i += input[i].Argument
		}
	}

	// next, we'll re-run the code, changing a single jmp/nop instruction each
	// time.
	// this could take awhile; let's speed up the process using goroutines!

	// first, we'll make a channel that the goroutines can use to send their
	// execution results back to the main thread.
	result := make(chan executionResult, len(candidateInstructions))

	// next, we'll make a channel that the main goroutine can use to shut down all other children.
	done := make(chan interface{})

	for _, snapshot := range candidateInstructions {
		go executeProgram(input, snapshot, result, done)
	}

	reports := 0
	for r := range result {
		reports++
		if r.Exited {
			close(done)
			fmt.Printf("collected %d reports from %d candidates\n", reports, len(candidateInstructions))
			return (r.Accumulator)
		}
		if reports == len(candidateInstructions) {
			panic(errors.New("no solution found"))
		}
	}
	panic(errors.New("unreachable"))
}

const (
	CmdJmp = "jmp"
	CmdNop = "nop"
	CmdAcc = "acc"
)

func executeProgram(instructions []instruction, snapshot executionSnapshot, result chan executionResult, done chan interface{}) {
	accumulator := snapshot.Accumulator
	visited := make([]bool, len(instructions))
	pc := snapshot.ProgramCounter

	for pc < len(instructions) {
		select {
		case <-done: // main goroutine has its solution, stop running.
			return
		default:
		}

		// signal a loop detection and stop execution
		if visited[pc] {
			result <- executionResult{
				Exited:      false,
				Accumulator: accumulator,
			}
			return
		}

		command, arg := instructions[pc].Command, instructions[pc].Argument
		if pc == snapshot.ProgramCounter {
			switch command {
			case CmdNop:
				command = CmdJmp
			case CmdJmp:
				command = CmdNop
			}
		}
		visited[pc] = true
		switch command {
		case CmdJmp:
			pc += arg
		case CmdAcc:
			accumulator += arg
			pc++
		case CmdNop:
			pc++
		}
	}

	// if we got here, then we managed to exit correctly!
	result <- executionResult{
		Exited:      true,
		Accumulator: accumulator,
	}
}

type instruction struct {
	Command  string
	Argument int
}

func getInput() ([]instruction, error) {
	result := make([]instruction, 0)
	f, err := os.Open("inputs/day8.txt")
	if err != nil {
		return result, err
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		parts := strings.Fields(s.Text())
		argument, _ := strconv.Atoi(parts[1])
		result = append(result, instruction{
			Command:  parts[0],
			Argument: argument,
		})
	}
	return result, nil
}
