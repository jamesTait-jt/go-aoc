package two

import (
	"strconv"
	"strings"
)

func PartOne(lines []string) int {
	sum := 0
	for idx, game := range lines {
		game = strings.Split(game, ": ")[1]
		if validGame(game) {
			sum += idx + 1
		}
	}

	return sum
}

func PartTwo(lines []string) int {
	sum := 0
	for _, game := range lines {
		game = strings.Split(game, ": ")[1]
		minSet := getMinSet(game)
		power := getPower(minSet)
		sum += power
	}

	return sum
}

func validGame(game string) bool {
	for _, set := range strings.Split(game, "; ") {
		parsedSet := newSet(set)
		if parsedSet["red"] > 12 || parsedSet["green"] > 13 || parsedSet["blue"] > 14 {
			return false
		}
	}

	return true
}

func getMinSet(game string) map[string]int {
	minSet := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range strings.Split(game, "; ") {
		parsedSet := newSet(set)
		for _, colour := range []string{"red", "blue", "green"} {
			if parsedSet[colour] > minSet[colour] {
				minSet[colour] = parsedSet[colour]
			}
		}
	}

	return minSet
}

func getPower(set map[string]int) int {
	return set["red"] * set["blue"] * set["green"]
}

func newSet(str string) map[string]int {
	set := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, subset := range strings.Split(str, ", ") {
		spl := strings.Split(subset, " ")
		count := spl[0]
		colour := spl[1]
		n, _ := strconv.Atoi(count)
		set[colour] += n
	}

	return set
}
