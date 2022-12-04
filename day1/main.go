package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func part1(input []string) []int {
	elves := []int{}
	var sum int = 0
	for _, line := range input {
		if line == "" {
			elves = append(elves, sum)
			sum = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err.Error())
			}
			sum += num
		}
	}
	elves = append(elves, sum)
	return elves
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	// input := ReadByLine("sample.txt")
	input := ReadByLine("input.txt")
	elves := part1(input)

	sort.Ints(elves)
	reverse(elves)

	fmt.Println("Part 1 = ", elves[0])
	fmt.Println("Part 2 = ", elves[0]+elves[1]+elves[2])
}
