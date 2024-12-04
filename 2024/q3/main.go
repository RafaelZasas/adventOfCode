package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

var fileName = "input.txt"

// var fileName = "example.txt"
// var fileName = "exampleP2.txt"

var tokens = [][]byte{}
var tokensWithInstruction = [][]byte{}

func init() {
	start := time.Now()
	defer func() {
		fmt.Println("Init took: ", time.Since(start))
	}()

	// read file into memory
	file, err := os.ReadFile(fileName)
	check(err)

	//regex capture for mul(x,y)
	// mul\(\d+,\d+\)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	tokens = re.FindAll(file, -1)

	// new capture for instructions do() or don't()
	// as well as the tokens
	// do\(\) | don't\(\) | mul\(\d+,\d+\)
	re = regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)

	tokensWithInstruction = re.FindAll(file, -1)
}

func partOne() {
	start := time.Now()
	defer func() {
		fmt.Println("Part one took: ", time.Since(start))
	}()

	sum := 0

	for _, token := range tokens {
		// fmt.Println(string(token))
		var x, y int
		fmt.Sscanf(string(token), "mul(%d,%d)", &x, &y)
		sum += x * y
	}

	fmt.Println("Part One Sum: ", sum)
}

// at the beginning of the tokens, multiplications is enabled
func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("Part two took: ", time.Since(start))
	}()

	sum := 0
	multiplicationEnabled := true

	for _, token := range tokensWithInstruction {
		// fmt.Println(string(token))
		switch string(token) {
		case "do()":
			multiplicationEnabled = true
		case "don't()":
			multiplicationEnabled = false
		default:
			if multiplicationEnabled {
				var x, y int
				fmt.Sscanf(string(token), "mul(%d,%d)", &x, &y)
				sum += x * y
			}
		}
	}

	fmt.Println("Part Two Sum: ", sum)
}

func main() {
	partOne()
	partTwo()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
