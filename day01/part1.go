package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func day1Part1() {
	file, err := os.Open("./day01/input/day01.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		firstDigitString := getNumberStringFromLine(line)
		secondDigitString := getNumberStringFromLine(reverseString(line))

		sum += convertStringToInt(firstDigitString + secondDigitString)
	}

	fmt.Printf("%d", sum)

	err = file.Close()
	if err != nil {
		return
	}
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

func convertStringToInt(char string) int {
	number, err := strconv.Atoi(char)
	if err != nil {
		return 0
	}
	return number
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
