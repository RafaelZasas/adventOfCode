package q2

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
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
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

func part2() int {
	var h int
	var d int
	var aim int

	for _, inst := range instructions {
		switch inst.direction {
		case "forward":
			h += inst.dist
			d += aim * inst.dist
		case "down":
			aim += inst.dist
		case "up":
			aim -= inst.dist
		}
	}
	return h * d
}
