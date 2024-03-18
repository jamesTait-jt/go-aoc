package main

import (
	"fmt"
	"log"
	"plugin"

	"github.com/jamesTait-jt/go-aoc/internal/config"
	"github.com/jamesTait-jt/go-aoc/internal/input"
)

const pluginDir = "./internal/plugins"

// aocDay represents a single day of AOC solutions. You must implement this interface for each day as a plugin.
type aocDay interface {
	PartOne([]string) string
	PartTwo([]string) string
}

func main() {
	appConfig, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	runners, err := registerDays(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err = input.Download(appConfig); err != nil {
		log.Fatal(err)
	}

	if err = run(appConfig, runners); err != nil {
		log.Fatal(err)
	}
}

func registerDays(appConfig config.AppConfig) (map[int]map[int]aocDay, error) {
	days := map[int]map[int]aocDay{appConfig.Year: {}}

	for _, day := range appConfig.Days {
		pluginPath := fmt.Sprintf("%s/%d/%d.so", pluginDir, appConfig.Year, day)
		p, err := plugin.Open(pluginPath)
		if err != nil {
			return nil, err
		}

		symbolName := fmt.Sprintf("Day%d", day)
		sym, err := p.Lookup(symbolName)
		if err != nil {
			return nil, err
		}

		runnableDay, ok := sym.(aocDay)
		if !ok {
			return nil, fmt.Errorf("aocDay not implemented correctly for year=%d day=%d", appConfig.Year, day)
		}

		days[appConfig.Year][day] = runnableDay
	}

	return days, nil
}

func run(appConfig config.AppConfig, runnableDays map[int]map[int]aocDay) error {
	for _, dayToRun := range appConfig.Days {
		input, err := input.Read(appConfig.Year, dayToRun, appConfig.Input)
		if err != nil {
			return err
		}

		fmt.Printf("~~~ year=%d day=%d\n", appConfig.Year, dayToRun)

		runner := runnableDays[appConfig.Year][dayToRun]
		fmt.Println("Part 1: ", runner.PartOne(input))
		fmt.Println("Part 2: ", runner.PartTwo(input))
	}

	return nil
}
