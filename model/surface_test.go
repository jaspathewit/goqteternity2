// Unit tests for the Tile1 entity
package model

import (
	"fmt"
	"math/rand"
	"time"

	. "gopkg.in/check.v1"
)

// Test the lines of a given tile
func (context *EternitySuite) Test_NewSurface(c *C) {
	//surface := NewSurface()

	//fmt.Printf("%s\n", surface)

}

// Test the filling from the tile set
func (context *EternitySuite) Test_AverageEnergy(c *C) {
	noSurfaces := 1000
	totalEnergy := 0
	for i := 0; i < noSurfaces; i++ {

		surface := NewSurface()
		surface.FillFromTileSet()

		valid, err := surface.IsValid()
		if !valid {
			fmt.Printf("Not Valid %s", err.Error())
			c.Fail()
		}

		totalEnergy += surface.Energy()
	}
	average := float32(totalEnergy) / float32(noSurfaces)

	fmt.Printf("Average energy %f\n", average)
}

// Test cloning the surface
func (context *EternitySuite) Test_Clone(c *C) {
	noSurfaces := 100

	for i := 0; i < noSurfaces; i++ {

		surface := NewSurface()
		surface.FillFromTileSet()

		expected := surface.String()

		clone := surface.Clone()
		actual := clone.String()

		c.Assert(actual, Equals, expected)
	}
}

// Test Rotate tile
func (context *EternitySuite) Test_RotateTileAt(c *C) {
	surface := NewSurface()
	surface.FillFromTileSet()

	// copy the original state
	clone := surface.Clone()

	// rotate all the middle tiles in the surface
	for i := 2; i < 16; i++ {
		for j := 2; j < 16; j++ {
			tile := surface.GetTile(i, j)
			tile.Rotate()
		}
	}

	expected := surface.String()

	// rotate all the tiles in the clone
	for i := range clone.Tiles {
		for j := range clone.Tiles[i] {
			tile := clone.GetTile(i, j)
			number := tile.Number
			clone.RotateTile(number)
		}
	}
	actual := clone.String()

	c.Assert(actual, Equals, expected)
}

// Test rand
func (context *EternitySuite) Test_Rand(c *C) {
	// initilise the random seed
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 1000; i++ {
		prob := rand.Intn(16)
		if prob == 16 {
			c.Fail()
		}
	}
}
