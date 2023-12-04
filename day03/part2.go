package main

import (
	"github.com/avenue9977/advent-of-code-2023/utils"
	"os"
	"strings"
)

type Gear struct {
	number        int
	starIndex     int
	starLineIndex int
}

func day3Part2() {
	fileScanner, err := utils.GetFileScanner("./day03/input/day03.txt")
	if err != nil {
		os.Exit(1)
	}

	sum := 0

	var lines [][]string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineSlice := strings.Split(line, "")
		lines = append(lines, lineSlice)
	}

	var gears []Gear

	for _, part := range getParts(&lines) {
		gear := Gear{
			starIndex:     -1,
			starLineIndex: -1,
		}

		if part.prevCharRune == 42 {
			gear.number = part.number
			gear.starLineIndex = part.lineIndex
			gear.starIndex = part.startingIndex - 1
			gears = append(gears, gear)
			continue
		}

		if part.nextCharRune == 42 {
			gear.number = part.number
			gear.starLineIndex = part.lineIndex
			gear.starIndex = part.endingIndex + 1
			gears = append(gears, gear)
			continue
		}

		lineIndex := -1
		starIndex := -1

		if part.lineIndex > 0 {
			lineIndex = part.lineIndex - 1
			starIndex = getStarsFromRow(lines[lineIndex], part.startingIndex-1, part.endingIndex+1)

			if starIndex > -1 {
				addStarToGear(&gear, lineIndex, starIndex, part.number)
				gears = append(gears, gear)
				continue
			}
		}

		if part.lineIndex < len(lines)-1 {
			lineIndex = part.lineIndex + 1
			starIndex = getStarsFromRow(lines[part.lineIndex+1], part.startingIndex-1, part.endingIndex+1)

			if starIndex > -1 {
				addStarToGear(&gear, lineIndex, starIndex, part.number)
				gears = append(gears, gear)
				continue
			}
		}
	}

	for gearIndex, gear := range gears {
		for i := gearIndex + 1; i < len(gears); i++ {
			g := gears[i]
			if g.starLineIndex == gear.starLineIndex && g.starIndex == gear.starIndex {
				sum += gear.number * g.number
			}
		}
	}

	println(sum)
}

func getStarsFromRow(line []string, fromIndex int, toIndex int) (starIndex int) {
	lineSize := len(line) - 1
	if fromIndex == -1 {
		fromIndex = 0
	}

	if toIndex > lineSize {
		toIndex = lineSize
	}

	starIndex = -1

	for i := fromIndex; i <= toIndex; i++ {
		char := line[i][0]

		if char == 42 {
			starIndex = i
			break
		}
	}

	return starIndex
}

func addStarToGear(gear *Gear, lineIndex int, starIndex int, number int) {
	gear.number = number
	gear.starLineIndex = lineIndex
	gear.starIndex = starIndex
}
