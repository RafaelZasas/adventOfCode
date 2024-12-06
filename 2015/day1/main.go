package main

import (
	"fmt"
	"os"
	"time"
)

var file = "input.txt"

// var file = "example.txt"

var fileBuffer []byte

func init() {
	start := time.Now()
	defer func() {
		fmt.Println("init took", time.Since(start))
	}()

	fileBuffer, _ = os.ReadFile(file)
}

func partOne() {
	start := time.Now()
	defer func() {
		fmt.Println("partOne took", time.Since(start))
	}()

	floor := 0
	for _, char := range fileBuffer {

		if char == '(' {
			floor++
			continue
		}

		if char == ')' {
			floor--
			continue
		}
	}
	fmt.Println(floor)
}

func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("partTwo took", time.Since(start))
	}()

	basementIdx := 0
	floor := 0

	for i, char := range fileBuffer {
		if char == '(' {
			floor++

			if floor == -1 {
				basementIdx = i + 1
				break
			}
			continue
		}

		if char == ')' {
			floor--
			if floor == -1 {
				basementIdx = i + 1
				break
			}
			continue
		}
	}

	fmt.Println(basementIdx)
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
