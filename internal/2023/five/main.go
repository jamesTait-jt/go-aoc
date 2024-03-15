package five

import (
	"fmt"
	"math"
	"strings"

	"github.com/jamesTait-jt/go-aoc/internal/parse"
)

type Almanac struct {
	seedToSoil            []Mapping
	soilToFertiliser      []Mapping
	fertiliserToWater     []Mapping
	waterToLight          []Mapping
	lightToTemperature    []Mapping
	temperatureToHumidity []Mapping
	humidityToLocation    []Mapping
}

func (a Almanac) SeedToLocation(seed int) int {
	soil := a.mapValue(a.seedToSoil, seed)
	fertiliser := a.mapValue(a.soilToFertiliser, soil)
	water := a.mapValue(a.fertiliserToWater, fertiliser)
	light := a.mapValue(a.waterToLight, water)
	temp := a.mapValue(a.lightToTemperature, light)
	humidity := a.mapValue(a.temperatureToHumidity, temp)
	location := a.mapValue(a.humidityToLocation, humidity)

	return location
}

func (a Almanac) mapValue(mappings []Mapping, key int) int {
	for _, mapping := range mappings {
		if !mapping.Maps(key) {
			continue
		}

		return mapping.Dest(key)
	}

	return key
}

type Mapping struct {
	destRangeStart int
	srcRangeStart  int
	rangeLen       int
}

func (m Mapping) Maps(n int) bool {
	return m.srcRangeStart <= n && n <= m.srcRangeStart+m.rangeLen
}

func (m Mapping) Dest(n int) int {
	return m.destRangeStart + (n - m.srcRangeStart)
}

func PartOne(lines []string) int {
	seeds := parse.Nums(strings.Split(lines[0], ": ")[1], " ")

	almanac := parseAlmanac(lines[1:])

	lowestLocation := math.MaxUint32
	for _, seed := range seeds {
		location := almanac.SeedToLocation(seed)

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func PartTwo(lines []string) int {
	seedInput := parse.Nums(strings.Split(lines[0], ": ")[1], " ")
	seeds := []int{}
	i := 0
	for i < len(seedInput)-1 {
		initialSeed := seedInput[i]
		numSeeds := seedInput[i+1]
		for j := 0; j < numSeeds; j++ {
			seeds = append(seeds, initialSeed+j)
		}

		i += 2
	}

	fmt.Println(seeds)

	// almanac := parseAlmanac(lines[1:])

	// fmt.Println("=== Parsed Almanac ===")

	// lowestLocation := math.MaxUint32
	// for _, seed := range seeds {
	// 	location := almanac.SeedToLocation(seed)

	// 	if location < lowestLocation {
	// 		lowestLocation = location
	// 	}
	// }

	// return lowestLocation

	return 0 
}

func parseAlmanac(lines []string) Almanac {
	a := Almanac{}

	i := 0
	for i < len(lines) {
		switch lines[i] {
		case "seed-to-soil map:":
			i += 1
			a.seedToSoil = parseMappings(lines, i)
			i += len(a.seedToSoil)

		case "soil-to-fertilizer map:":
			i += 1
			a.soilToFertiliser = parseMappings(lines, i)
			i += len(a.soilToFertiliser)

		case "fertilizer-to-water map:":
			i += 1
			a.fertiliserToWater = parseMappings(lines, i)
			i += len(a.fertiliserToWater)

		case "water-to-light map:":
			i += 1
			a.waterToLight = parseMappings(lines, i)
			i += len(a.waterToLight)

		case "light-to-temperature map:":
			i += 1
			a.lightToTemperature = parseMappings(lines, i)
			i += len(a.lightToTemperature)

		case "temperature-to-humidity map:":
			i += 1
			a.temperatureToHumidity = parseMappings(lines, i)
			i += len(a.temperatureToHumidity)

		case "humidity-to-location map:":
			i += 1
			a.humidityToLocation = parseMappings(lines, i)
			i += len(a.humidityToLocation)

		default:
			i += 1
		}
	}

	return a
}

func parseMappings(lines []string, startIdx int) []Mapping {
	mappings := []Mapping{}

	for i := startIdx; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}

		rangeInfo := parse.Nums(lines[i], " ")
		mappings = append(mappings, Mapping{rangeInfo[0], rangeInfo[1], rangeInfo[2]})
	}

	return mappings
}
