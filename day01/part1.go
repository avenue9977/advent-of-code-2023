package main

import (
	"github.com/avenue9977/advent-of-code-2023/utils"
	"os"
	"unicode"
)

func day1Part1() {
	fileScanner, err := utils.GetFileScanner("./day01/input/day01.txt")
	if err != nil {
		os.Exit(1)
	}

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		firstDigitString := getNumberStringFromLine(line)
		secondDigitString := getNumberStringFromLine(utils.ReverseString(line))

		sum += int(10*(firstDigitString[0]-'0') + (secondDigitString[0] - '0'))
	}

	println(sum)
}

func getNumberStringFromLine(line string) string {
	var digit string
	for _, character := range line {
		if unicode.IsDigit(character) {
			digit = string(character)
			break
		}
	}

	return digit
}
