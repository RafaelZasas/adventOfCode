package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var grid [][]byte

func init() {

	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert the line into a byte slice
		row := []byte(line)
		grid = append(grid, row)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

}

// in q1 we need to find all the numbers in the grid which are adjacent to any character that is
// not a period (.) or another number. We then need to sum up all the numbers we find.
// Numbers are adjacent if they are directly above, below, left or right OR diagonal to a
// symbol (anything other than period or number)
func q1() {
	sum := 0
	cs := make(map[int]struct{})

	for rowIdx, row := range grid {

		for colIdx, char := range row {

			// Check if the character is a number
			if (char >= '0' && char <= '9') || char == '.' {
				continue

			}

			// Check if the character is adjacent to a symbol
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					// Calculate the new coordinates
					newR, newC := rowIdx+dr, colIdx+dc

					// Check if the new coordinates are within bounds
					if newR >= 0 && newR < len(grid) && newC >= 0 && newC < len(grid[newR]) {
						// Check if the character at the new coordinates is a digit
						if grid[newR][newC] >= '0' && grid[newR][newC] <= '9' {
							// Move to the start of the number
							for newC > 0 && grid[newR][newC-1] >= '0' && grid[newR][newC-1] <= '9' {
								newC--
							}

							// Add the coordinates to the map
							cs[newR*1000+newC] = struct{}{}
						}
					}
				}
			}

		}

	}

	ns := []int{} // A slice to store part numbers

	// Iterate through the unique coordinates
	for coord := range cs {
		r, c := coord/1000, coord%1000
		numStr := ""

		// Extract the number
		for c < len(grid[r]) && grid[r][c] >= '0' && grid[r][c] <= '9' {
			numStr += string(grid[r][c])
			c++
		}

		// Parse the number and append it to the slice
		if num, err := strconv.Atoi(numStr); err == nil {
			ns = append(ns, num)
		}
	}

	// Calculate and print the sum of part numbers
	for _, num := range ns {
		sum += num
	}

	fmt.Println("Q1) Sum of part numbers:", sum)

}

// Function to calculate the gear ratios of gears in the engine schematic
func q2() {
	total := 0

	for r, row := range grid {
		for c, ch := range row {
			if ch != '*' {
				continue
			}

			cs := make(map[int]struct{})
			directions := []int{-1, 0, 1}

			for _, dr := range directions {
				for _, dc := range directions {
					cr, cc := r+dr, c+dc
					if cr < 0 || cr >= len(grid) || cc < 0 || cc >= len(grid[cr]) || !isdigit(grid[cr][cc]) {
						continue
					}
					for cc > 0 && isdigit(grid[cr][cc-1]) {
						cc--
					}
					cs[cr*1000+cc] = struct{}{}
				}
			}

			if len(cs) != 2 {
				continue
			}

			ns := []int{}

			for coord := range cs {
				cr, cc := coord/1000, coord%1000
				s := ""
				for cc < len(grid[cr]) && isdigit(grid[cr][cc]) {
					s += string(grid[cr][cc])
					cc++
				}
				num, _ := strconv.Atoi(s)
				ns = append(ns, num)
			}

			total += ns[0] * ns[1]
		}
	}

	fmt.Println("Q2) Sum of gear ratios:", total)
}

func main() {
	q1()
	q2()
}

// ---------------------------------------------------------
// Helper functions
// ---------------------------------------------------------

func isdigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
