// Package contains the data domain for the Eternity2 tiles.
package model

import (
	"fmt"

	"bitbucket.com/jaspathewit/eternity2/util"
)

func PrintTileStatistics() {

	// calculate the statistics for the tiles
	northTiles := TileSet[:][ORIENTATION_NORTH]
	middleTiles := filterTiles(northTiles, funcTypeOf(TYPEOF_MIDDLE))
	edgeTiles := filterTiles(northTiles, funcTypeOf(TYPEOF_EDGE))

	fmt.Printf("Middle 180 Tile (same pattern opposite sides).\n")
	tiles180 := filterTiles(middleTiles, matchingSides180)
	printTiles(tiles180)
	fmt.Printf("\n")

	fmt.Printf("Middle 90 Tile (same pattern neighbouring sides).\n")
	tiles90 := filterTiles(middleTiles, matchingSides90)
	printTiles(tiles90)
	fmt.Printf("\n")

	fmt.Printf("Middle Tile with same pattern three sides.\n")
	tilesThreeSides := filterTiles(middleTiles, matchingThreeSides)
	printTiles(tilesThreeSides)
	fmt.Printf("\n")

	// calculate statistics for the sides of all tiles
	sides := allSidesFrom(northTiles)
	keys := util.SortedKeys(sides)
	fmt.Printf("All Tile side frequencies.\n")
	printSideFrequencies(keys, sides)
	fmt.Printf("\n")

	// calculate statistics for the sides of the edge tiles
	sides = northSidesFrom(edgeTiles)
	keys = util.SortedKeys(sides)
	fmt.Printf("All edge Tile north side frequencies.\n")
	printSideFrequencies(keys, sides)
	fmt.Printf("\n")

	fmt.Printf("All Tiles grouped by matching side.\n")
	printTilesGroupedBySides(keys, sides, northTiles)
	fmt.Printf("\n")

	fmt.Printf("Middle Tiles grouped by matching side.\n")
	middleTilesGroupedByMiddleTileSides()
	fmt.Printf("\n")

	fmt.Printf("Middle 180 Tiles grouped by matching side.\n")
	middle180TilesGroupedByMiddleTileSides()
	fmt.Printf("\n")

	fmt.Printf("Middle 90 Tiles grouped by matching side.\n")
	middle90TilesGroupedByMiddleTileSides()
	fmt.Printf("\n")

	// print out which tiles match the sides ordered by the keys
	fmt.Printf("Middle Tiles grouped by match Edge tile sides.\n")
	middleTilesGroupedByEdgeTileSides()
	fmt.Printf("\n")

	// print out a dependency matrix for all tiles
	fmt.Printf("All Tiles dependency matrix.\n")
	printDependencyMatrix(northTiles)
	fmt.Printf("\n")

	/*fmt.Printf("Sides of tiles with same pattern three sides.\n")
	sides = allSidesFrom(tilesThreeSides)
	keys = util.SortedKeys(sides)
	printSideFrequencies(keys, sides)

	fmt.Printf("Tiles that have no sides that match with North sides of side tiles.\n")
	middleTilesThatCannotConnectToTheSideStatistics()

	middleTile3MatchingSides() */

}

// Function prints out the tiles that have no side that matches any of the edge
// tiles (ie tiles that cannon be connected to an Edge tile
func middleTilesThatCannotConnectToTheSideStatistics() {
	northTiles := TileSet[:][ORIENTATION_NORTH]
	// get the side tiles
	sideTiles := filterTiles(northTiles, funcTypeOf(TYPEOF_EDGE))

	// get the north sides of the side tiles
	sides := northSidesFrom(sideTiles)
	keys := util.SortedKeys(sides)
	printSideFrequencies(keys, sides)

	// get the middle tiles
	middleTiles := filterTiles(northTiles, funcTypeOf(TYPEOF_MIDDLE))

	// get the tiles that have no side that matches the edge sides
	tilesNotPossiblyAttachedToTheSide := filterTiles(middleTiles, funcHasMatchingSide(false, keys))
	printTiles(tilesNotPossiblyAttachedToTheSide)
}

// Function prints out the tiles that have a matching side to the Edge tiles
// in Edge tile side frequency
func middleTilesGroupedByEdgeTileSides() {
	northTiles := TileSet[:][ORIENTATION_NORTH]
	sideTiles := filterTiles(northTiles, funcTypeOf(TYPEOF_EDGE))
	middleTiles := filterTiles(northTiles, funcTypeOf(TYPEOF_MIDDLE))
	sides := northSidesFrom(sideTiles)
	keys := util.SortedKeys(sides)
	printTilesGroupedBySides(keys, sides, middleTiles)
}

// Function prints out dependency matrix for the given tiles
func printDependencyMatrix(tiles []*Tile1) {
	// print the column headers
	fmt.Printf("    ")
	for _, tile := range tiles {
		fmt.Printf("%s ", tile.GetNumberAsString())
	}

	fmt.Printf("\n")

	// loop through all the tiles
	for _, tileA := range tiles {
		fmt.Printf("%s ", tileA.GetNumberAsString())

		for _, tileB := range tiles {

			number := " "
			if tileA.CanConnectTo(tileB) {
				// get the number of sides in common between the two tiles
				if tileA.NumberOfSidesMatchingTile(tileB) > 0 {
					number = "*"
				}
			}

			// fmt.Printf("% 3d ", number)
			fmt.Printf("% 3s ", number)
		}

		fmt.Printf("\n")
	}
}

// Function prints out the middle tiles grouped by side
func middleTilesGroupedByMiddleTileSides() {
	northTiles := TileSet[:][ORIENTATION_NORTH]
	middleTiles := filterTiles(northTiles, funcTypeOf(TYPEOF_MIDDLE))
	sides := allSidesFrom(middleTiles)
	keys := util.SortedKeys(sides)
	printTilesGroupedBySides(keys, sides, middleTiles)
}

// Function prints out the middle 180 degree tiles grouped by side
func middle180TilesGroupedByMiddleTileSides() {
	northTiles := TileSet[:][ORIENTATION_NORTH]
	middleTiles := filterTiles(northTiles, funcTypeOf(TYPEOF_MIDDLE))
	middle180Tiles := filterTiles(middleTiles, matchingSides180)
	sides := allSidesFrom(middle180Tiles)
	keys := util.SortedKeys(sides)
	printTilesGroupedBySides(keys, sides, middle180Tiles)
}

// Function prints out the middle 180 degree tiles grouped by side
func middle90TilesGroupedByMiddleTileSides() {
	northTiles := TileSet[:][ORIENTATION_NORTH]
	middleTiles := filterTiles(northTiles, funcTypeOf(TYPEOF_MIDDLE))
	middle90Tiles := filterTiles(middleTiles, matchingSides90)
	sides := allSidesFrom(middle90Tiles)
	keys := util.SortedKeys(sides)
	printTilesGroupedBySides(keys, sides, middle90Tiles)
}

func middleTile3MatchingSides() {
	fmt.Printf("Tiles that have three sides in common.\n")
	northTiles := TileSet[:][ORIENTATION_NORTH]

	// get the middle tiles
	middleTiles := filterTiles(northTiles, funcTypeOf(TYPEOF_MIDDLE))

	for _, middleTile := range middleTiles {
		tile1 := middleTile.Clone()
		for _, middleTile := range middleTiles {
			tile2 := middleTile.Clone()
			for o := 0; o < 4; o++ {
				matchingSides, nonMatchingSides := tile1.SidesMatching(tile2)
				if len(matchingSides) == 3 {
					fmt.Printf("Tile1 %s Tile2 %s MatchingSides %v NonMatching %v\n",
						tile1.GetNumberAsString(),
						tile2.GetNumberAsString(),
						matchingSides,
						nonMatchingSides)
				}
				tile2.Rotate()
			}
		}
	}
}

// filter the set of tiles from the given tile set occording to the predicate function
func filterTiles(tiles []*Tile1, predicate func(*Tile1) bool) []*Tile1 {
	var result []*Tile1
	for _, tile := range tiles {
		clone := tile.Clone()
		if predicate(clone) {
			result = append(result, clone)
		}
	}
	return result
}

// Function returns true if the tile has matching sides 180 degrees
func matchingSides180(tile *Tile1) bool {
	result := (tile.Top == tile.Bottom) ||
		(tile.Left == tile.Right) &&
			!matchingThreeSides(tile)
	return result
}

// Function returns true if the tile has matching sides 90 degrees
func matchingSides90(tile *Tile1) bool {
	result := tile.TypeOf != TYPEOF_CORNER &&
		!matchingThreeSides(tile) &&
		(tile.Top == tile.Right ||
			tile.Right == tile.Bottom ||
			tile.Bottom == tile.Left ||
			tile.Left == tile.Top)
	return result
}

// Function returns true if the tile has matching sides 90 degrees
func matchingThreeSides(tile *Tile1) bool {
	clone := tile.Clone()
	result := false
	for i := 0; i < 3; i++ {
		result = (clone.Top == clone.Right &&
			clone.Right == clone.Bottom)
		if result {
			break
		}
		clone.Rotate()
	}

	return result
}

// Function returns a function that tests a tile for the given type
func funcTypeOf(typeOf byte) func(*Tile1) bool {
	function := func(tile *Tile1) bool {
		result := tile.TypeOf == typeOf
		return result
	}
	return function
}

// Function returns a function that test a tile any sides that matches
// any of the sides given in the slice of sides. If Should match is true the function returns true
// if one of the sides matches if its false it returns false if one of the sides matches
func funcHasMatchingSide(shouldMatch bool, sides []string) func(*Tile1) bool {

	// define the function
	function := func(tile *Tile1) bool {
		result := !shouldMatch
		clone := tile.Clone()
		// run through the four orientations
		for i := 0; i < 4; i++ {
			// get the side definition
			side, _ := SideToDefinition(clone.Top)

			// see if this side is in the sides
			found := util.Contains(sides, side)

			// if its found then the side matches
			if found {
				result = shouldMatch
				break
			}

			// rotate the tile one position
			clone.Rotate()
		}

		return result
	}

	return function
}

// Return a map containing the different sides and the count in the given tiles
func allSidesFrom(tiles []*Tile1) map[string]int {
	result := make(map[string]int)
	for _, tile := range tiles {
		clone := tile.Clone()
		for i := 0; i < 4; i++ {
			side, _ := SideToDefinition(clone.Top)
			result[side] += 1
			clone.Rotate()
		}
	}
	return result
}

// Return a map containing the different sides and the count in the given tiles
func northSidesFrom(tiles []*Tile1) map[string]int {
	result := make(map[string]int)
	for _, tile := range tiles {
		side, _ := SideToDefinition(tile.Top)
		result[side] += 1
	}
	return result
}

// Print out the set of tiles
func printTiles(tiles []*Tile1) {
	for _, tile := range tiles {
		fmt.Printf("Tile %s\n", tile.GetNumberAsString())
	}
	fmt.Printf("Total %d\n", len(tiles))
}

// Print out the sides ordered by key
func printSideFrequencies(keys []string, sides map[string]int) {
	for _, key := range keys {
		fmt.Printf("Side %s, %d\n", key, sides[key])
	}
	fmt.Printf("Total %d\n", len(sides))
}

// Print out which tiles match the sides ordered by key
func printTilesGroupedBySides(keys []string, sides map[string]int, tiles []*Tile1) {
	for _, key := range keys {
		// get the frequency for the given key
		frequency := sides[key]

		// make a slice containing the single side to match
		definitionToMatch := make([]string, 1)
		definitionToMatch[0] = key

		// get the tiles that match this side
		matchingTiles := filterTiles(tiles, funcHasMatchingSide(true, definitionToMatch))

		// get a slice containing just the string references for the tiles
		var numbers []string
		for _, tile := range matchingTiles {

			number := " "

			// if this tile has more than one side that matches the definition prefix the number with *
			if tile.NumberOfSidesMatchingDefinitions(definitionToMatch) > 1 {
				number = "*"
			}

			number = tile.GetNumberAsString() + number
			numbers = append(numbers, number)
		}

		fmt.Printf("Side %s, %d : %v\n", key, frequency, numbers)
	}
}
