package main

import (
	"bufio"
	"fmt"
	"os"
)

var total int = 0

var handScores = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func partOne() {
	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		roundChoices := scanner.Text()

		opponentScore := handScores[string(roundChoices[0])]
		myScore := handScores[string(roundChoices[2])] // second char is a space

		total += myScore
        
        // not taking into account any losing scenarios since score wouldnt change

		if myScore == 1 && opponentScore == 3 { // if I have rock and they have scissors
			total += 6
		}

        if myScore == 2 && opponentScore == 1 { // if I have paper and they have rock
            total += 6
        }

        if myScore == 3 && opponentScore == 2 { // if I have scissors and they have paper

            total += 6
        } 

		if myScore == opponentScore {
			total += 3
		}

	}
}


func partTwo() {
    total = 0 // reset total after part one

	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		roundChoices := scanner.Text()

		opponentScore := handScores[string(roundChoices[0])]
		outcome := string(roundChoices[2]) // whether I need to win lose or draw

		if outcome == "Z" { // if I need to win
            total += 6 // automatic +6 because I will win
            if opponentScore == 1 { // opponent has rock, I need paper
                total += 2
            }

            if opponentScore == 2 { // opponent has paper, I need scissors
                total += 3
            }

            if opponentScore == 3 { // opponent has scissors, I need rock
                total += 1
            }
		}

        if outcome == "Y" { // If I need to draw 
            total += 3 // automatic +3 since I will draw
            total += opponentScore // since I will always have the same hand as my opponent
        }

        if outcome == "X" { // if I need to lose
            if opponentScore == 1 { // opponent has rock, I need scissors
                total += 3
            }

            if opponentScore == 2 { // opponent has paper, I need rock
                total += 1
            }

            if opponentScore == 3 { // opponent has scissors, I need paper
                total += 2
            }
        } 

	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
    partOne()
	fmt.Printf("Part 1 (Total Score): %v\n", total)

    partTwo()
	fmt.Printf("Part 2 (Total Score): %v\n", total)
}
