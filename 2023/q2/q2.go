package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Event struct {
	green int
	blue  int
	red   int
}

func (e Event) String() string {
	return fmt.Sprintf("green: %d, blue: %d, red: %d\n", e.green, e.blue, e.red)
}

type game struct {
	id     int
	events []Event
}

func (g game) String() string {
	return fmt.Sprintf("Game %d\n%v\n", g.id, g.events)
}

var games []game

func init() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	gameIdx := 1

	for scanner.Scan() {

		var g game

		g.id = gameIdx

		// split line by colon
		line := strings.Split(scanner.Text(), ":")[1]

		// split events by semicolon, ie. 10 green , 5 blue; 3 red;3 green,2 blue
		gameEvents := strings.Split(line, ";")

		// balls are split by comma, ie. 10 blue, 5 red, 3 green
		for _, ge := range gameEvents {

			var event Event
			balls := strings.Split(ge, ",")

			// ball is split by space, ie. blue 10
			for _, b := range balls {
				res := strings.Split(b, " ")

				// add ball to game
				switch res[2] {
				case "green":
					num, err := strconv.Atoi(res[1])
					if err != nil {
						log.Fatal(err)
					}
					event.green = num

				case "blue":
					num, err := strconv.Atoi(res[1])
					if err != nil {
						log.Fatal(err)
					}
					event.blue = num

				case "red":
					num, err := strconv.Atoi(res[1])
					if err != nil {
						log.Fatal(err)
					}
					event.red = num
				}
			}

			g.events = append(g.events, event)
		}

		games = append(games, g)

		gameIdx++
	}

}

func q1() {

	sum := 0

	for _, g := range games {

		gameValid := true
		// in any of the games events, there may be no more than 12 red, 13 green or 14 blue balls
		for _, e := range g.events {
			if e.red > 12 || e.green > 13 || e.blue > 14 {
				gameValid = false
				continue
			}
		}

		if gameValid {
			sum += g.id
		}

	}

	fmt.Println(sum)
}

// find the minimum number of balls that can be used to play all the games
// calculate the power of each game by multiplying the minimumm number of balls of each color for the game
// the total power is the sum of the power of all games
func q2() {

	sum := 0
	for _, g := range games {

		// set them temporarily to the first game's first event
		minGreen := g.events[0].green
		minBlue := g.events[0].blue
		minRed := g.events[0].red

		for _, e := range g.events {
			if e.green > minGreen {
				minGreen = e.green
			}
			if e.blue > minBlue {
				minBlue = e.blue
			}
			if e.red > minRed {
				minRed = e.red
			}
		}

		//fmt.Printf("Game %d: minGreen: %d, minBlue: %d, minRed: %d\n", g.id, minGreen, minBlue, minRed)

		gamePower := minGreen * minBlue * minRed
		sum += gamePower
	}

	fmt.Printf("Total power: %d\n", sum)

}

func main() {
	q1()
	q2()
}
