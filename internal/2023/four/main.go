package four

import (
	"strconv"
	"strings"
)

func PartOne(lines []string) int {
	total := 0
	for _, line := range lines {
		input := strings.Split(line, ": ")[1]
		winningNumbers := getWinningNumbers(input)
		myNumbersStr := strings.Split(input, " | ")[1]
		myNumbers := strings.Split(myNumbersStr, " ")

		points := 0
		for _, myNumber := range myNumbers {
			if myNumber == "" {
				continue
			}

			n, _ := strconv.Atoi(myNumber)
			_, ok := winningNumbers[n]
			if !ok {
				continue
			}
			
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
		total += points
	}

	return total
}

func PartTwo(lines []string) int {
	scratchcards := map[int]int{}
	for idx, line := range lines {
		cardNumber := idx+1

		// Add the initial copy of the scratch card
		if _, ok := scratchcards[cardNumber]; !ok {
			scratchcards[cardNumber] = 1
		}else {
			scratchcards[cardNumber] = scratchcards[cardNumber] + 1
		}

		input := strings.Split(line, ": ")[1]
		winningNumbers := getWinningNumbers(input)
		myNumbersStr := strings.Split(input, " | ")[1]
		myNumbers := strings.Split(myNumbersStr, " ")

		matches := 0
		for _, myNumber := range myNumbers {
			if myNumber == "" {
				continue
			}

			n, _ := strconv.Atoi(myNumber)
			_, ok := winningNumbers[n]
			if !ok {
				continue
			}
			
			matches += 1	
		}

		for i := cardNumber + 1 ; i < cardNumber + 1 + matches ; i++ {
			if _, ok := scratchcards[cardNumber]; !ok {
				scratchcards[cardNumber] = 1
			} else {
				scratchcards[i] = scratchcards[i] + scratchcards[cardNumber]
			}
		}
	}

	sum := 0
	for _, copies := range scratchcards {
		sum += copies
	}

	return sum
}

func getWinningNumbers(card string) map[int]struct{} {
	winningNumbersStr := strings.Split(card, " | ")[0]
	winningNumbersSl := strings.Split(winningNumbersStr, " ")
	winningNumbers := map[int]struct{}{}
	for _, winningNumber := range winningNumbersSl {
		if winningNumber == "" {
			continue
		}

		n, _ := strconv.Atoi(winningNumber)
		winningNumbers[n] = struct{}{}
	}

	return winningNumbers
}

