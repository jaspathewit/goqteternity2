// Unit tests for the Tile1 entity
package model

import (
	"fmt"

	. "gopkg.in/check.v1"
)

// Test the lines of a given tile
func (context *EternitySuite) Test_Tile1_Lines(c *C) {
	expected := "+--T--+\n" +
		"|  *  |\n" +
		"A 011 D\n" +
		"|     |\n" +
		"+--X--+\n"

	tile := TileSet[ORIENTATION_NORTH][byte(10)]

	actual := tile.String()
	c.Assert(actual, Equals, expected)
}

// Test the cloning of a given tile
func (context *EternitySuite) Test_Tile1_Clone(c *C) {
	expected := "+--T--+\n" +
		"|  *  |\n" +
		"A 011 D\n" +
		"|     |\n" +
		"+--X--+\n"

	tile := TileSet[ORIENTATION_NORTH][byte(10)].Clone()

	actual := tile.String()
	c.Assert(actual, Equals, expected)
}

// Test the print lines of a given tile
func (context *EternitySuite) Test_Tile1_CloneEnergy(c *C) {
	tile := TileSet[ORIENTATION_NORTH][byte(10)].Clone()

	tile.Energy = 5

	clone := tile.Clone()

	c.Assert(tile.Energy, Equals, clone.Energy)
}

// Test the rotation of a given tile
func (context *EternitySuite) Test_Tile1_Rotate(c *C) {
	expected := "+--A--+\n" +
		"|     |\n" +
		"X 011*T\n" +
		"|     |\n" +
		"+--D--+\n"

	tile := TileSet[ORIENTATION_NORTH][byte(10)].Clone()

	tile.Rotate()

	actual := tile.String()
	c.Assert(actual, Equals, expected)
}

// Test the print lines of a given tile
func (context *EternitySuite) Test_Tile1_RotateTill(c *C) {
	expectedTile := TileSet[ORIENTATION_WEST][byte(10)].Clone()

	expected := expectedTile.String()

	tile := TileSet[ORIENTATION_NORTH][byte(10)].Clone()
	tile.RotateTill(ORIENTATION_WEST)

	actual := tile.String()
	c.Assert(actual, Equals, expected)
}

// Test the sides matching between two tiles
func (context *EternitySuite) Test_Tile1_SidesMatching(c *C) {
	tile := TileSet[ORIENTATION_NORTH][byte(10)].Clone()

	test := TileSet[ORIENTATION_NORTH][byte(10)].Clone()

	matching, nonMatching := tile.SidesMatching(test)
	c.Assert(len(matching), Equals, 4)
	c.Assert(len(nonMatching), Equals, 0)

	test.Rotate()

	matching, nonMatching = tile.SidesMatching(test)
	c.Assert(len(matching), Equals, 0)
	c.Assert(len(nonMatching), Equals, 8)
}

// Test the Number of matching sides
func (context *EternitySuite) Test_Tile1_NumberOfMatchingSides(c *C) {
	tile := TileSet[ORIENTATION_NORTH][byte(172)].Clone()
	definitions := []string{"070607"}
	numberOf := tile.NumberOfSidesMatchingDefinitions(definitions)

	c.Assert(numberOf, Equals, 1)

	definitions = []string{"100406"}
	numberOf = tile.NumberOfSidesMatchingDefinitions(definitions)

	c.Assert(numberOf, Equals, 3)
}

// Test the round trip marshalling
func (context *EternitySuite) Test_Tile1_Round_Trip(c *C) {
	for j := 0; j < 4; j++ {
		for i := 0; i < 256; i++ {
			expected := TileSet[j][i].Clone()

			bytes := expected.Marshal()
			actual := TILE1INSTANCE.Unmarshal(bytes)
			c.Assert(actual.String(), Equals, expected.String())
		}
	}
}

// Test the paterns of the tiles
func (context *EternitySuite) Test_Tile1_Patterns(c *C) {
	for j := 0; j < 4; j++ {
		for i := 0; i < 256; i++ {
			tile := TileSet[j][i].Clone()
			c.Assert(tile.GetTopPattern(), Equals, fmt.Sprintf("%c", tile.Top))
			c.Assert(tile.GetRightPattern(), Equals, fmt.Sprintf("%c", tile.Right))
			c.Assert(tile.GetBottomPattern(), Equals, fmt.Sprintf("%c", tile.Bottom))
			c.Assert(tile.GetLeftPattern(), Equals, fmt.Sprintf("%c", tile.Left))
		}
	}
}

// Test the ids of the tiles
func (context *EternitySuite) Test_Tile1_Ids(c *C) {
	for j := 0; j < 4; j++ {
		for i := 0; i < 256; i++ {
			tile := TileSet[j][i].Clone()
			// Number, orientation
			bytes := []byte{byte(i), byte(j)}
			c.Assert(tile.GetId(), Equals, fmt.Sprintf("%X", bytes))
		}
	}
}
