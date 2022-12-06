package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func firstUniqueSeries(code string, count int) int {
  lastLetters := []rune{}
  for index, newLetter := range code {
    lastLetters = append(lastLetters, newLetter)
    if len(lastLetters) < count {
      continue
    }

    seen := false

    for lastLetterIndex, letter := range lastLetters {
      for i := lastLetterIndex + 1; i < count; i++ {
        if lastLetters[i] == letter {
          seen = true
        }
      }
    }

    if seen {
      lastLetters = lastLetters[1:]
    } else {
      return index + 1
    }
  }
  return 0
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

  var code string
  for sc.Scan() {
    code = sc.Text()
  }

  partOneAnswer := firstUniqueSeries(code, 4)
  fmt.Printf("Part one answer: %d\n", partOneAnswer)
  partTwoAnswer := firstUniqueSeries(code, 14)
  fmt.Printf("Part two answer: %d\n", partTwoAnswer)
}
