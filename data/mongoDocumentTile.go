// Single Eternity tile.
package data

import (
	"time"
)

// Eternity tile
type MongoDocumentTile struct {
	// The identity of the tile
	Id          string
	Status      byte
	Date        time.Time
	TypeOf      byte
	Orientation byte

	// The side patterns of the tile
	TopPattern    string
	RightPattern  string
	BottomPattern string
	LeftPattern   string

	TileData []byte
}
