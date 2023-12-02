package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

type Game struct {
	id    int
	games []MiniGame
}

func (g *Game) isValid() bool {
	isValid := true

	for _, game := range g.games {
		if !game.isValid() {
			isValid = false
			break
		}
	}

	return isValid
}

type MiniGame struct {
	red   int
	green int
	blue  int
}

func (g *MiniGame) isValid() bool {
	return g.red <= MAX_RED && g.green <= MAX_GREEN && g.blue <= MAX_BLUE
}

func day2Part1() {
	file, err := os.Open("./day02/input/day02.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, ":")
		games := getMiniGames(strings.Trim(split[1], " "))

		game := Game{
			id:    getGameId(split[0]),
			games: games,
		}

		if game.isValid() {
			sum += game.id
		}
	}

	println(sum)

	err = file.Close()
	if err != nil {
		return
	}
}

func getGameId(str string) int {
	split := strings.Fields(str)
	num, err := strconv.Atoi(split[1])

	if err != nil {
		panic("Failed to convert gameId")
	}

	return num
}

func getMiniGames(str string) []MiniGame {
	var games []MiniGame
	innerGames := strings.Split(str, "; ")

	for _, game := range innerGames {
		gameColors := strings.Split(game, ", ")
		miniGame := MiniGame{}

		for _, colorString := range gameColors {
			colorPair := strings.Split(colorString, " ")
			colorCount, _ := strconv.Atoi(colorPair[0])
			color := colorPair[1]

			switch color {
			case "red":
				miniGame.red = colorCount
				break
			case "green":
				miniGame.green = colorCount
				break
			case "blue":
				miniGame.blue = colorCount
				break
			}
		}

		games = append(games, miniGame)
	}

	return games
}
