package five

import (
	"math"
	"strings"

	"github.com/jamesTait-jt/go-aoc/internal/parse"
)

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
	seedRanges := [][]int{}
	i := 0
	for i < len(seedInput)-1 {
		seedRanges = append(seedRanges, []int{seedInput[i], seedInput[i+1]})

		i += 2
	}

	almanac := parseAlmanac(lines[1:])

	for i := 0; i < math.MaxInt32; i++ {
		seed := almanac.LocationToSeed(i)
		for _, seedRange := range seedRanges {
			if seedRange[0] <= seed && seed < seedRange[0]+seedRange[1] {
				return i
			}
		}
	}

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

func (a Almanac) LocationToSeed(location int) int {
	humidity := a.reverseMapValue(a.humidityToLocation, location)
	temp := a.reverseMapValue(a.temperatureToHumidity, humidity)
	light := a.reverseMapValue(a.lightToTemperature, temp)
	water := a.reverseMapValue(a.waterToLight, light)
	fertiliser := a.reverseMapValue(a.fertiliserToWater, water)
	soil := a.reverseMapValue(a.soilToFertiliser, fertiliser)
	seed := a.reverseMapValue(a.seedToSoil, soil)

	return seed
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

func (a Almanac) reverseMapValue(mappings []Mapping, dest int) int {
	for _, mapping := range mappings {
		if !mapping.MapsTo(dest) {
			continue
		}

		return mapping.Src(dest)
	}

	return dest
}

type Mapping struct {
	destRangeStart int
	srcRangeStart  int
	rangeLen       int
}

func (m Mapping) Maps(n int) bool {
	return m.srcRangeStart <= n && n < m.srcRangeStart+m.rangeLen
}

func (m Mapping) MapsTo(n int) bool {
	return m.destRangeStart <= n && n < m.destRangeStart+m.rangeLen
}

func (m Mapping) Dest(n int) int {
	return m.destRangeStart + (n - m.srcRangeStart)
}

func (m Mapping) Src(n int) int {
	return m.srcRangeStart + (n - m.destRangeStart)
}
