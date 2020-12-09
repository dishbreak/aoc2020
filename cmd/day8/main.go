package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

	// first, we'll make a channel for transmitting the instruction to flip out
	// to the goroutines.
	flippedInstruction := make(chan executionSnapshot, len(candidateInstructions))
	// next, we'll make a channel that the goroutines can use to send their
	// exeuction results back to the main thread.
	result := make(chan executionResult, len(candidateInstructions))

	var wg sync.WaitGroup
	// start 10 workers to run exectuions of the code.
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(flippedInstruction <-chan executionSnapshot, result chan<- executionResult) {
			defer wg.Done()
			for j := range flippedInstruction {
				// re-execute the program.
				// note that this time the program counter is declared outside
				// the for loop because we need it to check if the program exited.
				accumulator := j.Accumulator
				visited := make([]bool, len(input))
				k := j.ProgramCounter

				for k < len(input) && !visited[k] {
					visited[k] = true
					command := input[k].Command

					// flip the command if the program counter matches the input
					// we got.
					if k == j.ProgramCounter {
						switch command {
						case "jmp":
							command = "nop"
						case "nop":
							command = "jmp"
						}
					}

					// execute the command as normal.
					switch command {
					case "nop":
						k++
					case "acc":
						accumulator += input[k].Argument
						k++
					case "jmp":
						k += input[k].Argument
					}
				}
				// record the result
				runResult := executionResult{
					Exited:      k == len(input),
					Accumulator: accumulator,
				}
				result <- runResult
			}
		}(flippedInstruction, result)
	}

	// feed the inputs to the workers.
	for _, inst := range candidateInstructions {
		flippedInstruction <- inst
	}
	// important! we need to close the channel so workers know to exit.
	close(flippedInstruction)

	accumulator = 0
	// collect results
	for i := 0; i < len(candidateInstructions); i++ {
		s := <-result
		if s.Exited {
			accumulator = s.Accumulator
		}
	}

	// make sure everything is done before returning.
	wg.Wait()
	return accumulator
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
