package main

import (
	"github.com/avenue9977/advent-of-code-2023/utils"
	"os"
	"slices"
	"strings"
)

type Game struct {
	playerNumbers  []string
	winningNumbers []string
}

func (game *Game) getPoints() int {
	sum := 0

	for _, number := range game.playerNumbers {
		if slices.Contains(game.winningNumbers, number) {
			if sum == 0 {
				sum++
			} else {
				sum *= 2
			}
		}
	}

	return sum
}

func day4Part1() {
	fileScanner, err := utils.GetFileScanner("./day04/input/day04.txt")
	if err != nil {
		os.Exit(1)
	}

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, ":")
		numbers := strings.Split(split[1], "|")

		game := Game{
			playerNumbers:  getGameNumbers(strings.Split(strings.TrimSpace(numbers[0]), " ")),
			winningNumbers: getGameNumbers(strings.Split(strings.TrimSpace(numbers[1]), " ")),
		}

		sum += game.getPoints()
	}

	println(sum)
}

func getGameNumbers(str []string) []string {
	var numbers []string

	for _, number := range str {
		if number != "" {
			numbers = append(numbers, number)
		}
	}

	return numbers
}
