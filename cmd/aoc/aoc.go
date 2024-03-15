package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jamesTait-jt/go-aoc/internal/2023/five"
	"github.com/jamesTait-jt/go-aoc/internal/2023/four"
	"github.com/jamesTait-jt/go-aoc/internal/2023/one"
	"github.com/jamesTait-jt/go-aoc/internal/2023/six"
	"github.com/jamesTait-jt/go-aoc/internal/2023/three"
	"github.com/jamesTait-jt/go-aoc/internal/2023/two"
	"github.com/jamesTait-jt/go-aoc/internal/input"
)

func main() {
	year := flag.Int("year", -1, "the year you would like to run")
	day := flag.Int("day", -1, "the day you would like to run")
	flag.Parse()

	funcs := registerFuncs()

	if *year == -1 {
		log.Fatal("must specify a year with -year")
	}

	if *day == -1 {
		runAllDays(*year, funcs[*year], false)
	} else {
		runSingleDay(*year, *day, false, funcs[*year][*day - 1])
	}
}

func registerFuncs() map[int][]func([]string) {
	return map[int][]func([]string){
		2023: {
			one.Run,
			two.Run,
			three.Run,
			four.Run,
			five.Run,
			six.Run,
		},
	}
}

func runAllDays(year int, daysToRun []func([]string), force bool) {
	for dayIdx, runner := range daysToRun {
		dayToRun := dayIdx + 1
		runSingleDay(year, dayToRun, force, runner)
	}
}

func runSingleDay(year int, day int, force bool, runner func([]string)) {
	input.Download(year, day, force)

	input, err := input.Read(year, day)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("~~~ Year=%d Day=%d\n", year, day)
	runner(input)
}