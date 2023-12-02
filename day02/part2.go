package main

import (
	"github.com/avenue9977/advent-of-code-2023/utils"
	"os"
	"strings"
)

type ColorGame struct {
	id    int
	red   int
	green int
	blue  int
}

func (g *ColorGame) getPowerOfColors() int {
	return g.red * g.blue * g.green
}

func day2Part2() {
	fileScanner, err := utils.GetFileScanner("./day02/input/day02.txt")
	if err != nil {
		os.Exit(1)
	}

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, ": ")
		colors := getGamesColors(split[1])

		game := ColorGame{
			id:    getGameId(split[0]),
			red:   colors["red"],
			green: colors["green"],
			blue:  colors["blue"],
		}

		sum += game.getPowerOfColors()
	}

	println(sum)
}

func getGamesColors(str string) map[string]int {
	games := getGames(str)
	gamesMap := map[string]int{}

	for _, game := range games {
		gameColors := getGameColors(game)

		for _, colorString := range gameColors {
			colorPair := getColorCountPair(colorString)
			colorCount := utils.ConvertStringToInt(colorPair[0])
			color := colorPair[1]

			value, hasValue := gamesMap[color]

			if hasValue {
				if value < colorCount {
					gamesMap[color] = colorCount
				}
			} else {
				gamesMap[color] = colorCount
			}
		}
	}

	return gamesMap
}
