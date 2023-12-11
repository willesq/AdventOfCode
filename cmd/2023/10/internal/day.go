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

type LRS int

const (
	Left LRS = iota
	Right
	Straight
)

func Part1(filename string) (*Challenge, *int) {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	Input := Challenge{RawData: derp}
	Input.init()
	Input.navigateMap()
	derpInt := Input.TravelValues[0].Length
	for _, num := range Input.TravelValues {
		if num.Length > derpInt {
			derpInt = num.Length
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
	Input.TravelValues[0].countLRS()
	Input.triggerTiles()
	totalCount := 0
	for _, row := range Input.Data {
		for _, tile := range row {
			if tile.Triggered && !tile.LoopTile {
				totalCount++
			}
		}
	}
	return &totalCount
}

type Challenge struct {
	RawData      *[]string
	Data         [][]*Tile
	SPosition    *Tile
	TravelValues []Path
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
	results := make(chan Path)
	if c.SPosition.North.Value == "|" ||
		c.SPosition.North.Value == "7" ||
		c.SPosition.North.Value == "F" {
		// Start the north tile
		wg.Add(1)
		go func() {
			defer wg.Done()
			pathTiles, res, foundS, lfs := c.SPosition.North.findRouteLength(South)
			pathResult := Path{FoundS: foundS, DirectionFromStart: South, Length: res, LeftRightStraight: lfs, PathTiles: append([]*Tile{c.SPosition}, pathTiles...)}
			results <- pathResult
		}()
	}
	if c.SPosition.West.Value == "-" ||
		c.SPosition.West.Value == "L" ||
		c.SPosition.West.Value == "F" {
		// Start the west tile
		wg.Add(1)
		go func() {
			defer wg.Done()
			pathTiles, res, foundS, lfs := c.SPosition.West.findRouteLength(East)
			pathResult := Path{FoundS: foundS, DirectionFromStart: East, Length: res, LeftRightStraight: lfs, PathTiles: append([]*Tile{c.SPosition}, pathTiles...)}
			results <- pathResult
		}()
	}
	if c.SPosition.South.Value == "|" ||
		c.SPosition.South.Value == "L" ||
		c.SPosition.South.Value == "J" {
		// Start the south tile
		wg.Add(1)
		go func() {
			defer wg.Done()
			pathTiles, res, foundS, lfs := c.SPosition.South.findRouteLength(North)
			pathResult := Path{FoundS: foundS, DirectionFromStart: North, Length: res, LeftRightStraight: lfs, PathTiles: append([]*Tile{c.SPosition}, pathTiles...)}
			results <- pathResult
		}()
	}
	if c.SPosition.East.Value == "-" ||
		c.SPosition.East.Value == "J" ||
		c.SPosition.East.Value == "7" {
		// Start the east tile
		wg.Add(1)
		go func() {
			defer wg.Done()
			pathTiles, res, foundS, lfs := c.SPosition.East.findRouteLength(West)
			pathResult := Path{FoundS: foundS, DirectionFromStart: West, Length: res, LeftRightStraight: lfs, PathTiles: append([]*Tile{c.SPosition}, pathTiles...)}
			results <- pathResult
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

	for _, loopSlice := range c.TravelValues {
		for _, loopTile := range loopSlice.PathTiles {
			loopTile.LoopTile = true
		}
	}
}

func (c *Challenge) triggerTiles() []*Tile {
	for _, pt := range c.TravelValues[0].PathTiles {
		current := pt

		for {
			if current == nil {
				break
			}
			current.Triggered = !current.Triggered
			current = current.East
		}

	}
	return c.TravelValues[0].PathTiles
}

type Path struct {
	DirectionFromStart Route
	FoundS             *bool
	Length             int
	LeftRightStraight  []LRS
	TotalLeft          int
	TotalRight         int
	TotalStraight      int
	LFSMajority        LRS
	PathTiles          []*Tile
	InteriorTiles      []*Tile
}

func (p *Path) countLRS() {
	for _, val := range p.LeftRightStraight {
		if val == Left {
			p.TotalLeft++
		}
		if val == Right {
			p.TotalRight++
		}
		if val == Straight {
			p.TotalStraight++
		}
	}
	if p.TotalLeft > p.TotalRight {
		p.LFSMajority = Left
	} else if p.TotalRight > p.TotalLeft {
		p.LFSMajority = Right
	}
}

type Tile struct {
	North     *Tile
	South     *Tile
	East      *Tile
	West      *Tile
	Value     string
	Route1    *Route
	Route2    *Route
	LoopTile  bool
	Triggered bool
}

func (t *Tile) defineRoute1and2(valMap map[string][]Route) {
	if t.Value == "." || t.Value == "S" {
		return
	}
	t.Route1 = &valMap[t.Value][0]
	t.Route2 = &valMap[t.Value][1]
}

func (t *Tile) findRouteLength(comingFrom Route, Start ...bool) ([]*Tile, int, *bool, []LRS) {
	lfs := []LRS{}
	if t.Value == "S" {
		foundS := true // return if back at start
		return []*Tile{}, 1, &foundS, lfs
	}
	var nextPath Route
	if *t.Route1 == comingFrom {
		nextPath = *t.Route2
	} else {
		nextPath = *t.Route1
	}
	var res int
	var foundS *bool
	var lfsres []LRS
	var tRes []*Tile
	switch nextPath {
	case North:
		lfs = append(lfs, t.findLeftRightStraight(comingFrom, North))
		tRes, res, foundS, lfsres = t.North.findRouteLength(South)
	case South:
		lfs = append(lfs, t.findLeftRightStraight(comingFrom, South))
		tRes, res, foundS, lfsres = t.South.findRouteLength(North)
	case East:
		lfs = append(lfs, t.findLeftRightStraight(comingFrom, East))
		tRes, res, foundS, lfsres = t.East.findRouteLength(West)
	case West:
		lfs = append(lfs, t.findLeftRightStraight(comingFrom, West))
		tRes, res, foundS, lfsres = t.West.findRouteLength(East)
	}
	return append(tRes, t), res + 1, foundS, append(lfs, lfsres...)
}

func (t *Tile) findLeftRightStraight(from, to Route) LRS {
	opPair := map[Route]Route{
		North: South,
		South: North,
		West:  East,
		East:  West,
	}
	if opPair[from] == to {
		return Straight
	}
	var path LRS
	switch from {
	case North:
		switch to {
		case West:
			path = Right
		case East:
			path = Left
		}
	case South:
		switch to {
		case West:
			path = Left
		case East:
			path = Right
		}
	case East:
		switch to {
		case North:
			path = Right
		case South:
			path = Left
		}
	case West:
		switch to {
		case North:
			path = Left
		case South:
			path = Right
		}
	}
	return path
}

type Index struct {
	X int
	Y int
}
