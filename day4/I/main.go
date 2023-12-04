package main

import (
  "bufio"
  "fmt"
  "log"
  "math"
  "os"
  "slices"
  "strings"
)

func main() {
  fmt.Println("Day 4 2023 - Part I\n")

  // Read input file
  file, err := os.Open("./day4/input.txt")
  if err != nil {
    log.Fatalf("Error while reading file: %v", err)
  }
  defer file.Close()

  sum := 0

  // Iterate over each line
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    row := scanner.Text()
    winningNumbersFound := -1

    // Remove card's prefix
    cardPrefix := strings.Split(row, ":")[0]
    row = strings.Replace(row, cardPrefix+":", "", 1)

    // Split the row into winning numbers and numbers you have
    cardSections := strings.Split(row, "|")

    // Winning numbers as slice
    winningNumbers := strings.FieldsFunc(cardSections[0], func(c rune) bool {
      return c == ' '
    })

    // Numbers as slice
    numbers := strings.FieldsFunc(cardSections[1], func(c rune) bool {
      return c == ' '
    })

    // Iterate over each number and check if it's in the winning numbers
    for _, number := range numbers {
      if slices.Contains(winningNumbers, number) {
        winningNumbersFound++ // Increase amount of winning numbers
      }
    }

    // No winning numbers -> -1
    if winningNumbersFound > -1 {
      sum += int(math.Pow(2.0, float64(winningNumbersFound)))
    }
  }

  // Print result
  fmt.Printf("The overall sum is: %d\n", sum)
}
