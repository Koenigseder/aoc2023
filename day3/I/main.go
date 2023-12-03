package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Day 3 2023 - Part II\n")

	// Read the input as matrix
	matrix := readInputAsMatrix()

	sum := 0

	// Iterate over each row in the matrix
	for i, row := range matrix {
		var number string // Cache the current number

		// Iterate over each element in a row
		for j, element := range row {
			// Check if the element is a digit
			if unicode.IsDigit([]rune(element)[0]) {
				number += element

				// Check if end is reached or the next char is not a digit -> number found
				if j+1 == len(row) || !unicode.IsDigit([]rune(row[j+1])[0]) {
					// Check if a symbol is near the number
					foundSymbol := checkForSurroundingSymbol(matrix, i, j-len(number)+1, j)

					// In case there is a number, add it to the sum
					if foundSymbol {
						numberInt, err := strconv.Atoi(number)
						if err != nil {
							log.Fatalf("Error converting string to number: %v", err)
						}

						sum += numberInt
					}
				}
			} else {
				number = "" // Reset number string
			}
		}
	}

	// Print result
	fmt.Printf("The overall sum is: %d\n", sum)
}

// Read the input as matrix
func readInputAsMatrix() [][]string {
	var matrix [][]string

	// Read input file
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		matrix = append(matrix, row)
	}

	return matrix
}

// Check for a surrounding symbol
func checkForSurroundingSymbol(matrix [][]string, row, startIndex, endIndex int) bool {
	if startIndex-1 >= 0 {
		leftChar := matrix[row][startIndex-1]
		if !unicode.IsDigit([]rune(leftChar)[0]) && leftChar != "." {
			return true
		}
	}

	if endIndex+1 < len(matrix[row]) {
		rightChar := matrix[row][endIndex+1]
		if !unicode.IsDigit([]rune(rightChar)[0]) && rightChar != "." {
			return true
		}
	}

	for i := startIndex - 1; i <= endIndex+1; i++ {
		if i < 0 || i > len(matrix[row]) {
			continue
		}

		if row-1 >= 0 && i < len(matrix[row]) {
			charToCheck := matrix[row-1][i]
			if !unicode.IsDigit([]rune(charToCheck)[0]) && charToCheck != "." {
				return true
			}
		}

		if row+1 < len(matrix) && i < len(matrix[row]) {
			charToCheck := matrix[row+1][i]
			if !unicode.IsDigit([]rune(charToCheck)[0]) && charToCheck != "." {
				return true
			}
		}
	}

	return false
}
