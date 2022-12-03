package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

func partOne(games [][]string) int {
  score := 0
  for _, game := range games {
    switch game[1] {
      case "X":
        score += 1
        switch game[0] {
          case "A":
            score += 3
          case "C":
            score += 6
        }
      case "Y":
        score += 2
        switch game[0] {
          case "A":
            score += 6
          case "B":
            score += 3
        }
      case "Z":
        score += 3
        switch game[0] {
          case "B":
            score += 6
          case "C":
            score += 3
        }
    }
  }

  return score
}

func partTwo(games [][]string) int {
  score := 0
  for _, game := range games {
    switch game[1] {
      case "X":
        switch game[0] {
          case "A":
            score += 3
          case "B":
            score += 1
          case "C":
            score += 2
        }
      case "Y":
        score += 3
        switch game[0] {
          case "A":
            score += 1
          case "B":
            score += 2
          case "C":
            score += 3
        }
      case "Z":
        score += 6
        switch game[0] {
          case "A":
            score += 2
          case "B":
            score += 3
          case "C":
            score += 1
        }
    }
  }

  return score
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
  games := make([][]string, 0)

  for sc.Scan() {
    games = append(games, strings.Fields(sc.Text()))
  }
  
  partOneScore := partOne(games)
  fmt.Printf("Part one score: %d\n", partOneScore)
  partTwoScore := partTwo(games)
  fmt.Printf("Part two score: %d\n", partTwoScore)
}
