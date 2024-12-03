package main

import (
	"bufio"
	"os"
)

var grid [][]string

func init() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var chars = []string{}
		for _, char := range line {
			chars = append(chars, string(char))
		}
		grid = append(grid, chars)
	}
}

func q1() {

}

func main() {
	q1()
}
