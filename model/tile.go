// Constants used in the definition of Eternity tiles.
package model

// Tile type constants
const (
	TYPEOF_BORDER = byte(0)
	TYPEOF_CORNER = byte(1)
	TYPEOF_EDGE   = byte(2)
	TYPEOF_MIDDLE = byte(3)
)

// constants defining Orientation of tile
const (
	ORIENTATION_UNKNOWN = byte('U')
	ORIENTATION_NORTH   = byte(0)
	ORIENTATION_EAST    = byte(1)
	ORIENTATION_SOUTH   = byte(2)
	ORIENTATION_WEST    = byte(3)
)
