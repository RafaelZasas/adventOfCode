package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var fileName = "input.txt"

// var fileName = "example.txt"

var fileBuffer []byte

var grid [][]string

var startingRow = 0
var startingCol = 0

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
		grid = append(grid, strings.Split(row, ""))
	}

}

func partOne() {
	start := time.Now()
	defer func() {
		fmt.Println("Part one took: ", time.Since(start))
	}()

	guardRow := 0
	guardCol := 0

	guardDirection := "up"

outer:
	for i, row := range grid {
		for j, col := range row {
			if col == "^" {
				guardRow = i
				guardCol = j
				break outer
			}
		}
	}

	startingRow = guardRow
	startingCol = guardCol

	visited := make(map[string]bool)

	for guardRow < len(grid) && guardCol < len(grid[0]) {
		// if the item in front of the guard in the directionn they are facing is a
		// wall (#) then turn by 90 degrees to the right

		// fmt.Println("Guard Row: ", guardRow)
		// fmt.Println("Guard Col: ", guardCol)
		// fmt.Println("Guard Direction: ", guardDirection)

		if guardDirection == "up" {
			if guardRow-1 < 0 || grid[guardRow-1][guardCol] == "#" {
				guardDirection = "right"
				continue
			}
		}

		if guardDirection == "right" {
			if guardCol+1 >= len(grid[0]) || grid[guardRow][guardCol+1] == "#" {
				guardDirection = "down"
				continue
			}
		}

		if guardDirection == "down" {
			if guardRow+1 >= len(grid) || grid[guardRow+1][guardCol] == "#" {
				guardDirection = "left"
				continue
			}
		}

		if guardDirection == "left" {
			if guardCol-1 < 0 || grid[guardRow][guardCol-1] == "#" {
				guardDirection = "up"
				continue
			}
		}

		// move the guard in the direction they are facing
		if guardDirection == "up" {
			guardRow--
		}

		if guardDirection == "right" {
			guardCol++
		}

		if guardDirection == "down" {
			guardRow++
		}

		if guardDirection == "left" {
			guardCol--
		}

		key := fmt.Sprintf("%d,%d", guardRow, guardCol)
		visited[key] = true

		if guardRow == 0 || guardCol == 0 || guardRow == len(grid)-1 || guardCol == len(grid[0])-1 {
			break
		}

	}

	fmt.Println("Sum: ", len(visited))

}

// Need to find out how many new ubstructions we can place in the grid to block
// the guard from escaping
// ie. the guard needs to be in a loop
// obsctructions are placed in the grid by replacing a . with a #
// and cant be placed on the guards starting position
// and cant be placed on another obstruction
func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("Part two took: ", time.Since(start))
	}()

	validObstructions := 0

	for i, row := range grid {
		for j, col := range row {
			// Skip the guard's starting position and existing obstructions
			if (i == startingRow && j == startingCol) || col == "#" {
				continue
			}

			// temp obstruction for this run
			grid[i][j] = "#"

			guardRow := startingRow
			guardCol := startingCol
			guardDirection := "up"

			visited := make(map[string]bool)
			inLoop := false

			// move the guard around the grid to check if they end up in a loop
			for steps := 0; steps < len(grid)*len(grid[0]); steps++ {

				// check if the guard has already been in this position, and facing the
				// same direction
				state := fmt.Sprintf("%d,%d,%s", guardRow, guardCol, guardDirection)
				if visited[state] {
					inLoop = true
					break
				}
				visited[state] = true

				// same stuff as part one
				if guardDirection == "up" {
					if guardRow-1 < 0 || grid[guardRow-1][guardCol] == "#" {
						guardDirection = "right"
						continue
					}
					guardRow--
				} else if guardDirection == "right" {
					if guardCol+1 >= len(grid[0]) || grid[guardRow][guardCol+1] == "#" {
						guardDirection = "down"
						continue
					}
					guardCol++
				} else if guardDirection == "down" {
					if guardRow+1 >= len(grid) || grid[guardRow+1][guardCol] == "#" {
						guardDirection = "left"
						continue
					}
					guardRow++
				} else if guardDirection == "left" {
					if guardCol-1 < 0 || grid[guardRow][guardCol-1] == "#" {
						guardDirection = "up"
						continue
					}
					guardCol--
				}

				// Stop if the guard leaves the grid bc obvs they are not in a loop
				if guardRow == 0 || guardCol == 0 || guardRow == len(grid)-1 || guardCol == len(grid[0])-1 {
					break
				}
			}

			// If the guard ends up in a loop, count this obstruction as valid
			if inLoop {
				validObstructions++
			}

			// get rid of the obstruction in this pos
			grid[i][j] = "."
		}
	}

	fmt.Println("Valid obstructions: ", validObstructions)
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
