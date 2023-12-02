package main

import (
	"github.com/avenue9977/advent-of-code-2023/utils"
	"os"
	"strings"
	"unicode"
)

var numbersMap = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func day1Part2() {
	fileScanner, err := utils.GetFileScanner("./day01/input/day01.txt")
	if err != nil {
		os.Exit(1)
	}

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		digits := getNumberStringFromLine2(line)

		sum += utils.ConvertStringToInt(digits[0] + digits[len(digits)-1])
	}

	println(sum)
}

func getNumberStringFromLine2(line string) []string {
	var digits []string
	for index, character := range line {
		if unicode.IsDigit(character) {
			digits = append(digits, string(character))
		}

		for _, number := range numbers {
			if strings.HasPrefix(line[index:], number) {
				digits = append(digits, numbersMap[number])
			}
		}
	}

	return digits
}
