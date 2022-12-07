package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getDirSizeMap(commands []string) map[string]int {
	currentDir := []string{}
	dirSizeMap := map[string]int{}

	for _, command := range commands {
		dir := ""
		fileSize := 0
		fmt.Sscanf(command, "$ cd %s", &dir)
		fmt.Sscanf(command, "%d %s", &fileSize)

		if dir != "" {
			if dir == "/" {
				continue
			}
			if dir == ".." {
				currentDir = currentDir[:len(currentDir)-1]
				continue
			}
			currentDir = append(currentDir, dir)
		}

		if fileSize != 0 {
			for i := 0; i <= len(currentDir); i++ {
				dirString := strings.Join(currentDir[:len(currentDir)-i], "/")
				if _, ok := dirSizeMap[dirString]; ok {
					dirSizeMap[dirString] += fileSize
				} else {
					dirSizeMap[dirString] = fileSize
				}
			}
		}
	}

	return dirSizeMap
}

func sizeOfSmallDir(dirSizeMap map[string]int) int {
	total := 0
	for _, size := range dirSizeMap {
		if size <= 100000 {
			total += size
		}
	}
	return total
}

func sizeOfSmallestDirToFree(dirSizeMap map[string]int) int {
	totalDisk := 70000000
	newProgram := 30000000

	needed := newProgram - (totalDisk - dirSizeMap[""])
	smallestPossible := totalDisk
	for _, size := range dirSizeMap {
		if size >= needed && smallestPossible > size {
			smallestPossible = size
		}
	}

	return smallestPossible
}

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

	commands := []string{}
	for sc.Scan() {
		commands = append(commands, sc.Text())
	}

	dirSizeMap := getDirSizeMap(commands)

	partOneAnswer := sizeOfSmallDir(dirSizeMap)
	fmt.Printf("Part one answer: %d\n", partOneAnswer)
	partTwoAnswer := sizeOfSmallestDirToFree(dirSizeMap)
	fmt.Printf("Part two answer: %d\n", partTwoAnswer)
}
