package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetFileScanner(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, nil
}
