package main

import (
	"github.com/avenue9977/advent-of-code-2023/utils"
	"os"
	"strings"
)

func day4Part2() {
	fileScanner, err := utils.GetFileScanner("./day04/input/day04.txt")
	if err != nil {
		os.Exit(1)
	}

	var games []Game

	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, ":")
		game := getGame(strings.Split(split[0], " ")[1], split)
		games = append(games, *game)
	}

	var count = new(int)
	matchCount := getWinningGamesC(games, count)

	println(matchCount)
}

func getWinningGamesC(games []Game, count *int) int {
	if len(games) == 0 {
		return 0
	}

	for i, game := range games {
		matches := game.getMatches()
		getWinningGamesC(games[i+1:i+matches+1], count)
		*count++
	}

	return *count
}
