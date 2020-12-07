package lib

import (
	"bufio"
	"os"
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
