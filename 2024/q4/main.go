package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

var fileName = "input.txt"

// var fileName = "example.txt"

var file []byte
var lines [][]byte

func init() {
	start := time.Now()
	defer func() {
		fmt.Println("Init took: ", time.Since(start))
	}()

	// read file into memory
	f, err := os.ReadFile(fileName)
	check(err)
	file = f

	// split file into lines
	lines = bytes.Split(file, []byte("\n"))
}

// Need to find word XMAS in the input file
// The words can be reversed or up and down or diagonal
func partOne() {
	start := time.Now()
	defer func() {
		fmt.Println("Part one took: ", time.Since(start))
	}()

	xmasCount := 0
	// first find XMAS left to right in file buffer
	xmasCount = strings.Count(string(file), "XMAS")
	// fmt.Println("XMAS count: ", xmasCount)

	// find XMAS right to left in file buffer
	xmasCount += strings.Count(string(file), "SAMX")

	// fmt.Println("XMAS count: ", xmasCount)

	// find XMAS up and down in file buffer
	// First transpose the file buffer
	// Then find XMAS in the transposed file buffer
	transposed := transpose(lines)

	// this is equivalent to finding XMAS down
	for _, line := range transposed {
		xmasCount += bytes.Count(line, []byte("XMAS"))
	}

	// fmt.Println("XMAS count: ", xmasCount)

	// this is equivalent to finding XMAS up
	for _, line := range transposed {
		xmasCount += bytes.Count(line, []byte("SAMX"))
	}

	// fmt.Println("XMAS count: ", xmasCount)

	// find XMAS diagonal in file buffer
	xmasCount += countDiagonalOccurrences(lines, "XMAS")

	xmasCount += countDiagonalOccurrences(lines, "SAMX")

	fmt.Println("XMAS count: ", xmasCount)

}
func countDiagonalOccurrences(grid [][]byte, word string) int {
	// Get dimensions of the grid
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := getMaxRowLength(grid) // Get the longest row length
	wordLen := len(word)
	count := 0

	// Helper function to check a word along a diagonal
	checkDiagonal := func(startRow, startCol, dirRow, dirCol int) bool {
		for i := 0; i < wordLen; i++ {
			r := startRow + i*dirRow
			c := startCol + i*dirCol
			// Check bounds
			if r < 0 || r >= rows || c < 0 || c >= cols || r >= len(grid) || c >= len(grid[r]) || grid[r][c] != word[i] {
				return false
			}
		}
		return true
	}

	// Iterate over each cell as a starting point
	for r := 0; r < rows; r++ {
		for c := 0; c < len(grid[r]); c++ {
			// Check diagonals from the current cell
			if checkDiagonal(r, c, 1, 1) { // Top-left to bottom-right
				count++
			}
			if checkDiagonal(r, c, 1, -1) { // Top-right to bottom-left
				count++
			}
		}
	}

	return count
}

func getMaxRowLength(grid [][]byte) int {
	maxLength := 0
	for _, row := range grid {
		if len(row) > maxLength {
			maxLength = len(row)
		}
	}
	return maxLength
}

func transpose(matrix [][]byte) [][]byte {
	if len(matrix) == 0 {
		return matrix
	}

	// create a new matrix
	transposed := make([][]byte, len(matrix[0]))

	for i := range transposed {
		transposed[i] = make([]byte, len(matrix))
	}

	for i, row := range matrix {
		for j, val := range row {
			transposed[j][i] = val
		}
	}

	return transposed

}

// For this part we need to find ocurrences of X-MAS
// ie. MAS in diagonals
// eg:
// M.S
// .A.
// M.S
func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("Part two took: ", time.Since(start))
	}()

	masPairCount := 0

	// Initial assumption is to loop through each col and row
	// and check for the char A, if found then check for M and S
	// in the diagonals

	for i, line := range lines {
		for j, char := range line {

			if char != 'A' {
				continue
			}

			firstDiagonalOkay := false
			secondDiagonalOkay := false

			// Check top-left and bottom-right diagonal
			if i > 0 && j > 0 && i+1 < len(lines) && j+1 < len(lines[i+1]) {
				if lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S' || lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M' {
					firstDiagonalOkay = true
				}
			}

			// Check top-right and bottom-left diagonal
			if i > 0 && j+1 < len(line) && i+1 < len(lines) && j-1 >= 0 && j-1 < len(lines[i+1]) {
				if lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S' || lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M' {
					secondDiagonalOkay = true
				}
			}

			if firstDiagonalOkay && secondDiagonalOkay {
				masPairCount++
			}

		}
	}

	fmt.Println("MAS count: ", masPairCount)

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
