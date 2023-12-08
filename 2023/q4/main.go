package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var myCards [][]int
var winningCards [][]int

func init() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Eg. of input file:
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")[1]
		winningNumbers := strings.Split(line, "|")[0]
		myNumbers := strings.Split(line, "|")[1]

		myCards = append(myCards, convertToIntArray(myNumbers))
		winningCards = append(winningCards, convertToIntArray(winningNumbers))
	}
}

// q1- Calculate the total points of all the cards (rows of input)
// The first matching number is worth one point, subsequent matches double the points.
func q1() {
	totalPoints := 0

outerLoop:
	for i, card := range myCards {
		var matches int
		for _, number := range card {
			if contains(winningCards[i], number) {
				matches++
			}
		}
		if matches < 0 {
			continue outerLoop
		}

		points := math.Pow(2, float64(matches-1))
		totalPoints += int(points)
	}

	fmt.Printf("Q1) Total points: %d\n", totalPoints)
}

// q2 - Calculate the total number of scratch cards at the end of the game.
// Each matching number generates a copy of the next card in the pile.
// eg. If card 1 has 2 matches, then card 2 and card 3 is coppied.
// If card 2 has 3 matches, then card 3, 4 and 5 is coppied.
// this function will compute the cards recursively.
func q2() (totalCards int) {
	scratchCards := make(map[int]int, len(myCards))

	for x := range myCards {
		scratchCards[x] = 1
	}

	for i, card := range myCards {
		var count int
		for _, number := range card {
			if contains(winningCards[i], number) {
				count++
			}
		}

		for j := i + 1; j < i+count+1; j++ {
			scratchCards[j] += 1
		}

		fmt.Println(scratchCards)

	}

	for _, v := range scratchCards {
		totalCards += v
	}

	return

}

func main() {
	q1()

	totalCards := q2()
	fmt.Printf("Q2) Total number of cards: %d\n", totalCards)
}

// Helper functions

func contains(numbers []int, number int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}

func convertToIntArray(numbers string) []int {
	var result []int
	for _, number := range strings.Split(numbers, " ") {
		if number == "" {
			continue
		}
		n, err := strconv.Atoi(strings.Trim(number, " "))
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}
	return result
}
