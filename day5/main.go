package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Peek() string {
	index := len(*s) - 1
	element := (*s)[index]
	return element
}

func (s *Stack) Reverse() {
	for i := 0; i < len(*s)/2; i++ {
		j := len(*s) - i - 1
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

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

func splitInput(input []string) ([]string, []string, int) {
	var (
		drawing []string
		moves   []string
	)
	blankLineIndex := 0
	numOfStacks := 0
	for i, line := range input {
		if string(line[0]) == " " && string(line[1]) != " " {
			numOfStacks, _ = strconv.Atoi(string(line[len(line)-2]))
			blankLineIndex = i + 1
			break
		}
		drawing = append(drawing, line)
	}
	for i, line := range input {
		if i > blankLineIndex {
			moves = append(moves, line)
		}
	}
	return drawing, moves, numOfStacks
}

func createStacks(input []string) []Stack {
	drawing, _, numOfStacks := splitInput(input)
	stacks := make([]Stack, numOfStacks)
	i := 0
	for _, line := range drawing {
		for j := 1; j < len(line); j += 4 {
			if string(line[j]) != " " {
				stacks[i].Push(string(line[j]))
			}
			i++
		}
		i = 0
	}
	for i := range stacks {
		stacks[i].Reverse()
	}
	return stacks
}

func clearMoves(input []string) [][]int {
	_, moves, _ := splitInput(input)
	intMoves := make([][]int, len(moves))
	for i, line := range moves {
		str := strings.ReplaceAll(line, "move ", "")
		str = strings.ReplaceAll(str, "from ", "")
		str = strings.ReplaceAll(str, "to ", "")
		move := strings.Split(str, " ")
		intMoves[i] = make([]int, 3)
		intMoves[i][0], _ = strconv.Atoi(move[0])
		intMoves[i][1], _ = strconv.Atoi(move[1])
		intMoves[i][2], _ = strconv.Atoi(move[2])
	}
	return intMoves
}

func part1(input []string) string {
	stacks := createStacks(input)
	moves := clearMoves(input)
	toReturn := ""

	for _, move := range moves {
		for i := 0; i < move[0]; i++ {
			toMove, _ := stacks[move[1]-1].Pop()
			stacks[move[2]-1].Push(toMove)
		}
	}

	for _, stack := range stacks {
		toReturn += stack.Peek()
	}

	return toReturn
}

func main() {
	// input := ReadByLine("sample.txt")
	input := ReadByLine("input.txt")
	topCrates := part1(input)
	fmt.Println("Part 1 =", topCrates)
	fmt.Println("Part 2 =")
}
