package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getElves(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	currentElfCalories := make([]int, 0)
	allElves := make([][]int, 0)

	for sc.Scan() {
		if len(sc.Text()) == 0 {
			allElves = append(allElves, currentElfCalories)
			currentElfCalories = make([]int, 0)
			continue
		}
		calories, err := strconv.Atoi(sc.Text())
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		currentElfCalories = append(currentElfCalories, calories)
	}

	return allElves, nil
}

func topElfTotal(elves [][]int) int {
	max := 0

	for _, elf := range elves {
		elfCalories := 0
		for _, calories := range elf {
			elfCalories += calories
		}

		if elfCalories > max {
			max = elfCalories
		}
	}

	return max
}

func topThreeElvesTotal(elves [][]int) int {
	topElves := [3]int{0, 0, 0}

	for _, elf := range elves {
		elfCalories := 0
		for _, calories := range elf {
			elfCalories += calories
		}

		if elfCalories > topElves[0] {
			topElves[2] = topElves[1]
			topElves[1] = topElves[0]
			topElves[0] = elfCalories
			continue
		}
		if elfCalories > topElves[1] {
			topElves[2] = topElves[1]
			topElves[1] = elfCalories
			continue
		}
		if elfCalories > topElves[2] {
			topElves[2] = elfCalories
			continue
		}
	}

	total := 0
	for _, calories := range topElves {
		total += calories
	}

	return total
}

func main() {
	filename := os.Args[1]

	allElves, err := getElves(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	partOneScore := topElfTotal(allElves)
	fmt.Printf("Part one score: %d\n", partOneScore)
	partTwoScore := topThreeElvesTotal(allElves)
	fmt.Printf("Part two score: %d\n", partTwoScore)
}
