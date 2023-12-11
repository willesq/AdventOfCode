package adventhelper

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

type Tile struct {
	North     *Tile
	South     *Tile
	East      *Tile
	West      *Tile
	Value     string
	Triggered bool
}

func Create4LinkedGrid(data *[]string) [][]*Tile {
	returnData := [][]*Tile{}
	for lineIdx, line := range *data {
		lineSlice := []*Tile{}
		var tmpTile *Tile
		for charIdx, char := range line {
			charStr := string(char)
			tmpTile = &Tile{Value: charStr}
			if lineIdx != 0 { // North  & South Path
				// Ignored on the first line in the grid
				// Set Current Tile North Value
				tmpTile.North = returnData[lineIdx-1][charIdx]
				// Set Previous Line same tile Index South value
				preprevTile := returnData[lineIdx-1][charIdx]
				preprevTile.South = tmpTile
			}
			if charIdx != 0 { // West & East Path
				// Ignores the first element in the line
				// Set current Tile west Value
				tmpTile.West = lineSlice[charIdx-1]
				// Set Previous Tile East Value
				lineSlice[charIdx-1].East = tmpTile
			}

			lineSlice = append(lineSlice, tmpTile)
		}
		returnData = append(returnData, lineSlice)
	}
	return returnData
}
