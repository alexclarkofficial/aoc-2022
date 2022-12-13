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

	monkies := []monkey{}
	globalDiviser := 1
	var currentMonkey monkey
	for sc.Scan() {
		var opOp string
		var opNumber string
		var divisor int
		var trueReceiver int
		var falseReceiver int

		prompt := strings.Split(sc.Text(), ": ")
		if prompt[0] == "  Starting items" {
			itemsStrings := strings.Split(prompt[1], ", ")
			items := []int{}
			for _, val := range itemsStrings {
				item, _ := strconv.Atoi(val)
				items = append(items, item)
			}
			currentMonkey = monkey{items: items}
		}

		_, notFound := fmt.Sscanf(sc.Text(), "  Operation: new = old %s %s", &opOp, &opNumber)
		if notFound == nil {
			currentMonkey.opOp = opOp
			currentMonkey.opNumber = opNumber
		}

		_, notFound = fmt.Sscanf(sc.Text(), "  Test: divisible by %d", &divisor)
		if notFound == nil {
			currentMonkey.divisor = divisor
			globalDiviser = globalDiviser * divisor
		}

		_, notFound = fmt.Sscanf(sc.Text(), "    If true: throw to monkey %d", &trueReceiver)
		if notFound == nil {
			currentMonkey.trueReceiver = trueReceiver
		}

		_, notFound = fmt.Sscanf(sc.Text(), "    If false: throw to monkey %d", &falseReceiver)
		if notFound == nil {
			currentMonkey.falseReceiver = falseReceiver
			monkies = append(monkies, currentMonkey)
		}
	}

	partOneAnswer := monkeyPlay(monkies, 20, 3, globalDiviser)
	fmt.Printf("Part one answer: %d\n", partOneAnswer)
	partTwoAnswer := monkeyPlay(monkies, 10000, 1, globalDiviser)
	fmt.Printf("Part two answer: %d\n", partTwoAnswer)
}

func monkeyPlay(monkies []monkey, rounds int, roundDivisor int, globalDiviser int) int {
	gameMonkies := []monkey{}
	for _, monkey := range monkies {
		gameMonkies = append(gameMonkies, monkey)
	}

	for i := 0; i < rounds; i++ {
		for monkeyIndex, currentMonkey := range gameMonkies {
			for _, item := range currentMonkey.items {
				currentMonkey.handledCount++
				var opNumber int
				var itemVal int
				var receiverIndex int

				switch currentMonkey.opNumber {
				case "old":
					opNumber = item
				default:
					opNumber, _ = strconv.Atoi(currentMonkey.opNumber)
				}

				switch currentMonkey.opOp {
				case "+":
					itemVal = item + opNumber
				case "*":
					itemVal = item * opNumber
				}

				itemVal = itemVal / roundDivisor
				itemMod := itemVal % currentMonkey.divisor
				if itemMod == 0 {
					receiverIndex = currentMonkey.trueReceiver
				} else {
					receiverIndex = currentMonkey.falseReceiver
				}

				receiver := gameMonkies[receiverIndex]
				receiver.items = append(receiver.items, itemVal%(roundDivisor*globalDiviser))
				gameMonkies[receiverIndex] = receiver
			}
			currentMonkey.items = []int{}
			gameMonkies[monkeyIndex] = currentMonkey
		}
	}

	return twoBiggestMultiple(gameMonkies)
}

func twoBiggestMultiple(monkies []monkey) int {
	biggest := 0
	secondBiggest := 0
	for _, currentMonkey := range monkies {
		handledCount := currentMonkey.handledCount
		if handledCount > biggest {
			secondBiggest = biggest
			biggest = handledCount
			continue
		} else if handledCount > secondBiggest {
			secondBiggest = handledCount
			continue
		}
	}
	return biggest * secondBiggest
}

type monkey struct {
	items         []int
	opOp          string
	opNumber      string
	divisor       int
	trueReceiver  int
	falseReceiver int
	handledCount  int
}
