// Creation of Eternity tile surface.
package model

import (
	"math/rand"
	"time"
)

// Initilise the random number generator
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Create a new surface for the eternity tiles.
func NewSurface() *Surface {
	result := new(Surface)

	// make the slices for the tiles
	result.Tiles = make([][]*Tile1, 18)
	for i := range result.Tiles {
		result.Tiles[i] = make([]*Tile1, 18)
	}

	// make the slice for the locations
	result.Locations = make([]*Location, 256)

	for i := range result.Locations {
		result.Locations[i] = new(Location)
	}

	return result
}

// FillFromTileSet fills the surface with tiles randomly from the set of possible tiles.
// Its not entierly random as the only corner, edge and middle tiles are place where
// apropriate.
func (surface *Surface) FillFromTileSet() {

	// fill the surface with border tiles
	surface.fillWith(TILE1INSTANCE)

	surface.fillCornersFromTileSet()
	surface.fillEdgesFromTileSet()
	surface.fillMiddlesFromTileSet()
}

// Clone clones the surface.
// Returns a pointer to the newly created Surface.
func (surface *Surface) Clone() *Surface {
	result := NewSurface()

	// clone its tiles
	for i := range surface.Tiles {
		for j := range surface.Tiles[i] {
			tile := surface.GetTile(i, j)
			clone := tile.Clone()
			result.SetTile(i, j, clone)
		}
	}
	return result
}

// fillWith fills the surface with the given tile.
func (surface *Surface) fillWith(tile *Tile1) {
	for i := range surface.Tiles {
		for j := range surface.Tiles[i] {
			surface.SetTile(i, j, tile)
		}
	}
}

// fillCornersFromTileSet fill corners of the surface with random corner tiles
// It's not random as the four corner tiles will be used but they are place randomly.
func (surface *Surface) fillCornersFromTileSet() {
	// add the four corner tiles randomly
	indexes := rand.Perm(4)

	// Top Left
	index := indexes[0]
	tile := TileSet[ORIENTATION_EAST][index]
	clone := tile.Clone()
	surface.SetTile(1, 1, clone)

	// Top Right
	index = indexes[1]
	tile = TileSet[ORIENTATION_SOUTH][index]
	clone = tile.Clone()
	surface.SetTile(1, 16, clone)

	// Bottom Right
	index = indexes[2]
	tile = TileSet[ORIENTATION_WEST][index]
	clone = tile.Clone()
	surface.SetTile(16, 16, clone)

	// Bottom Left
	index = indexes[3]
	tile = TileSet[ORIENTATION_NORTH][index]
	clone = tile.Clone()
	surface.SetTile(16, 1, clone)
}

// fillEdgesFromTileSet fills the edges with random Edge tiles.
// It's not random as the 56 Edge tiles will be used but they are place randomly.
func (surface *Surface) fillEdgesFromTileSet() {
	tileSetOffset := 4
	// add the 56 side tiles
	indexes := rand.Perm(56)

	// Top
	col := 2
	for i := 0; i < 14; i++ {
		index := indexes[i] + tileSetOffset
		tile := TileSet[ORIENTATION_SOUTH][index]
		clone := tile.Clone()
		surface.SetTile(1, col, clone)
		col++
	}

	// Bottom
	col = 2
	for i := 14; i < 28; i++ {
		index := indexes[i] + tileSetOffset
		tile := TileSet[ORIENTATION_NORTH][index]
		clone := tile.Clone()
		surface.SetTile(16, col, clone)
		col++
	}

	// Left
	row := 2
	for i := 28; i < 42; i++ {
		index := indexes[i] + tileSetOffset
		tile := TileSet[ORIENTATION_EAST][index]
		clone := tile.Clone()
		surface.SetTile(row, 1, clone)
		row++
	}

	// Right
	row = 2
	for i := 42; i < 56; i++ {
		index := indexes[i] + tileSetOffset
		tile := TileSet[ORIENTATION_WEST][index]
		clone := tile.Clone()
		surface.SetTile(row, 16, clone)
		row++
	}
}

// fillMiddlesFromTileSet fills the middle with random Middle tiles.
// It's not random as the 196 middle tiles will be used but they are place randomly.
func (surface *Surface) fillMiddlesFromTileSet() {
	tileSetOffset := 60
	// add the 196 middle tiles
	indexes := rand.Perm(196)

	pos := 0
	for j := 2; j < 16; j++ {
		for i := 2; i < 16; i++ {
			index := indexes[pos] + tileSetOffset
			orientation := rand.Intn(4)
			tile := TileSet[orientation][index]
			clone := tile.Clone()
			surface.SetTile(j, i, clone)
			pos++
		}
	}
}
