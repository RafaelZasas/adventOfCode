package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// var fileName = "input.txt"

var fileName = "example.txt"

var fileBuffer []byte

func init() {

	start := time.Now()
	defer func() {
		fmt.Println("Init took: ", time.Since(start))
	}()

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	// read file into buffer and split by two new lines
	fileBuffer, err := io.ReadAll(file)
	check(err)

	rows := strings.Split(string(fileBuffer), "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}

	}
}

func partOne() {

	start := time.Now()
	defer func() {
		fmt.Println("Part one took: ", time.Since(start))
	}()
}

func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("Part two took: ", time.Since(start))
	}()
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
