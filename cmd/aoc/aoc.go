package main

import (
	"fmt"
	"log"

	"github.com/jamesTait-jt/go-aoc/internal/2023/five"
	"github.com/jamesTait-jt/go-aoc/internal/2023/four"
	"github.com/jamesTait-jt/go-aoc/internal/2023/one"
	"github.com/jamesTait-jt/go-aoc/internal/2023/six"
	"github.com/jamesTait-jt/go-aoc/internal/2023/three"
	"github.com/jamesTait-jt/go-aoc/internal/2023/two"
	"github.com/jamesTait-jt/go-aoc/internal/config"
	"github.com/jamesTait-jt/go-aoc/internal/input"
)

func main() {
	appConfig, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	runners := registerRunners()

	if err = input.Download(appConfig); err != nil {
		log.Fatal(err)
	}

	if err = run(appConfig, runners); err != nil {
		log.Fatal(err)
	}
}

func registerRunners() map[int]map[int]func([]string) {
	return map[int]map[int]func([]string){
		2023: {
			1: one.Run,
			2: two.Run,
			3: three.Run,
			4: four.Run,
			5: five.Run,
			6: six.Run,
		},
	}
}

func run(appConfig config.AppConfig, runners map[int]map[int]func([]string)) error {
	for _, dayToRun := range appConfig.Days {
		yearRunners, ok := runners[appConfig.Year]
		if !ok {
			return fmt.Errorf("no runners found for year=%d - please ensure you have registered the runner with the Run() function for the given day", appConfig.Year)
		}

		runner, ok := yearRunners[dayToRun]
		if !ok {
			return fmt.Errorf("no runner found for year=%d day=%d - please ensure you have registered the runner with the Run() function for the given day", appConfig.Year, dayToRun)
		}

		input, err := input.Read(appConfig.Year, dayToRun, appConfig.Input)
		if err != nil {
			return err
		}

		fmt.Printf("~~~ year=%d day=%d\n", appConfig.Year, dayToRun)
		runner(input)
	}

	return nil
}
