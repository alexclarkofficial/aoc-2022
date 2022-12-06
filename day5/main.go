package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getStacks(filename string) ([][]string, [][3]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	stacks := []string{}
	instructions := [][3]int{}
	gatheringStacks := true

	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			gatheringStacks = false
		} else if gatheringStacks {
			stacks = append(stacks, line)
		} else {
			move := [3]int{}
			fmt.Sscanf(line, "move %d from %d to %d", &move[0], &move[1], &move[2])
			instructions = append(instructions, move)
		}
	}

	counts := stacks[len(stacks)-1]
	stacksWidth := (len(counts) + 1) / 4
	columns := make([][]string, stacksWidth)

	for i := len(stacks) - 2; i >= 0; i-- {
		for j, col := 1, 0; j <= len(counts); j, col = j+4, col+1 {
			crate := string(stacks[i][j])
			if crate != " " {
				columns[col] = append(columns[col], crate)
			}
		}
	}

	return columns, instructions, nil
}

func moveSingleCrates(columns [][]string, moves [][3]int) string {
	for _, move := range moves {
		crateCount := move[0]
		fromColumn := &columns[move[1]-1]
		toColumn := &columns[move[2]-1]

		for i := 0; i < crateCount; i++ {
			topCrate := (*fromColumn)[len(*fromColumn)-1]
			*toColumn = append(*toColumn, topCrate)
			*fromColumn = (*fromColumn)[:len(*fromColumn)-1]
		}
	}

	answer := ""
	for _, col := range columns {
		answer += col[len(col)-1]
	}
	return answer
}

func moveMultipleCrates(columns [][]string, moves [][3]int) string {
	for _, move := range moves {
		crateCount := move[0]
		fromColumn := &columns[move[1]-1]
		toColumn := &columns[move[2]-1]

		topCrates := []string{}
		for i := 0; i < crateCount; i++ {
			topCrates = append([]string{(*fromColumn)[len(*fromColumn)-1]}, topCrates...)
			*fromColumn = (*fromColumn)[:len(*fromColumn)-1]
		}
		*toColumn = append(*toColumn, topCrates...)
	}

	answer := ""
	for _, col := range columns {
		answer += col[len(col)-1]
	}
	return answer
}

func main() {
	filename := os.Args[1]

	stacks, instructions, err := getStacks(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	partOneAnswer := moveSingleCrates(stacks, instructions)
	fmt.Printf("Part one answer: %s\n", partOneAnswer)
	stacks, instructions, err = getStacks(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	partTwoAnswer := moveMultipleCrates(stacks, instructions)
	fmt.Printf("Part two answer: %s\n", partTwoAnswer)
}
