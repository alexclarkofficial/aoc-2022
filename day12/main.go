package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var forestMap = [][]rune{}
var searching = true
var start = [2]int{}
var end = [2]int{}

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

	as := [][2]int{}
	rowIndex := 0
	for sc.Scan() {
		row := []rune{}
		for columnIndex, elevation := range sc.Text() {
			if elevation == 83 { // if 'S'
				start = [2]int{columnIndex, rowIndex}
				row = append(row, 97) // value of 'a'
				continue
			}
			if elevation == 69 { // if 'E'
				end = [2]int{columnIndex, rowIndex}
				row = append(row, 123) // value of 'z' + 1
				continue
			}
			if elevation == 97 { // if 'a'
				as = append(as, [2]int{columnIndex, rowIndex})
			}
			row = append(row, elevation)
		}
		forestMap = append(forestMap, row)
		rowIndex++
	}

	partOneStartPaths := [][][2]int{}
	partOneStartPaths = append(partOneStartPaths, [][2]int{start})

	partTwoStartPaths := [][][2]int{}
	partTwoStartPaths = append(partTwoStartPaths, [][2]int{start})
	for _, a := range as {
		partTwoStartPaths = append(partTwoStartPaths, [][2]int{a})
	}

	partOneAnswer := findShortestPath(partOneStartPaths)
	fmt.Printf("Part one answer: %d\n", partOneAnswer)
	partTwoAnswer := findShortestPath(partTwoStartPaths)
	fmt.Printf("Part two answer: %d\n", partTwoAnswer)
}

func findShortestPath(possiblePaths [][][2]int) int {
	seenPoints := map[string]bool{}
	rounds := 0

	searching = true
	for searching {
		newPaths := [][][2]int{}
		for _, path := range possiblePaths {
			headX := int(path[0][0])
			headY := int(path[0][1])
			elevation := forestMap[headY][headX]

			if headX+1 < len(forestMap[0]) {
				newPaths = appendToNewPaths(headX+1, headY, elevation, seenPoints, path, newPaths)
			}
			if headX-1 >= 0 {
				newPaths = appendToNewPaths(headX-1, headY, elevation, seenPoints, path, newPaths)
			}
			if headY+1 < len(forestMap) {
				newPaths = appendToNewPaths(headX, headY+1, elevation, seenPoints, path, newPaths)
			}
			if headY-1 >= 0 {
				newPaths = appendToNewPaths(headX, headY-1, elevation, seenPoints, path, newPaths)
			}
		}
		possiblePaths = newPaths
		rounds++
	}

	return rounds
}

func appendToNewPaths(headX int, headY int, currentElevation rune, seenPoints map[string]bool, currentPath [][2]int, newPaths [][][2]int) [][][2]int {
	nextPoint := [2]int{headX, headY}
	nextPointElevation := forestMap[headY][headX]
	if nextPointElevation-currentElevation <= 1 {
		if nextPoint == end {
			searching = false
		}
		if !pointSeen(seenPoints, nextPoint) {
			newPath := append([][2]int{nextPoint}, currentPath...)
			newPaths = append(newPaths, newPath)
		}
	}
	return newPaths
}

func pointSeen(seenPoints map[string]bool, point [2]int) bool {
	key := fmt.Sprintf("%d.%d", point[0], point[1])
	if seenPoints[key] != true {
		seenPoints[key] = true
		return false
	}
	return true
}
