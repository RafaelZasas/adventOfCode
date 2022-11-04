package q4

type board [][]int

var boards []board
var numbers []int

// init formats the input text file into
// an array of bingo numbers
// and an array of 5X5 bingo boards
func init() {

}

// part1 finds the board with a winning row or column
// by marking a matching number and checking for a winning row or column,
// for each input number
func part1() int {

	for _, n := range numbers {
		for _, b := range boards {
			markBoard(&b, n)
			isWinningRow := checkWinningRow(b)
			isWinningCol := checkWinningCol(b)

			if isWinningCol || isWinningRow {
				return calcWinningScore(b, n)
			}
		}
	}
	return -1
}

// calcWinningScore sums all the unmarked numbers
// and returns the product of that and the last called number
func calcWinningScore(b board, n int) int {
	sum := 0

	for i, _ := range b {
		for j, _ := range b[i] {
			if b[i][j] != -1 {
				sum += b[i][j]
			}
		}
	}
	return sum * n
}

// markBoard marks a given board with an x
// at the position of the given number
// if the number is present on the board
func markBoard(b *board, n int) {
	for i, _ := range *b {
		for j, _ := range (*b)[i] {
			if (*b)[i][j] == n {
				(*b)[i][j] = -1
			}
		}
	}
}

// checkWinningRow checks a given board
// to see if there is a row with 5 X's
func checkWinningRow(b board) bool {

}

// checkWinningCol checks a given board
// to see if there is a column with 5 X's
func checkWinningCol(b board) bool {

}

func main() {

}
