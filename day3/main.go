package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadByLine(filename string) []string {
	lines := []string{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return lines
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func part1(input []string) {
	for _, line := range input {
		// TODO
		fmt.Println(line)
	}
}

func main() {
	input := ReadByLine("sample.txt")
	// input := ReadByLine("input.txt")
	part1(input)
}
