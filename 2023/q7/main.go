package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

// var fileName = "ex.txt"

var fileName = "in.txt"

type Hand struct {
	bid      int16
	hand     string
	handRank int8
}

func (h Hand) String() string {
	return fmt.Sprintf("Hand{bid: %d, hand: %s, handRank: %d}\n", h.bid, h.hand, h.handRank)
}

var hands []Hand

func init() {
	start := time.Now()
	defer func() {
		fmt.Println("init took", time.Since(start))
	}()

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()

		parts := bytes.Split(line, []byte(" "))

		hand := parts[0]
		bid, err := strconv.Atoi(string(parts[1]))
		check(err)

		handStrength := getHandStrength(hand)

		hands = append(hands, Hand{int16(bid), string(hand), handStrength})

	}
	// fmt.Println(hands)
}

func partOne() {
	start := time.Now()
	defer func() {
		fmt.Println("partOne took", time.Since(start))
	}()

	// sort hands by handRank
	// if handRank is the same, sort each card in the hand and compare
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handRank == hands[j].handRank {
			for k := 0; k < 5; k++ {
				card1 := getCardValue(hands[i].hand[k])
				card2 := getCardValue(hands[j].hand[k])

				if card1 == card2 {
					continue
				}

				if card1 < card2 {
					return true
				} else {
					return false
				}

			}

		}
		return hands[i].handRank < hands[j].handRank
	})

	// sum up winnings by multiplying bid by position in sorted list
	winnings := 0
	for i, h := range hands {
		fmt.Println(i, h)

		winnings += int(h.bid) * (i + 1)
	}
	fmt.Println("Winnings:", winnings)

}

func partTwo() {
	start := time.Now()
	defer func() {
		println("partTwo took", time.Since(start))
	}()
}

func main() {
	partOne()
	partTwo()
}

func getCardValue(c byte) int8 {
	if c == 'A' {
		return 14
	}
	if c == 'K' {
		return 13
	}
	if c == 'Q' {
		return 12
	}
	if c == 'J' {
		return 11
	}
	if c == 'T' {
		return 10
	}
	return int8(c - '0')
}

func getHandStrength(hand []byte) int8 {
	// check for 5 of a kind
	if hand[0] == hand[1] && hand[1] == hand[2] && hand[2] == hand[3] && hand[3] == hand[4] {
		return 7
	}

	// check for 4 of a kind
	for _, c := range hand {
		if bytes.Count(hand, []byte{c}) == 4 {
			return 6
		}
	}

	// check for full house
	for _, c := range hand {
		if bytes.Count(hand, []byte{c}) == 3 {
			for _, c2 := range hand {
				if bytes.Count(hand, []byte{c2}) == 2 {
					return 5
				}
			}
		}
	}

	// check for three of a kind
	for _, c := range hand {
		if bytes.Count(hand, []byte{c}) == 3 {
			return 4
		}
	}

	// check for two pair
	for _, c := range hand {
		if bytes.Count(hand, []byte{c}) == 2 {
			for _, c2 := range hand {
				if bytes.Count(hand, []byte{c2}) == 2 && c != c2 {
					return 3
				}
			}
		}
	}

	// check for one pair
	for _, c := range hand {
		if bytes.Count(hand, []byte{c}) == 2 {
			return 2
		}
	}

	return 1

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
