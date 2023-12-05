package internal

import (
	"adventOfCode/internal/adventhelper"
	"fmt"
	"strconv"
	"strings"
)

func Part1(filename string) (*Challenge, *int) {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	Input := Challenge{RawData: derp}
	Input.init()
	Input.findLocations()
	smallestLocation := 0

	for _, val := range Input.Seeds {
		if smallestLocation == 0 {
			smallestLocation = val.Location
			continue
		}
		if val.Location < smallestLocation {
			smallestLocation = val.Location
		}
	}

	return &Input, &smallestLocation
}

func Part2(input *Challenge) *int {
	var totalSum int
	return &totalSum
}

type Challenge struct {
	RawData        *[]string
	Seeds          []*Seed
	ConversionMaps map[string][]*SourceDestinationMap
}

func (c *Challenge) init() {
	seedSplit := strings.Split((*c.RawData)[0], ": ")
	for _, seed := range strings.Split(seedSplit[1], " ") {
		sInt, _ := strconv.Atoi(seed)
		tmp_seed := Seed{Value: sInt}
		c.Seeds = append(c.Seeds, &tmp_seed)
	}
	c.ConversionMaps = map[string][]*SourceDestinationMap{}
	var currentMap string
	for _, line := range *c.RawData {
		if strings.HasPrefix(line, "seeds: ") {
			// skip first line
			continue
		}
		if line == "" {
			// reset map items in preparation for next
			currentMap = ""
			continue
		}
		if strings.HasSuffix(line, ":") {
			// assign new mapName
			currentMap = strings.TrimSuffix(line, " map:")
			continue
		}
		lineSplit := strings.Split(line, " ")

		destinationRange, _ := strconv.Atoi(lineSplit[0])
		sourceRange, _ := strconv.Atoi(lineSplit[1])
		rangeLength, _ := strconv.Atoi(lineSplit[2])
		lineMap := SourceDestinationMap{
			DestinationRange: destinationRange,
			SourceRange:      sourceRange,
			RangeLength:      rangeLength,
		}
		c.ConversionMaps[currentMap] = append(c.ConversionMaps[currentMap], &lineMap)
	}
}

func (c *Challenge) findLocations() {
	for _, seed := range c.Seeds {
		// Find Soil
		found := false
		for _, mapValue := range c.ConversionMaps["seed-to-soil"] {
			sourceMin := mapValue.SourceRange
			sourceMax := sourceMin + mapValue.RangeLength
			if seed.Value >= sourceMin && seed.Value < sourceMax {
				offset := seed.Value - sourceMin
				seed.Soil = mapValue.DestinationRange + offset
				found = true
				break
			}
		}
		if !found {
			seed.Soil = seed.Value
		}

		// Find fertilizer
		found = false
		for _, mapValue := range c.ConversionMaps["soil-to-fertilizer"] {
			sourceMin := mapValue.SourceRange
			sourceMax := sourceMin + mapValue.RangeLength
			if seed.Soil >= sourceMin && seed.Soil < sourceMax {
				offset := seed.Soil - sourceMin
				seed.Fertilizer = mapValue.DestinationRange + offset
				found = true
				break
			}
		}
		if !found {
			seed.Fertilizer = seed.Soil
		}

		// Find water
		found = false
		for _, mapValue := range c.ConversionMaps["fertilizer-to-water"] {
			sourceMin := mapValue.SourceRange
			sourceMax := sourceMin + mapValue.RangeLength
			if seed.Fertilizer >= sourceMin && seed.Fertilizer < sourceMax {
				offset := seed.Fertilizer - sourceMin
				seed.Water = mapValue.DestinationRange + offset
				found = true
				break
			}
		}
		if !found {
			seed.Water = seed.Fertilizer
		}

		// Find light
		found = false
		for _, mapValue := range c.ConversionMaps["water-to-light"] {
			sourceMin := mapValue.SourceRange
			sourceMax := sourceMin + mapValue.RangeLength
			if seed.Water >= sourceMin && seed.Water < sourceMax {
				offset := seed.Water - sourceMin
				seed.Light = mapValue.DestinationRange + offset
				found = true
				break
			}
		}
		if !found {
			seed.Light = seed.Water
		}

		// Find temperature
		found = false
		for _, mapValue := range c.ConversionMaps["light-to-temperature"] {
			sourceMin := mapValue.SourceRange
			sourceMax := sourceMin + mapValue.RangeLength
			if seed.Light >= sourceMin && seed.Light < sourceMax {
				offset := seed.Light - sourceMin
				seed.Temp = mapValue.DestinationRange + offset
				found = true
				break
			}
		}
		if !found {
			seed.Temp = seed.Light
		}

		// Find humidity
		found = false
		for _, mapValue := range c.ConversionMaps["temperature-to-humidity"] {
			sourceMin := mapValue.SourceRange
			sourceMax := sourceMin + mapValue.RangeLength
			if seed.Temp >= sourceMin && seed.Temp < sourceMax {
				offset := seed.Temp - sourceMin
				seed.Humidity = mapValue.DestinationRange + offset
				found = true
				break
			}
		}
		if !found {
			seed.Humidity = seed.Temp
		}

		// Find location
		found = false
		for _, mapValue := range c.ConversionMaps["humidity-to-location"] {
			sourceMin := mapValue.SourceRange
			sourceMax := sourceMin + mapValue.RangeLength
			if seed.Humidity >= sourceMin && seed.Humidity < sourceMax {
				offset := seed.Humidity - sourceMin
				seed.Location = mapValue.DestinationRange + offset
				found = true
				break
			}
		}
		if !found {
			seed.Location = seed.Humidity
		}
	}
}

type Seed struct {
	Value      int
	Soil       int
	Fertilizer int
	Water      int
	Light      int
	Temp       int
	Humidity   int
	Location   int
}

type SourceDestinationMap struct {
	SourceRange      int
	DestinationRange int
	RangeLength      int
}
