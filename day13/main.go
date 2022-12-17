package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pairs = [][]value{}

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

	currentPair := []value{}
	for sc.Scan() {
		if len(sc.Text()) == 0 {
			pairs = append(pairs, currentPair)
			currentPair = []value{}
		} else {
			packet := parseLine(sc.Text())
			currentPair = append(currentPair, value{array: packet})
		}
	}

	partOneAnswer := partOne(pairs)
	fmt.Printf("Part one answer: %d\n", partOneAnswer)
	partTwoAnswer := partTwo(pairs)
	fmt.Printf("Part two answer: %d\n", partTwoAnswer)
}

func parseLine(line string) *[]value {
	current := &[]value{}
	stack := value{}

	for i, char := range line {
		if char >= 48 && char <= 57 {
			// If it's a second digit, skip. We already got that one.
			if line[i-1] >= 49 && line[i-1] <= 57 {
				continue
			}

			// If it's two digits it's a 10.
			if line[i+1] >= 48 && line[i+1] <= 57 {
				*current = append(*current, value{number: 10})
				continue
			}

			*current = append(*current, value{number: int(char) - 48})
		}
		if char == 91 { // Opening bracket
			if stack.array == nil {
				stack.array = &[]value{}
			}
			prev := append(*stack.array, value{array: current})
			stack.array = &prev
			current = &[]value{}
		}
		if char == 93 { // Closing bracket
			if len(*stack.array) == 0 {
				continue
			}
			prev := (*stack.array)[len(*stack.array)-1]
			poppedStack := (*stack.array)[:len(*stack.array)-1]
			stack.array = &poppedStack
			newCurrent := append(*prev.array, value{array: current})
			current = &newCurrent
		}
	}
	return current
}

func partOne(pairs [][]value) int {
	correctPairSum := 0
	for i, pair := range pairs {
		correctlySorted := compare(pair[0], pair[1])
		if correctlySorted == 1 {
			correctPairSum += i + 1
		}
	}
	return correctPairSum
}

func partTwo(pairs [][]value) int {
	// Combine pairs and add divider packets
	dividerOne := value{array: &[]value{value{array: &[]value{value{number: 2}}}}}
	dividerTwo := value{array: &[]value{value{array: &[]value{value{number: 6}}}}}
	totalList := []value{dividerOne, dividerTwo}
	for _, pair := range pairs {
		totalList = append(totalList, pair[0])
		totalList = append(totalList, pair[1])
	}

	sorted := false
	dividerOnePos := 0
	dividerTwoPos := 1
	for sorted == false {
		sorted = true
		for i := 0; i < len(totalList)-1; i += 1 {
			first := totalList[i]
			second := totalList[i+1]
			if compare(totalList[i], totalList[i+1]) == -1 {
				totalList[i] = second
				totalList[i+1] = first
				sorted = false
				if i == dividerOnePos {
					dividerOnePos++
				}
				if i == dividerTwoPos {
					dividerTwoPos++
				}
			}
		}
	}
	return (dividerOnePos + 1) * (dividerTwoPos + 1)
}

func compare(left value, right value) int {
	if left.array == nil && right.array == nil {
		if left.number > right.number {
			return -1
		}
		if left.number == right.number {
			return 0
		}
		if left.number < right.number {
			return 1
		}
	} else if left.array == nil {
		left = value{array: &[]value{value{number: left.number}}}
	} else if right.array == nil {
		right = value{array: &[]value{value{number: right.number}}}
	}
	minLength := len(*left.array)
	rightLarger := 1
	if len(*left.array) > len(*right.array) {
		minLength = len(*right.array)
		rightLarger = -1
	}
	if len(*left.array) == len(*right.array) {
		rightLarger = 0
	}
	for i := 0; i < minLength; i++ {
		correctOrder := compare((*left.array)[i], (*right.array)[i])
		if correctOrder == 1 || correctOrder == -1 {
			return correctOrder
		}
	}
	return rightLarger
}

func valueToString(valueToPrint value) string {
	if valueToPrint.array != nil {
		result := []string{}
		for i := 0; i < len(*valueToPrint.array); i++ {
			result = append(result, valueToString((*valueToPrint.array)[i]))
		}
		return fmt.Sprintf("%v", result)
	}
	return fmt.Sprint(valueToPrint.number)
}

type value struct {
	number int
	array  *[]value
}
