package internal

import (
	"adventOfCode/internal/adventhelper"
	"fmt"
	"sync"
)

type Route int

const (
	North Route = iota
	West
	South
	East
)

func Part1(filename string) (*Challenge, *int) {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	Input := Challenge{RawData: derp}
	Input.init()
	Input.navigateMap()
	derpInt := Input.TravelValues[0]
	for _, num := range Input.TravelValues {
		if num > derpInt {
			derpInt = num
		}
	}
	if derpInt%2 == 0 {
		derpInt = derpInt / 2
	} else {
		derpInt = derpInt/2 + 1
	}

	return &Input, &derpInt
}

func Part2(Input *Challenge) *int {
	lowest := -1
	return &lowest
}

type Challenge struct {
	RawData      *[]string
	Data         [][]*Tile
	SPosition    *Tile
	TravelValues []int
	ValMap       map[string][]Route
}

func (c *Challenge) init() {
	c.ValMap = map[string][]Route{
		"|": []Route{North, South},
		"-": []Route{East, West},
		"L": []Route{North, East},
		"J": []Route{North, West},
		"7": []Route{South, West},
		"F": []Route{South, East},
	}
	c.Data = [][]*Tile{}
	for lineIdx, line := range *c.RawData {
		lineSlice := []*Tile{}
		var tmpTile *Tile
		for charIdx, char := range line {
			charStr := string(char)
			tmpTile = &Tile{Value: charStr}
			if charStr == "S" {
				c.SPosition = tmpTile
			}
			if lineIdx != 0 { // North  & South Path
				// Ignored on the first line in the grid
				// Set Current Tile North Value
				tmpTile.North = c.Data[lineIdx-1][charIdx]
				// Set Previous Line same tile Index South value
				preprevTile := c.Data[lineIdx-1][charIdx]
				preprevTile.South = tmpTile
			}
			if charIdx != 0 { // West & East Path
				// Ignores the first element in the line
				// Set current Tile west Value
				tmpTile.West = lineSlice[charIdx-1]
				// Set Previous Tile East Value
				lineSlice[charIdx-1].East = tmpTile
			}
			tmpTile.defineRoute1and2(c.ValMap)

			lineSlice = append(lineSlice, tmpTile)
		}
		c.Data = append(c.Data, lineSlice)
	}

}

func (c *Challenge) navigateMap() {
	var wg, wgResults sync.WaitGroup
	results := make(chan int)
	if c.SPosition.North.Value == "|" ||
		c.SPosition.North.Value == "7" ||
		c.SPosition.North.Value == "F" {
		// Start the north tile
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := c.SPosition.North.findRouteLength(South)
			results <- res
		}()
	}
	if c.SPosition.West.Value == "-" ||
		c.SPosition.West.Value == "L" ||
		c.SPosition.West.Value == "F" {
		// Start the west tile
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := c.SPosition.West.findRouteLength(East)
			results <- res
		}()
	}
	if c.SPosition.South.Value == "|" ||
		c.SPosition.South.Value == "L" ||
		c.SPosition.South.Value == "J" {
		// Start the south tile
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := c.SPosition.South.findRouteLength(North)
			results <- res
		}()
	}
	if c.SPosition.East.Value == "-" ||
		c.SPosition.East.Value == "J" ||
		c.SPosition.East.Value == "7" {
		// Start the east tile
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := c.SPosition.East.findRouteLength(West)
			results <- res
		}()
	}
	wgResults.Add(1)
	go func() {
		defer wgResults.Done()
		for item := range results {
			c.TravelValues = append(c.TravelValues, item)
		}
	}()
	wg.Wait()
	close(results)
	wgResults.Wait()
}

type Tile struct {
	North  *Tile
	South  *Tile
	East   *Tile
	West   *Tile
	Value  string
	Route1 *Route
	Route2 *Route
}

func (t *Tile) defineRoute1and2(valMap map[string][]Route) {
	if t.Value == "." || t.Value == "S" {
		return
	}
	t.Route1 = &valMap[t.Value][0]
	t.Route2 = &valMap[t.Value][1]
}

func (t *Tile) findRouteLength(comingFrom Route, Start ...bool) int {
	if t.Value == "S" { // return if back at start
		return 1
	}
	var nextPath Route
	if *t.Route1 == comingFrom {
		nextPath = *t.Route2
	} else {
		nextPath = *t.Route1
	}
	switch nextPath {
	case North:
		return t.North.findRouteLength(South) + 1
	case South:
		return t.South.findRouteLength(North) + 1
	case East:
		return t.East.findRouteLength(West) + 1
	case West:
		return t.West.findRouteLength(East) + 1
	}
	return 0
}
