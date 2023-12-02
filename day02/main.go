package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("AoC 2023 Day 2!\n")

	// open the input file ("input.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// close the file when we're done
	defer file.Close()

	games := ReadGames(file)

	// part1
	part1(games)

	// part2
	part2(games)

}

type Game struct {
	Id       int
	MaxRed   int
	MaxGreen int
	MaxBlue  int
}

func (g Game) IsValidFrom(max Game) bool {
	if g.MaxRed <= max.MaxRed && g.MaxGreen <= max.MaxGreen && g.MaxBlue <= max.MaxBlue {
		return true
	} else {
		return false
	}
}

func (g Game) GetPower() int {
	return g.MaxRed * g.MaxGreen * g.MaxBlue
}

func ReadGames(file *os.File) []Game {
	// create a buffer so we can read line by line
	scanner := bufio.NewScanner(file)

	var games []Game

	// loop through each line
	for scanner.Scan() {
		line := scanner.Text()

		// split line on ':'
		split := strings.Split(line, ":")

		gameAndId := split[0]
		rounds := split[1]

		// split game and id on ' '
		split = strings.Split(gameAndId, " ")
		id, err := strconv.Atoi(split[1])

		if err != nil {
			log.Printf("Error converting id to int: %v\n", split[1])
			log.Fatal(err)
		}

		game, err := findMaxForGame(id, rounds)

		if err != nil {
			log.Fatal(err)
		}

		games = append(games, game)
	}

	return games
}

func part1(games []Game) {
	fmt.Printf("Part 1\n")

	// set target
	maxGame := Game{Id: 999, MaxRed: 12, MaxGreen: 13, MaxBlue: 14}

	idSum := 0

	// check which games are valid
	for _, game := range games {
		if game.IsValidFrom(maxGame) {
			idSum += game.Id
		}
	}

	fmt.Printf("Sum of Ids = %v\n", idSum)
}

func part2(games []Game) {
	fmt.Printf("Part 2\n")

	// compute power for each game
	// and sum
	sum := 0

	for _, game := range games {
		sum += game.GetPower()
	}

	fmt.Printf("Sum of Powers = %v\n", sum)
}

func findMaxForGame(id int, game string) (Game, error) {
	// split rounds on ';'
	rounds := strings.Split(game, ";")

	// make a game
	gameObj := Game{Id: id}

	// loop through each round
	for _, round := range rounds {
		// split rounds on ','
		cubes := strings.Split(round, ",")
		// loop through each cube
		for _, cube := range cubes {
			// trim and split on ' ', [0] is number, [1] is color
			cubeSplit := strings.Split(strings.TrimSpace(cube), " ")

			// convert cubeSplit[0] to int
			// if err, return error
			val, err := strconv.Atoi(cubeSplit[0])

			if err != nil {
				log.Printf("Error converting value %v to int %e", cubeSplit[0], err)
				return Game{}, err
			}

			switch cubeSplit[1] {
			case "red":
				if val > gameObj.MaxRed {
					gameObj.MaxRed = val
				}
			case "green":
				if val > gameObj.MaxGreen {
					gameObj.MaxGreen = val
				}
			case "blue":
				if val > gameObj.MaxBlue {
					gameObj.MaxBlue = val
				}
			}
		}
	}

	return gameObj, nil
}
