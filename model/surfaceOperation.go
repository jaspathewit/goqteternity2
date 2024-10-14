// Operations performed on a surface of Eternity tiles.
package model

// "errors"
// "fmt"
// "os"
// "strconv"

// SetTile sets the given tile pointer at the given row and col in the surface.
// Updates the Locations to reflect the change
func (surface *Surface) SetTile(row int, col int, tile *Tile1) {
	// update the location of this tile
	// if its a non Border tile
	if tile.TypeOf != TYPEOF_BORDER {
		number := tile.Number
		location := surface.Locations[number]
		location.Row = row
		location.Column = col
	}

	// put the tile in the surface
	surface.Tiles[row][col] = tile
}

// GetTile returns a pointer to the tile at the given row and col.
func (surface *Surface) GetTile(row int, col int) *Tile1 {
	result := surface.Tiles[row][col]
	return result
}

// GetLocation returns the current location of the tile with the given tile number.
func (surface *Surface) GetLocation(number byte) *Location {
	result := surface.Locations[number]
	return result
}

// Rotate the tile at the given row and column positions
func (surface *Surface) RotateTile(tileNumber byte) {
	location := surface.GetLocation(tileNumber)
	tile := surface.GetTile(location.Row, location.Column)

	// only rotate middle tiles
	if tile.TypeOf == TYPEOF_MIDDLE {
		tile.Rotate()
	}
}

// Switch tiles at the given source and destination positions
func (surface *Surface) SwitchTiles(srcNumber byte, destNumber byte) {
	srcLocation := surface.GetLocation(srcNumber).Clone()
	srcTile := surface.GetTile(srcLocation.Row, srcLocation.Column)

	destLocation := surface.GetLocation(destNumber).Clone()
	destTile := surface.GetTile(destLocation.Row, destLocation.Column)

	// only switch matching tiles
	if srcTile.TypeOf == destTile.TypeOf {
		srcOrientation := srcTile.Orientation
		destOrientation := destTile.Orientation

		// orientate the tiles correctly
		srcTile.RotateTill(destOrientation)
		destTile.RotateTill(srcOrientation)

		// switch them
		surface.SetTile(destLocation.Row, destLocation.Column, srcTile)
		surface.SetTile(srcLocation.Row, srcLocation.Column, destTile)
	}
}
