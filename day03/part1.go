package main

import (
	"github.com/avenue9977/advent-of-code-2023/utils"
	"os"
	"strings"
	"unicode"
)

type Part struct {
	number        int
	lineIndex     int
	startingIndex int
	endingIndex   int
	prevCharRune  rune
	nextCharRune  rune
}

func day3Part1() {
	fileScanner, err := utils.GetFileScanner("./day03/input/day03.txt")
	if err != nil {
		os.Exit(1)
	}

	var lines [][]string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineSlice := strings.Split(line, "")
		lines = append(lines, lineSlice)
	}

	sum := 0

	for _, part := range getParts(&lines) {
		sum += getPartNumber(part, lines)
	}

	println(sum)
}

func getParts(lines *[][]string) []Part {
	var parts []Part

	for i, line := range *lines {
		var numberString string

		part := Part{
			lineIndex: i,
		}

		for y, char := range line {
			currentCharRune := rune(char[0])
			prevCharRune, nextCharRune := getAdjacentCharRunes(line, y)

			isPrevCharDigit := unicode.IsDigit(prevCharRune)
			isCurrentCharDigit := unicode.IsDigit(currentCharRune)
			isNextCharDigit := unicode.IsDigit(nextCharRune)

			if !isCurrentCharDigit {
				continue
			}

			if !isPrevCharDigit {
				if isCurrentCharDigit {
					part.startingIndex = y
					part.prevCharRune = prevCharRune
				}
			}

			if isNextCharDigit {
				numberString += char
			} else {
				part.number = utils.ConvertStringToInt(numberString + char)
				part.endingIndex = y
				part.nextCharRune = nextCharRune
				parts = append(parts, part)
				numberString = ""
			}
		}
	}

	return parts
}

func getAdjacentCharRunes(line []string, currentIndex int) (rune, rune) {
	prevChar := ""
	nextChar := ""

	if currentIndex != 0 {
		prevChar = line[currentIndex-1]
	}

	if currentIndex+1 < len(line) {
		nextChar = line[currentIndex+1]
	}

	prevRune := rune(0)
	nextRune := rune(0)

	if prevChar != "" {
		prevRune = rune(prevChar[0])
	}

	if nextChar != "" {
		nextRune = rune(nextChar[0])
	}

	return prevRune, nextRune
}

func getPartNumber(part Part, lines [][]string) int {
	partNumber := 0
	hasLeftSymbol := false
	hasRightSymbol := false
	hasTopSymbol := false
	hasBottomSymbol := false

	if part.prevCharRune != 46 {
		hasLeftSymbol = unicode.IsPunct(part.prevCharRune) || unicode.IsSymbol(part.prevCharRune)
	}

	if part.nextCharRune != 46 {
		hasRightSymbol = unicode.IsPunct(part.nextCharRune) || unicode.IsSymbol(part.nextCharRune)
	}

	if part.lineIndex > 0 {
		hasTopSymbol = hasAdjacentLineSymbol(lines[part.lineIndex-1], part.startingIndex-1, part.endingIndex+1)
	}

	if part.lineIndex < len(lines)-1 {
		hasBottomSymbol = hasAdjacentLineSymbol(lines[part.lineIndex+1], part.startingIndex-1, part.endingIndex+1)
	}

	if hasLeftSymbol || hasRightSymbol || hasTopSymbol || hasBottomSymbol {
		partNumber = part.number
	}

	return partNumber
}

func hasAdjacentLineSymbol(line []string, fromIndex int, toIndex int) bool {
	lineSize := len(line) - 1
	if fromIndex == -1 {
		fromIndex = 0
	}

	if toIndex > lineSize {
		toIndex = lineSize
	}

	hasAdjacentSymbol := false
	for i := fromIndex; i <= toIndex; i++ {
		char := line[i][0]
		if char == 46 {
			hasAdjacentSymbol = false
			continue
		}
		if unicode.IsPunct(rune(char)) || unicode.IsSymbol(rune(char)) {
			hasAdjacentSymbol = true
			break
		}
	}
	return hasAdjacentSymbol
}
