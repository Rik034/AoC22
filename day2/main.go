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

func part1(input []string) int {
	hand := map[string]string{
		"A X": "tie",
		"A Y": "win",
		"A Z": "lose",

		"B X": "lose",
		"B Y": "tie",
		"B Z": "win",

		"C X": "win",
		"C Y": "lose",
		"C Z": "tie",
	}

	points := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,

		"win":  6,
		"tie":  3,
		"lose": 0,
	}

	totalPoints := 0

	for _, line := range input {
		you := string(line[2])
		winTieLose := hand[line]
		totalPoints += points[winTieLose] + points[you]
	}
	return totalPoints
}

func part2(input []string) int {
	handToPlay := map[string]string{
		"A X": "C",
		"A Y": "A",
		"A Z": "B",

		"B X": "A",
		"B Y": "B",
		"B Z": "C",

		"C X": "B",
		"C Y": "C",
		"C Z": "A",
	}

	points := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,

		"A": 1,
		"B": 2,
		"C": 3,
	}

	totalPoints := 0

	for _, line := range input {
		you := handToPlay[line]
		winTieLose := string(line[2])
		totalPoints += points[you] + points[winTieLose]
	}
	return totalPoints
}

func main() {
	// input := ReadByLine("sample.txt")
	input := ReadByLine("input.txt")

	points := part1(input)
	fmt.Println("part1 = ", points)

	points = part2(input)
	fmt.Println("part2 = ", points)
}
