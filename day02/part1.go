package main

import (
	"github.com/avenue9977/advent-of-code-2023/utils"
	"os"
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
	fileScanner, err := utils.GetFileScanner("./day02/input/day02.txt")
	if err != nil {
		os.Exit(1)
	}

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
}

func getGameId(str string) int {
	split := strings.Fields(str)
	return utils.ConvertStringToInt(split[1])
}

func getMiniGames(str string) []MiniGame {
	games := getGames(str)
	var miniGames []MiniGame

	for _, game := range games {
		gameColors := getGameColors(game)
		miniGame := MiniGame{}

		for _, colorString := range gameColors {
			colorPair := getColorCountPair(colorString)
			colorCount := utils.ConvertStringToInt(colorPair[0])
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

		miniGames = append(miniGames, miniGame)
	}

	return miniGames
}

func getGames(str string) []string {
	return strings.Split(str, "; ")
}

func getGameColors(str string) []string {
	return strings.Split(str, ", ")
}

func getColorCountPair(colorString string) []string {
	return strings.Split(colorString, " ")
}
