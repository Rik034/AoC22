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
	hand := make(map[string]string)

	hand["A X"] = "tie"
	hand["A Y"] = "win"
	hand["A Z"] = "lose"

	hand["B X"] = "lose"
	hand["B Y"] = "tie"
	hand["B Z"] = "win"

	hand["C X"] = "win"
	hand["C Y"] = "lose"
	hand["C Z"] = "tie"

	points := make(map[string]int)

	points["X"] = 1
	points["Y"] = 2
	points["Z"] = 3

	points["win"] = 6
	points["tie"] = 3
	points["lose"] = 0

	totalPoints := 0

	for _, line := range input {
		you := string(line[2])
		winTieLose := hand[line]
		totalPoints += points[winTieLose] + points[you]
	}
	return totalPoints
}

func part2(input []string) int {
	handToPlay := make(map[string]string)

	handToPlay["A X"] = "C"
	handToPlay["A Y"] = "A"
	handToPlay["A Z"] = "B"

	handToPlay["B X"] = "A"
	handToPlay["B Y"] = "B"
	handToPlay["B Z"] = "C"

	handToPlay["C X"] = "B"
	handToPlay["C Y"] = "C"
	handToPlay["C Z"] = "A"

	points := make(map[string]int)

	points["X"] = 0
	points["Y"] = 3
	points["Z"] = 6

	points["A"] = 1
	points["B"] = 2
	points["C"] = 3

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
