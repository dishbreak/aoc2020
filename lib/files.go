package lib

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

//GetInput will return a slice of strings, one entry per line, in-order.
func GetInput(filename string) ([]string, error) {
	f, err := os.Open(filename)
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

// GetInputAsSections returns back the input as a slice of string slices, where
// each top-level slice represents a section of the input. Useful for puzzle
// inputs with multiple sections e.g 2020 day 19.
func GetInputAsSections(filename string) ([][]string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	s := string(b)

	result := make([][]string, 0)

	for _, section := range strings.Split(s, "\n\n") {
		result = append(result, strings.Split(section, "\n"))
	}

	return result, nil
}
