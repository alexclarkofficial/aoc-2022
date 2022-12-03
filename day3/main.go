package main

import (
  "bufio"
  "log"
  "fmt"
  "os"
  "strings"
  "unicode"
)

func scoreItem(item rune) int {
  if unicode.IsUpper(item) {
    return int(item - 38)
  } else {
    return int(item - 96)
  }
}

func getRucks(filename string) ([]string, error) {
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
    return nil, err
  }
  defer file.Close()

  sc := bufio.NewScanner(file)
  allRucks := make([]string, 0)

  for sc.Scan() {
      ruck := sc.Text()
      allRucks = append(allRucks, ruck)
  }
  
  return allRucks, nil;
}

func getBalancedItemScore(rucks []string) int {
  score := 0
  for _, ruck := range rucks {
    compA := ruck[0:len(ruck)/2]
    compB := ruck[len(ruck)/2:]

    matches := map[rune]bool{}
    for _, item := range compA {
      if strings.ContainsRune(compB, item) {
        matches[item] = true
      }
    }
    for item, _ := range matches {
      score += scoreItem(item)
    }
  }

  return score
}

func getCommonGroupScore(rucks []string) int {
  score := 0
  for i := 0; i < len(rucks); i += 3 {
    matches := map[rune]bool{}
    for _, item := range rucks[i] {
      if strings.ContainsRune(rucks[i+1], item) && strings.ContainsRune(rucks[i+2], item) {
        matches[item] = true
      }
    }
    for item, _ := range matches {
      score += scoreItem(item)
    }
  }

  return score
}

func main() {
  filename := os.Args[1]

  allRucks, err := getRucks(filename)
  if err != nil {
    log.Fatal(err)
    return
  }

  partOneValue := getBalancedItemScore(allRucks) 
  fmt.Printf("Part one score: %d\n", partOneValue)
  partTwoValue := getCommonGroupScore(allRucks)
  fmt.Printf("Part two score: %d\n", partTwoValue)
}
