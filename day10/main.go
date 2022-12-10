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

	cycle := 1
	x := 1

	cycleOfInterestTotal := 0
	crdLine := []string{}
	crdDisplay := [][]string{}

	for sc.Scan() {
		action := ""
		mag := 0

		fmt.Sscanf(sc.Text(), "%s %d", &action, &mag)

		crdLine = calcCRDLine(crdLine, x, cycle)
		cycle += 1

		cycleOfInterestTotal = addCycleOfInterest(cycle, x, cycleOfInterestTotal)
		crdDisplay, crdLine = resetCRDLine(crdDisplay, crdLine, cycle)

		if mag != 0 {
			crdLine = calcCRDLine(crdLine, x, cycle)
			cycle += 1
			x += mag

			cycleOfInterestTotal = addCycleOfInterest(cycle, x, cycleOfInterestTotal)
			crdDisplay, crdLine = resetCRDLine(crdDisplay, crdLine, cycle)
		}
	}

	fmt.Printf("Part one answer: %d\n", cycleOfInterestTotal)
	fmt.Printf("Part two answer: \n")
	for _, line := range crdDisplay {
		log.Println(line)
	}
}

func addCycleOfInterest(cycle int, x int, currentTotal int) int {
	cyclesOfInterest := []int{20, 60, 100, 140, 180, 220}

	for _, i := range cyclesOfInterest {
		if cycle == i {
			currentTotal += x * cycle
		}
	}
	return currentTotal
}

func spriteOverlap(x int, cycle int) bool {
	index := (cycle - 1) % 40
	if index == x-1 || index == x || index == x+1 {
		return true
	}
	return false
}

func calcCRDLine(crdLine []string, x int, cycle int) []string {
	if spriteOverlap(x, cycle) {
		crdLine = append(crdLine, "#")
	} else {
		crdLine = append(crdLine, ".")
	}
	return crdLine
}

func resetCRDLine(crdDisplay [][]string, crdLine []string, cycle int) ([][]string, []string) {
	if (cycle-1)%40 == 0 {
		crdDisplay = append(crdDisplay, crdLine)
		return crdDisplay, []string{}
	}
	return crdDisplay, crdLine
}
