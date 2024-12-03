package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func q1() {
	fmt.Println()
}

func main() {
	q1()
}
