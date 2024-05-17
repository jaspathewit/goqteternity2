// A surface of Eternity tiles.
package model

import (
	// "errors"
	// "fmt"
	// "os"
	// "strconv"
	"bytes"
	"fmt"
)

// Type defining a surface of tiles
type Surface struct {
	// The tiles forming the surface
	Tiles [][]*Tile1

	// the locations of a particular tile
	Locations []*Location
}

// calculate the energy of the surface (updates the individual tile energies as a side effect)
func (surface *Surface) Energy() int {
	result := 0
	for i := 1; i < 17; i++ {
		for j := 1; j < 17; j++ {

			tile := surface.GetTile(i, j)
			energy := 0

			// check north
			north := surface.north(i, j)
			if tile.Top != north.Bottom {
				energy++
			}

			// check west
			west := surface.west(i, j)
			if tile.Left != west.Right {
				energy++
			}

			// check south
			//south := surface.south(i, j)
			//if tile.Bottom != south.Top {
			//	energy++
			//}

			// check east
			//east := surface.east(i, j)
			//if tile.Right != east.Left {
			//	energy++
			//}

			tile.Energy = energy
			result += tile.Energy
		}
	}

	return result
}

// returns the tile to the north of the given position
func (surface *Surface) north(row int, col int) *Tile1 {
	result := surface.GetTile(row-1, col)
	return result
}

// returns the tile to the east of the given position
func (surface *Surface) east(row int, col int) *Tile1 {
	result := surface.GetTile(row, col+1)
	return result
}

// returns the tile to the south of the given position
func (surface *Surface) south(row int, col int) *Tile1 {
	result := surface.GetTile(row+1, col)
	return result
}

// returns the tile to the westh of the given position
func (surface *Surface) west(row int, col int) *Tile1 {
	result := surface.Tiles[row][col-1]
	return result
}

// function returns the Surface as a string
// change the for ranges
func (surface *Surface) String() string {
	var result string
	for i := range surface.Tiles {
		for lineIndex := 0; lineIndex < 5; lineIndex++ {
			for j := range surface.Tiles[i] {
				tile := surface.GetTile(i, j)
				if tile != nil {
					result += tile.Line(lineIndex)
				}
			}
			result += "\n"
		}
	}
	return result
}

// function returns the Surface as a html representation
func (surface *Surface) Html() string {
	var buffer bytes.Buffer
	header := getHtmlHeader()
	buffer.WriteString(header)

	buffer.WriteString("    <table>")
	for i := range surface.Tiles {
		buffer.WriteString("      <tr>")
		for j := range surface.Tiles[i] {
			buffer.WriteString("        <td>")
			tile := surface.GetTile(i, j)
			if tile != nil {
				buffer.WriteString(fmt.Sprintf("<img src='%0.3d' alt='%0.3d' />", tile.Number, tile.Number))
				//result += tile.Line(lineIndex)
			}
			buffer.WriteString("        </td>\n")
		}
		buffer.WriteString("      </tr>\n")
	}
	buffer.WriteString("    </table>\n")
	return buffer.String()
}

// function returns the Surface as a html representation
func getHtmlHeader() string {
	var buffer bytes.Buffer
	buffer.WriteString("<!DocType html>")
	buffer.WriteString("<html>")
	buffer.WriteString("  <body>")

	return buffer.String()
}
