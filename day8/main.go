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

	forest := [][]int{}
	for sc.Scan() {
    row := []int{}
    for _, tree := range sc.Text() {
      row = append(row, int(tree - 48))
    }

		forest = append(forest, row)
	}


	partOneAnswer := countVisibleTrees(forest)
	fmt.Printf("Part one answer: %d\n", partOneAnswer)
  partTwoAnswer := maxViewDepth(forest)
	fmt.Printf("Part two answer: %d\n", partTwoAnswer)
}

func countVisibleTrees(forest [][]int) int {
  count := 0

  height := len(forest)
  width := len(forest[0])
  for rowIndex, row := range forest {
    for columnIndex, tree := range row {
      if columnIndex == 0 || rowIndex == 0 || rowIndex == height - 1 || columnIndex == width - 1 {
        count++
        continue
      }

      tallest := true
      for i := 0; i < rowIndex; i++ {
        if forest[i][columnIndex] >= tree {
          tallest = false
        }
      }
      if tallest {
        count++
        continue
      }

      tallest = true
      for i := rowIndex + 1; i < height; i++ {
        if forest[i][columnIndex] >= tree {
          tallest = false
        }
      }
      if tallest {
        count++
        continue
      }

      tallest = true
      for i := 0; i < columnIndex; i++ {
        if forest[rowIndex][i] >= tree {
          tallest = false
        }
      }
      if tallest {
        count++
        continue
      }

      tallest = true
      for i := columnIndex + 1; i < width; i++ {
        if forest[rowIndex][i] >= tree {
          tallest = false
        }
      }
      if tallest {
        count++
        continue
      }
    }
  }
  return count
}

func maxViewDepth(forest [][]int) int {
  maxView := 0

  height := len(forest)
  width := len(forest[0])
  for rowIndex, row := range forest {
    for columnIndex, tree := range row {
      if columnIndex == 0 || rowIndex == 0 || rowIndex == height - 1 || columnIndex == width - 1 {
        continue
      }
      upCount := 0
      downCount := 0
      leftCount := 0
      rightCount := 0

      for i := rowIndex - 1; i >= 0; i-- {
        upCount++
        if forest[i][columnIndex] >= tree {
          break
        }
      }

      for i := rowIndex + 1; i < height; i++ {
        downCount++
        if forest[i][columnIndex] >= tree {
          break
        }
      }

      for i := columnIndex - 1; i >= 0; i-- {
        leftCount++
        if forest[rowIndex][i] >= tree {
          break
        }
      }

      for i := columnIndex + 1; i < width; i++ {
        rightCount++
        if forest[rowIndex][i] >= tree {
          break
        }
      }

      treeCount := upCount * downCount * leftCount * rightCount
      if treeCount > maxView {
        maxView = treeCount
      }
    }
  }
  return maxView
}
