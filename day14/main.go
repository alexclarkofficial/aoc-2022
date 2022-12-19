package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	lines := [][][2]int{}
	ceiling := 0
	floor := 0
	start := 99999999
	finish := 0
	for sc.Scan() {
		rawLine := strings.Split(sc.Text(), " -> ")
		line := [][2]int{}
		for _, p := range rawLine {
			coords := strings.Split(p, ",")
			x, err := strconv.Atoi(coords[0])
			y, err := strconv.Atoi(coords[1])
			if err != nil {
				log.Fatal(err)
				return
			}
			if x < start {
				start = x
			}
			if x > finish {
				finish = x
			}
			if y < ceiling {
				ceiling = y
			}
			if y > floor {
				floor = y
			}
			line = append(line, [2]int{x, y})
		}
		lines = append(lines, line)
	}

	width := finish - start + 1
	height := floor - ceiling + 1
	padding := 1000
	offset := start - (padding / 2)

	board := drawBoard(lines, height, width, padding, start)
	partOneAnswer := partOne(board, offset, floor)
	fmt.Printf("Part one answer: %d\n", partOneAnswer)

	board = drawBoard(lines, height, width, padding, start)
	partTwoAnswer := partTwo(board, offset)
	fmt.Printf("Part two answer: %d\n", partTwoAnswer)
}

func partOne(board [][]string, offset int, floor int) int {
	spout := [2]int{500 - offset, 0}
	overTheEdge := false
	i := 0
	for overTheEdge == false {
		grain := spout
		atRest := false
		for atRest == false {
			nextGrain := findNextPosition(grain, board)
			if nextGrain != grain {
				grain = nextGrain
				continue
			}
			atRest = true
			board[grain[1]][grain[0]] = "o"
			if grain[1] >= floor {
				overTheEdge = true
				break
			}
		}
		i++
	}
	return i - 1
}

func partTwo(board [][]string, offset int) int {
	spout := [2]int{500 - offset, 0}
	overTheEdge := false
	i := 0
	for overTheEdge == false {
		grain := spout
		atRest := false
		for atRest == false {
			nextGrain := findNextPosition(grain, board)
			if nextGrain != grain {
				grain = nextGrain
				continue
			}
			atRest = true
			board[grain[1]][grain[0]] = "o"
			if grain[1] == 0 {
				overTheEdge = true
				break
			}
		}
		i++
	}
	return i
}

func drawBoard(lines [][][2]int, height int, width int, padding int, start int) [][]string {
	board := [][]string{}
	for i := 0; i <= height; i++ {
		row := make([]string, width+padding)
		for i, _ := range row {
			row[i] = "."
		}
		board = append(board, row)
	}

	// Add floor
	row := make([]string, width+padding)
	for i, _ := range board[0] {
		row[i] = "#"
	}
	board = append(board, row)

	offset := start - (padding / 2)
	for _, line := range lines {
		for i := 0; i < len(line)-1; i++ {
			first := line[i]
			second := line[i+1]
			for x := first[0]; x <= second[0]; x++ {
				board[first[1]][x-offset] = "#"
			}
			for y := first[1]; y <= second[1]; y++ {
				board[y][first[0]-offset] = "#"
			}
			for x := second[0]; x <= first[0]; x++ {
				board[first[1]][x-offset] = "#"
			}
			for y := second[1]; y <= first[1]; y++ {
				board[y][first[0]-offset] = "#"
			}
		}
	}
	return board
}

func findNextPosition(grain [2]int, board [][]string) [2]int {
	if board[grain[1]+1][grain[0]] == "." {
		grain = [2]int{grain[0], grain[1] + 1}
	} else if board[grain[1]+1][grain[0]-1] == "." {
		grain = [2]int{grain[0] - 1, grain[1] + 1}
	} else if board[grain[1]+1][grain[0]+1] == "." {
		grain = [2]int{grain[0] + 1, grain[1] + 1}
	}
	return grain
}

func isFreeSpace(takenSpaces [][2]int, space [2]int) bool {
	for _, position := range takenSpaces {
		if space == position {
			return false
		}
	}
	return true
}
