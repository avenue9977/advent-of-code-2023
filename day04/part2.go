package main

import (
	"github.com/avenue9977/advent-of-code-2023/utils"
	"os"
	"strings"
)

func day4Part2() {
	fileScanner, err := utils.GetFileScanner("./day03/input/test.txt")
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

	println(sum)
}
