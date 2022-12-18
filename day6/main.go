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

func uniqueChars(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}

func part1(input []string) int {
	stream := input[0]

	for i := 0; i < len(stream); i++ {
		if i-4 >= 0 {
			str := stream[i-4 : i]
			if uniqueChars(str) {
				return i
			}
		}
	}
	return 0
}

func part2(input []string) int {
	stream := input[0]

	for i := 0; i < len(stream); i++ {
		if i-14 >= 0 {
			str := stream[i-14 : i]
			if uniqueChars(str) {
				return i
			}
		}
	}
	return 0
}

func main() {
	// input := ReadByLine("sample.txt")
	input := ReadByLine("input.txt")
	fmt.Println("Part 1 =", part1(input))
	fmt.Println("Part 2 =", part2(input))
}
