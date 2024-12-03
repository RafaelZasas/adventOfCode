// *NOTE* This solution is not complete. I have not been able to figure out how to sort the hands
// The handType sort works, but the card value sort does not.
// The hands of the same type should be decreasing in strength based on the card values.
// ie. Aces in the first position should be the strongest hand of that type, and 2s in the first position
// should be the weakest hand of that type.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Hand struct {
	cards    []int
	handType int
	bid      int
}

func (h Hand) String() string {
	var cards string
	for _, c := range h.cards {
		// convert back to String
		switch c {
		case 14:
			cards = fmt.Sprintf("%s%s", cards, "A")
		case 13:
			cards = fmt.Sprintf("%s%s", cards, "K")
		case 12:
			cards = fmt.Sprintf("%s%s", cards, "Q")
		case 11:
			cards = fmt.Sprintf("%s%s", cards, "J")
		case 10:
			cards = fmt.Sprintf("%s%s", cards, "T")
		default:
			cards = fmt.Sprintf("%s%d", cards, c)
		}
	}
	return fmt.Sprintf("%v: %d\n", cards, h.bid)
}

var hands []Hand

func init() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		cards := getCards(parts[0])

		handType := getHandType(cards)

		hand := Hand{
			cards:    cards,
			bid:      bid,
			handType: handType,
		}

		hands = append(hands, hand)
	}
}

// in q1 we need to find the sum of all of the calculated
// the bid ammounts.
// The bid ammounts are calculated by multiplying the bid by the rank of the hand.
func q1() {
	start := time.Now()

	sortedHands := sortHands(hands)

	for i := 0; i < len(sortedHands); i++ {
		sortedHands = sortHands(sortedHands)
	}

	bidAmmounts := make([]int, len(sortedHands))
	for i, h := range sortedHands {
		bidAmmounts[i] = h.bid * (i + 1)
	}

	sum := 0
	for _, h := range bidAmmounts {
		sum += h
	}

	fmt.Printf("Q1) The sum of the bid ammounts is: %d\n", sum)
	fmt.Printf("Running time: %s\n\n", time.Since(start))
	fmt.Println(sortedHands)
}

func main() {
	q1()
}

// -------------------------------
// HELPER FUNCTIONS
// -------------------------------

// getCards takes a string of cards and returns a slice of ints
// representing the cards
func getCards(s string) []int {
	cards := make([]int, 0)

	for _, c := range s {
		switch c {
		case 'A':
			cards = append(cards, 14)
		case 'K':
			cards = append(cards, 13)
		case 'Q':
			cards = append(cards, 12)
		case 'J':
			cards = append(cards, 11)
		case 'T':
			cards = append(cards, 10)
		default:
			i, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			cards = append(cards, i)
		}
	}

	return cards
}

// getHandType takes a slice of ints representing cards and returns
// an int representing the type of the hand
// The types are ranked and is ordered as follows (higher is better):
// 1. High card (All labels are distinct)
// 2. Pair (2 cards same)
// 3. Two pair (2 cards same, 2 cards same)
// 4. Three of a kind (3 cards same)
// 5. Full House (3 cards same)
// 6. Four of a kind
// 7. 5 of a kind
func getHandType(cards []int) int {
	cardCounts := make(map[int]int)

	for _, c := range cards {
		cardCounts[c]++
	}

	// 5 of a kind
	if len(cardCounts) == 1 {
		return 7
	}

	// 4 of a kind
	if len(cardCounts) == 2 {
		for _, v := range cardCounts {
			if v == 4 {
				return 6
			}
		}
	}

	// Full house
	if len(cardCounts) == 2 {
		for _, v := range cardCounts {
			if v == 3 {
				return 5
			}
		}
	}

	// 3 of a kind
	if len(cardCounts) == 3 {
		for _, v := range cardCounts {
			if v == 3 {
				return 4
			}
		}
	}

	// Two pair
	if len(cardCounts) == 3 {
		return 3
	}

	// Pair
	if len(cardCounts) == 4 {
		return 2
	}

	// High cards
	return 1
}

// sortHands takes a slice of hands and returns a slice of hands
// sorted by the rules of the game
// The rules are as follows:
// The hand in the first position is the weakest hand and the hand in the last position
// is the strongest hand.
//
// First sorting priority is the hand type.
// higher hand type is better ie. handType 5 (full house) > handType 4 (four of a kind).
//
// Second sorting priority is calculated by looking at the values or labels of the
// cards in the same posiiton. ie. the higher card in the first position of the hand,
// if the first card is the same then the higher card in the second position of the hand,
// if the second card is the same then the higher card in the third position of the hand etc.
//
// At the end, the sorted hands should have five of a kind hands at the end of the array
// and high card hands at the beginning of the array.
//
// Hands of the same type should be decreasing in strength based on the card values.
// ie. Aces in the first position should be the strongest hand of that type, and 2s in the first position
// should be the weakest hand of that type.
// This is not working as intended.
func sortHands(hands []Hand) []Hand {
	sortedHands := make([]Hand, len(hands))

	copy(sortedHands, hands)

	// sort by hand type
	for i := 0; i < len(sortedHands); i++ {
		for j := i + 1; j < len(sortedHands); j++ {
			if sortedHands[i].handType > sortedHands[j].handType {
				sortedHands[i], sortedHands[j] = sortedHands[j], sortedHands[i]
			}
		}
	}

	// sort by card values
	for i := 0; i < len(sortedHands); i++ {
		for j := i + 1; j < len(sortedHands); j++ {
			for k := 0; k < len(sortedHands[i].cards); k++ {
				if sortedHands[i].cards[k] > sortedHands[j].cards[k] {
					sortedHands[i], sortedHands[j] = sortedHands[j], sortedHands[i]
				}
			}
		}
	}

	return sortedHands
}
