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
	fmt.Println("Day 2 2023 - Part I\n")

	// Read input file
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}
	defer file.Close()

	colorConfig := []int{12, 13, 14} // Pre-defined color configuration - red, green, blue
	sum := 0                         // Sum of all possible games

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()

		// Get the game number as an int
		gameNumber, err := strconv.Atoi(strings.Replace(strings.Split(lineText, ":")[0], "Game ", "", 1))
		if err != nil {
			log.Fatalf("Error converting game number to int: %v", err)
		}

		// Split the game into sets and remove boilerplate
		sets := strings.Split(lineText, "; ")
		sets[0] = strings.Replace(sets[0], fmt.Sprintf("Game %d: ", gameNumber), "", 1)

		// Let's assume every game is valid as a standard
		validGame := true

		// Iterate over every set in a game
		for _, set := range sets {
			// Split the set into the different cubes taken out
			cubes := strings.Split(set, ", ")

			// Iterate over all cubes and check if that's possible - if not, the complete game is invalid
			for _, cube := range cubes {
				switch {
				case strings.Contains(cube, "red"):
					cube = strings.ReplaceAll(cube, " red", "")
					if num, err := strconv.Atoi(cube); err != nil {
						log.Fatalf("Error converting string to int: %v", err)
					} else if num > colorConfig[0] {
						validGame = false
						break
					}

				case strings.Contains(cube, "green"):
					cube = strings.ReplaceAll(cube, " green", "")
					if num, err := strconv.Atoi(cube); err != nil {
						log.Fatalf("Error converting string to int: %v", err)
					} else if num > colorConfig[1] {
						validGame = false
						break
					}

				case strings.Contains(cube, "blue"):
					cube = strings.ReplaceAll(cube, " blue", "")
					if num, err := strconv.Atoi(cube); err != nil {
						log.Fatalf("Error converting string to int: %v", err)
					} else if num > colorConfig[2] {
						validGame = false
						break
					}
				}
			}

			// If the game is invalid we can skip the remaining sets
			if !validGame {
				break
			}
		}

		// If the game is valid add the game number to the sum
		if validGame {
			sum += gameNumber
		}
	}

	// Print result
	fmt.Printf("The overall sum is: %d\n", sum)
}
