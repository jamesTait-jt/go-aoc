package three

import (
	"fmt"
	"strconv"
	"unicode"
)

func Run(lines []string) {
	fmt.Println("Part 1: ", partOne(lines))
	fmt.Println("Part 2: ", partTwo(lines))
}

func partOne(lines []string) int {
	sum := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if !unicode.IsDigit(rune(lines[i][j])) {
				continue
			}

			num := parseNum(lines[i][j:])

			leBound := j - 1
			riBound := j + len(num)
			upBound := i - 1
			loBound := i + 1

			if adjacentSymbol(lines, leBound, riBound, upBound, loBound) {
				n, _ := strconv.Atoi(num)
				sum += n
			}

			j += len(num)
		}
	}

	return sum
}

func partTwo(lines []string) int {
	starsAndAdjacentNumbers := map[coord][]string{}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if !unicode.IsDigit(rune(lines[i][j])) {
				continue
			}

			num := parseNum(lines[i][j:])

			leBound := j - 1
			riBound := j + len(num)
			upBound := i - 1
			loBound := i + 1

			adjacentStars := getAdjacentStars(lines, leBound, riBound, upBound, loBound)

			for _, star := range adjacentStars {
				starsAndAdjacentNumbers[star] = append(starsAndAdjacentNumbers[star], num)
			}	

			j += len(num)
		}
	}

	sum := 0
	for _, adjacentNumbers := range starsAndAdjacentNumbers {
		if len(adjacentNumbers) == 2 {
			x, _ := strconv.Atoi(adjacentNumbers[0])
			y, _ := strconv.Atoi(adjacentNumbers[1])

			sum += x * y
		}
	}

	return sum
}

func parseNum(s string) string {
	soFar := ""
	i := 0
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		soFar += string(s[i])
		i++
	}

	return soFar
}

func adjacentSymbol(lines []string, leBound, riBound, upBound, loBound int) bool {
	for i := upBound; i <= loBound; i++ {
		for j := leBound; j <= riBound; j++ {
			if i < 0 || j < 0 || i >= len(lines) || j >= len(lines[0]) {
				continue
			}

			if unicode.IsDigit(rune(lines[i][j])) {
				continue
			}

			if lines[i][j] == []byte(".")[0] {
				continue
			}

			return true
		}
	}

	return false
}

type coord struct {
	i int
	j int
}

func getAdjacentStars(lines []string, leBound, riBound, upBound, loBound int) []coord {
	stars := []coord{}
	for i := upBound; i <= loBound; i++ {
		for j := leBound; j <= riBound; j++ {
			if i < 0 || j < 0 || i >= len(lines) || j >= len(lines[0]) {
				continue
			}

			if lines[i][j] == []byte("*")[0] {
				stars = append(stars, coord{i, j})
			}
		}
	}

	return stars
}