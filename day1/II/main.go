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
	fmt.Println("Day 1 2023 - Part II\n")

	// Read input file
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}
	defer file.Close()

	sum := 0 // Final sum of all numbers

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		firstInt := -1
		lastInt := -1

		// Replace words with numbers
		lineText = replaceWordsWithNumbers(lineText)

		// Loop over each and every char in a line and check if it is a number
		for _, char := range lineText {
			if num, err := strconv.Atoi(fmt.Sprintf("%c", char)); err == nil {
				// Check if firstInt is already filled with a valid value
				if firstInt == -1 {
					firstInt = num
					continue
				}

				// Set lastInt to the next number's value - in the end you have the last number's value
				lastInt = num
			}
		}

		// In case there is only one number, set the lastInt to the value of firstInt's
		if lastInt == -1 {
			lastInt = firstInt
		}

		// Add value of combined numbers to sum
		sum += firstInt*10 + lastInt
	}

	// Check for errors during scanning the file
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	// Print result
	fmt.Printf("The overall sum is: %d\n", sum)
}

// Function to replace words with numbers
// Place numbers between the starting and ending letters in case other words use them
// Example: twone -> 21
func replaceWordsWithNumbers(text string) string {
	text = strings.ReplaceAll(text, "one", "o1e")
	text = strings.ReplaceAll(text, "two", "t2o")
	text = strings.ReplaceAll(text, "three", "t3e")
	text = strings.ReplaceAll(text, "four", "f4r")
	text = strings.ReplaceAll(text, "five", "f5e")
	text = strings.ReplaceAll(text, "six", "s6x")
	text = strings.ReplaceAll(text, "seven", "s7n")
	text = strings.ReplaceAll(text, "eight", "e8t")
	text = strings.ReplaceAll(text, "nine", "n9e")

	return text
}
