package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	signalIdx   = 0
	messageIdx  = 0
	inputStream []byte
)

func init() {
	file, err := os.ReadFile("./input.txt")
	check(err)
	inputStream = file
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne() {
	for i := range inputStream {
		tmpBuf := inputStream[i : i+4]

		isUnique := true
		for _, c := range tmpBuf {
			numRepetitions := strings.Count(string(tmpBuf), string(c))
			if numRepetitions > 1 {
				isUnique = false
				break
			}
		}

		if isUnique {
			signalIdx = i + 4 // Plus 4 for the end of marker
			break
		}
	}
}

func partTwo() {
	for i := range inputStream {
		tmpBuf := inputStream[i : i+14]

		isUnique := true
		for _, c := range tmpBuf {
			numRepetitions := strings.Count(string(tmpBuf), string(c))
			if numRepetitions > 1 {
				isUnique = false
				break
			}
		}

		if isUnique {
			messageIdx = i + 14 // Plus 4 for the end of marker
			break
		}
	}
}

func main() {
	partOne()
	fmt.Printf("Part One (Signal Marker): %v\n", signalIdx)

	partTwo()
	fmt.Printf("Part Two (Message Marker): %v\n", messageIdx)
}
