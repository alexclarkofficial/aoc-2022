package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	fmt.Println(filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	forest := []instruction{}
	for sc.Scan() {
		dir := ""
		mag := 0

		fmt.Sscanf(sc.Text(), "%s %d", &dir, &mag)
		forest = append(forest, instruction{dir: dir, mag: mag})
	}

	partOneAnswer := followInstructions(forest, 1)
	fmt.Printf("Part one answer: %d\n", partOneAnswer)
	partTwoAnswer := followInstructions(forest, 9)
	fmt.Printf("Part two answer: %d\n", partTwoAnswer)
}

func followInstructions(instructions []instruction, ropeLength int) int {
	headPos := [2]int{0, 0}
	knotPositions := make([][2]int, ropeLength)
	for i, _ := range knotPositions {
		knotPositions[i] = [2]int{0, 0}
	}
	tailHistory := map[string]bool{}
	for _, step := range instructions {
		for i := 1; i <= step.mag; i++ {
			headPos = moveHead(step.dir, headPos)
			for j, _ := range knotPositions {
				if j == 0 {
					knotPositions[j] = moveTail(step.dir, knotPositions[j], headPos)
				} else {
					knotPositions[j] = moveTail(step.dir, knotPositions[j], knotPositions[j-1])
				}
			}
			tailHistory[fmt.Sprintf("%d.%d", knotPositions[ropeLength-1][0], knotPositions[ropeLength-1][1])] = true
		}
	}
	count := 0
	for range tailHistory {
		count++
	}

	return count
}

func moveHead(dir string, currentPos [2]int) [2]int {
	switch dir {
	case "U":
		return [2]int{currentPos[0], currentPos[1] + 1}
	case "D":
		return [2]int{currentPos[0], currentPos[1] - 1}
	case "R":
		return [2]int{currentPos[0] + 1, currentPos[1]}
	case "L":
		return [2]int{currentPos[0] - 1, currentPos[1]}
	}
	return [2]int{0, 0}
}

func moveTail(dir string, currentPos [2]int, leadPos [2]int) [2]int {
	xDist := 0
	yDist := 0
	if leadPos[0] > currentPos[0] {
		xDist = 1
	}
	if leadPos[0] < currentPos[0] {
		xDist = -1
	}
	if leadPos[1] > currentPos[1] {
		yDist = 1
	}
	if leadPos[1] < currentPos[1] {
		yDist = -1
	}
	if absVal(leadPos[0]-currentPos[0]) <= 1 && absVal(leadPos[1]-currentPos[1]) <= 1 {
		return currentPos
	}

	return [2]int{currentPos[0] + xDist, currentPos[1] + yDist}
}

func absVal(integer int) int {
	if integer < 0 {
		return -integer
	} else {
		return integer
	}
}

type instruction struct {
	dir string
	mag int
}
