package main

import (
	"github.com/avenue9977/advent-of-code-2023/utils"
	"os"
	"strings"
)

func day3Part2() {
	fileScanner, err := utils.GetFileScanner("./day02/input/test.txt")
	if err != nil {
		os.Exit(1)
	}

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, ": ")
		println(split)
	}

	println(sum)
}
