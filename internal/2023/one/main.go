package one

import (
	"fmt"
	"strconv"
	"unicode"
	"strings"
)

func PartOne(lines []string) int {
	currTotal := 0
	for i := 0 ; i < len(lines) ; i++ {
		digits := []rune{} 
		for _, char := range lines[i] {
			if !unicode.IsDigit(char) {
				continue
			}

			digits = append(digits, char)
		}

		concat := fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1])
		num, _ := strconv.Atoi(concat)
		currTotal += num
	}

	return currTotal
}

func PartTwo(lines []string) int {
	currTotal := 0
	for i := 0 ; i < len(lines) ; i++ {
		digits := []rune{}
		currRune := 0
		for currRune < len(lines[i]) {
			digit, ok := parseDigit(lines[i][currRune:])
			if !ok {
				currRune += 1
				continue
			}

			digits = append(digits, digit)
			currRune += 1
		}

		concat := fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1])
		num, _ := strconv.Atoi(concat)
		currTotal += num
	}

	return currTotal
}

func parseDigit(s string) (rune, bool) {
	switch {
	case unicode.IsDigit(rune(s[0])):
		return rune(s[0]), true
	case strings.HasPrefix(s, "one"):
		return rune("1"[0]), true
	case strings.HasPrefix(s, "two"):
		return rune("2"[0]), true
	case strings.HasPrefix(s, "three"):
		return rune("3"[0]), true
	case strings.HasPrefix(s, "four"):
		return rune("4"[0]), true
	case strings.HasPrefix(s, "five"):
		return rune("5"[0]), true
	case strings.HasPrefix(s, "six"):
		return rune("6"[0]), true
	case strings.HasPrefix(s, "seven"):
		return rune("7"[0]), true
	case strings.HasPrefix(s, "eight"):
		return rune("8"[0]), true
	case strings.HasPrefix(s, "nine"):
		return rune("9"[0]), true
	}

	return 0, false
}