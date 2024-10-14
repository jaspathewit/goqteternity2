// Single Eternity tile.
package model

// the location of a tile, basically the coordinates
type Location struct {
	Row    int
	Column int
}

// function clones a location
func (location *Location) Clone() *Location {
	result := new(Location)

	result.Row = location.Row
	result.Column = location.Column

	return result
}
