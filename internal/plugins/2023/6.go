package main

import (
	"strconv"
	"strings"

	"github.com/jamesTait-jt/go-aoc/internal/parse"
)

var Day6 day6

type day6 struct{}

func (d *day6) PartOne(lines []string) string {
	times := parse.Nums(strings.Split(lines[0], ":")[1], " ")
	dists := parse.Nums(strings.Split(lines[1], ":")[1], " ")

	mult := 1
	for raceIdx := 0; raceIdx < len(times); raceIdx++ {
		waysToWin := 0
		for secondsHeld := 0; secondsHeld <= times[raceIdx]; secondsHeld++ {
			distanceTravelled := (times[raceIdx] - secondsHeld) * secondsHeld
			if distanceTravelled > dists[raceIdx] {
				waysToWin += 1
			}
		}
		mult *= waysToWin
	}

	return strconv.Itoa(mult)
}

func (d *day6) PartTwo(lines []string) string {
	timeWithWhiteSpace := strings.Split(lines[0], ":")[1]
	timeStr := strings.Join(strings.Fields(timeWithWhiteSpace), "")
	time, _ := strconv.Atoi(timeStr)

	distWithWhiteSpace := strings.Split(lines[1], ":")[1]
	distStr := strings.Join(strings.Fields(distWithWhiteSpace), "")
	dist, _ := strconv.Atoi(distStr)

	waysToWin := 0
	for secondsHeld := 0; secondsHeld <= time; secondsHeld++ {
		distanceTravelled := (time - secondsHeld) * secondsHeld
		if distanceTravelled > dist {
			waysToWin += 1
		}
	}

	return strconv.Itoa(waysToWin)
}
