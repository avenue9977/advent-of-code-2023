package utils

import "strconv"

func ConvertStringToInt(char string) int {
	number, err := strconv.Atoi(char)
	if err != nil {
		return 0
	}
	return number
}
