package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	dist      int
}

var instructions []instruction

func init() {

	file, err := os.Open("./q2.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var i instruction
		dist, err := strconv.Atoi(s[1])
		check(err)

		i.direction, i.dist = s[0], dist
		instructions = append(instructions, i)
	}
	file.Close()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println(part1())
}

func part1() int {
	var h int
	var d int

	for _, inst := range instructions {
		switch inst.direction {
		case "forward":
			h += inst.dist
		case "down":
			d += inst.dist
		case "up":
			d -= inst.dist
		}
	}
	return d * h
}
