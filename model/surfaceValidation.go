// Validation methods for a Surface.
package model

import (
	"errors"
	"fmt"
	// "os"
	// "strconv"
	// "math/rand"
	// "time"
)

// function determins if the surface is valid
func (surface *Surface) IsValid() (bool, error) {
	result, err := surface.isValidCorners()
	if result {
		result, err = surface.isValidSides()
		if result {
			result, err = surface.isValidMiddles()

			if result {
				result, err = surface.isValidLocations()
			}
		}
	}

	return result, err
}

// function validates the corner tiles
func (surface *Surface) isValidCorners() (bool, error) {
	// slice to record the tiles used
	tiles := make([]*Tile1, 4)

	// check Top Left
	tile := surface.GetTile(1, 1)
	if tile.Orientation != ORIENTATION_EAST {
		return false, errors.New("Tile orientation is wrong for top left")
	}
	tiles[tile.Number] = tile

	// Top Right
	tile = surface.GetTile(1, 16)
	if tile.Orientation != ORIENTATION_SOUTH {
		return false, errors.New("Tile orientation is wrong for top right")
	}
	tiles[tile.Number] = tile

	// Bottom Right
	tile = surface.GetTile(16, 16)
	if tile.Orientation != ORIENTATION_WEST {
		return false, errors.New("Tile orientation is wrong for bottom right")
	}
	tiles[tile.Number] = tile

	// Bottom Left
	tile = surface.GetTile(16, 1)
	if tile.Orientation != ORIENTATION_NORTH {
		return false, errors.New("Tile orientation is wrong for bottom left")
	}
	tiles[tile.Number] = tile

	result, err := checkTilesUsed(tiles, TYPEOF_CORNER)

	return result, err
}

// function validates the corner tiles
func (surface *Surface) isValidSides() (bool, error) {
	tileSetOffset := byte(4)
	// slice to record the tiles used
	tiles := make([]*Tile1, 56)

	// Top
	for i := 2; i < 16; i++ {
		tile := surface.GetTile(1, i)
		if tile.Orientation != ORIENTATION_SOUTH {
			return false, errors.New("Tile orientation is wrong for top side")
		}
		tiles[tile.Number-tileSetOffset] = tile
	}

	// Bottom
	for i := 2; i < 16; i++ {
		tile := surface.GetTile(16, i)
		if tile.Orientation != ORIENTATION_NORTH {
			return false, errors.New("Tile orientation is wrong for bottom side")
		}
		tiles[tile.Number-tileSetOffset] = tile
	}

	// Left
	for i := 2; i < 16; i++ {
		tile := surface.GetTile(i, 1)
		if tile.Orientation != ORIENTATION_EAST {
			return false, errors.New("Tile orientation is wrong for left side")
		}
		tiles[tile.Number-tileSetOffset] = tile
	}

	// Right
	for i := 2; i < 16; i++ {
		tile := surface.GetTile(i, 16)
		if tile.Orientation != ORIENTATION_WEST {
			return false, errors.New("Tile orientation is wrong for right side")
		}
		tiles[tile.Number-tileSetOffset] = tile
	}

	result, err := checkTilesUsed(tiles, TYPEOF_EDGE)

	return result, err
}

// function validates the middle tiles
func (surface *Surface) isValidMiddles() (bool, error) {
	tileSetOffset := byte(60)
	// slice to record the tiles used
	tiles := make([]*Tile1, 196)

	for i := 2; i < 16; i++ {
		for j := 2; j < 16; j++ {
			tile := surface.GetTile(i, j)
			tiles[tile.Number-tileSetOffset] = tile
		}
	}

	result, err := checkTilesUsed(tiles, TYPEOF_MIDDLE)

	return result, err
}

// function validates the middle tiles
func (surface *Surface) isValidLocations() (bool, error) {
	result := true
	var err error

	for i := 1; i < 17; i++ {
		for j := 1; j < 17; j++ {
			tile := surface.GetTile(i, j)
			number := tile.Number
			location := surface.Locations[number]
			if location.Row != i ||
				location.Column != j {
				err = errors.New(fmt.Sprintf("Tile %s is at location (%d,%d) but the location is recorded as (%d,%d)",
					tile.GetNumberAsString(), i, j, location.Row, location.Column))
				break
			}
		}
	}
	return result, err
}

// checks that all the tiles in the tiles slice or of the required type
// also checks that there are no nil pointers in the slice
func checkTilesUsed(tiles []*Tile1, typeOf byte) (bool, error) {
	result := true
	var err error

	for _, tile := range tiles {
		if tile == nil {
			result = false
			err = errors.New("Missing tile")
			break
		}

		if tile.TypeOf != typeOf {
			result = false
			err = errors.New("Tile is not of the correct type")
			break
		}
	}

	return result, err
}
