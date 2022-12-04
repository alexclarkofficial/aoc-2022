package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func getElves(filename string) ([][][]int, error) {
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
    return nil, err
  }
  defer file.Close()

  sc := bufio.NewScanner(file)
  pairs := make([][][]int, 0)

  for sc.Scan() {
    rawPair := strings.Split(sc.Text(), ",")
    splitPair := make([][]int, 0)
    for _, pair := range rawPair {
      split := strings.Split(pair, "-")
      start, err := strconv.Atoi(split[0])
      if err != nil {
        log.Fatal(err)
        return nil, err
      }
      end, err := strconv.Atoi(split[1])
      if err != nil {
        log.Fatal(err)
        return nil, err
      }
      splitPair = append(splitPair, []int{start, end})
    }
    pairs = append(pairs, splitPair)
  }

  return pairs, nil;
}

func findTotalOverlaps(pairs [][][]int) int {
  overlaps := 0

  for _, pair := range pairs {
    firstElf := pair[0]
    secondElf := pair[1]
    if firstElf[0] <= secondElf[0] && firstElf[1] >= secondElf[1] {
      overlaps++
    } else if firstElf[0] >= secondElf[0] && firstElf[1] <= secondElf[1] {
      overlaps++
    }
  }
  return overlaps
}

func findPartialOverlaps(pairs [][][]int) int {
  overlaps := 0

  for _, pair := range pairs {
    firstElf := pair[0]
    secondElf := pair[1]

    if firstElf[0] <= secondElf[0] && firstElf[1] >= secondElf[1] {
      overlaps++
    } else if firstElf[0] >= secondElf[0] && firstElf[1] <= secondElf[1] {
      overlaps++
    } else if firstElf[0] <= secondElf[1] && firstElf[0] >= secondElf[0] {
      overlaps++
    } else if firstElf[1] >= secondElf[0] && firstElf[1] <= secondElf[1] {
      overlaps++
    }
  }
  return overlaps
}

func main() {
  filename := os.Args[1]

  pairs, err := getElves(filename)
  if err != nil {
    log.Fatal(err)
    return
  }

  partOneAnswer := findTotalOverlaps(pairs)
  fmt.Printf("The answer for part one is: %d\n", partOneAnswer)
  partTwoAnswer := findPartialOverlaps(pairs)
  fmt.Printf("The answer for part two is: %d\n", partTwoAnswer)
}
