package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2 2023 - Part II\n")

	// Read input file
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}
	defer file.Close()

	sum := 0 // Sum of all possible games

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()

		// Prefix of line ("Game xy")
		prefix := strings.Split(lineText, ":")[0]

		// Split the game into sets and remove boilerplate
		sets := strings.Split(lineText, "; ")
		sets[0] = strings.Replace(sets[0], fmt.Sprintf("%s: ", prefix), "", 1)

		// Cache the highest values
		highestRed := math.MinInt
		highestGreen := math.MinInt
		highestBlue := math.MinInt

		// Iterate over every set in a game
		for _, set := range sets {
			// Split the set into the different cubes taken out
			cubes := strings.Split(set, ", ")

			// Iterate over all cubes and check for highest number of cubes in a game
			for _, cube := range cubes {
				switch {
				case strings.Contains(cube, "red"):
					cube = strings.ReplaceAll(cube, " red", "")
					if num, err := strconv.Atoi(cube); err != nil {
						log.Fatalf("Error converting string to int: %v", err)
					} else if num > highestRed {
						highestRed = num
					}

				case strings.Contains(cube, "green"):
					cube = strings.ReplaceAll(cube, " green", "")
					if num, err := strconv.Atoi(cube); err != nil {
						log.Fatalf("Error converting string to int: %v", err)
					} else if num > highestGreen {
						highestGreen = num
					}

				case strings.Contains(cube, "blue"):
					cube = strings.ReplaceAll(cube, " blue", "")
					if num, err := strconv.Atoi(cube); err != nil {
						log.Fatalf("Error converting string to int: %v", err)
					} else if num > highestBlue {
						highestBlue = num
					}
				}
			}
		}

		// Add the power of all cubes to the sum
		sum += highestRed * highestGreen * highestBlue
	}

	// Print result
	fmt.Printf("The overall sum is: %d\n", sum)
}
